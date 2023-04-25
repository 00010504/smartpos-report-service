package storage

import (
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Invan2/invan_report_service/pkg/logger"
	chdb "github.com/Invan2/invan_report_service/storage/clickhouse"
	"github.com/Invan2/invan_report_service/storage/repo"
)

type storageCh struct {
	orderRepo repo.OrderICH
}

type StorageCh interface {
	Order() repo.OrderICH
}

func NewStorageCh(db clickhouse.Conn, log logger.Logger) StorageCh {
	return &storageCh{
		orderRepo: chdb.NewOrderCHRepo(db, log),
	}
}

func (s *storageCh) Order() repo.OrderICH {
	return s.orderRepo
}
