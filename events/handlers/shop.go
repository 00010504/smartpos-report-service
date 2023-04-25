package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateShop(ctx context.Context, event *kafka.Message) error {

	var req common.ShopCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err

	}

	e.log.Info("shop created", logger.Any("event", req))

	if err := e.strg.Shop().Upsert(&req); err != nil {
		return err
	}

	return nil

}
