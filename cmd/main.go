package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"genproto/report_service"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Invan2/invan_report_service/config"
	"github.com/Invan2/invan_report_service/events"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/services/listeners"
	"github.com/Invan2/invan_report_service/storage"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {

	cfg := config.Load()
	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
	case config.TestMode:
		loggerLevel = logger.LevelDebug
	default:
		loggerLevel = logger.LevelInfo
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)

	log.Info("config", logger.Any("data", cfg), logger.Any("env", os.Environ()))

	defer logger.Cleanup(log)

	log.Info("config", logger.Any("data", cfg), logger.Any("env", os.Environ()))

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	chDB, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{fmt.Sprintf("%s:%d", cfg.ClickHouseHost, cfg.ClickHousePort)},
		Auth: clickhouse.Auth{
			Database: cfg.ClickHouseDatabase,
			Username: cfg.ClickHouseUser,
			Password: cfg.ClickHousePassword,
		},
	})
	if err != nil {
		log.Error("error while connect clickhouse", logger.Error(err))
		return
	}

	err = chDB.Ping(ctx)
	if err != nil {
		log.Error("ping to click house", logger.Error(err))
		return
	}

	strgCh := storage.NewStorageCh(chDB, log)

	log.Info("connected to clickhouse db")

	postgresURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := pgx.Connect(ctx, postgresURL)
	if err != nil {
		log.Error("poostgres", logger.Error(err))
		return
	}

	strg := storage.NewStoragePg(log, psqlConn)

	conf := kafka.ConfigMap{
		"bootstrap.servers":                     cfg.KafkaUrl,
		"group.id":                              config.ConsumerGroupID,
		"auto.offset.reset":                     "earliest",
		"go.events.channel.size":                1000000,
		"socket.keepalive.enable":               true,
		"metadata.max.age.ms":                   900000,
		"metadata.request.timeout.ms":           30000,
		"retries":                               1000000,
		"message.timeout.ms":                    300000,
		"socket.timeout.ms":                     30000,
		"max.in.flight.requests.per.connection": 5,
		"heartbeat.interval.ms":                 3000,
		"enable.idempotence":                    true,
	}

	log.Info("kafka config", logger.Any("config", conf))

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Error("error while creating producer")
		return
	}

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Error("error while creating consumer", logger.Error(err))
		return
	}

	pubsubServer, err := events.NewPubSubServer(log, producer, consumer, strg, strgCh)
	if err != nil {
		log.Fatal("error creating pubSubServer", logger.Error(err))
		return
	}

	server := grpc.NewServer()

	listeners := listeners.NewReportService(strg, strgCh, log)

	report_service.RegisterReportServiceServer(server, listeners)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", cfg.HttpHost, cfg.HttpPort))
	if err != nil {
		log.Error("http", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
		if err := pubsubServer.Shutdown(); err != nil {
			log.Error("error while shutdown pub sub server")
			return
		}
	}()

	go pubsubServer.Run(ctx)

	if err := server.Serve(lis); err != nil {
		log.Fatal("serve", logger.Error(err))
		return
	}
}
