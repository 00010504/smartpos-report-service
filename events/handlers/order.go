package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) UpsertOrder(ctx context.Context, event *kafka.Message) error {

	var req common.OrderCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return errors.Wrap(err, "error while unmarshaling order")
	}

	e.log.Info("create order event", logger.Any("event", req))

	err := e.strgCh.Order().Upsert(ctx, &req)
	if err != nil {
		e.log.Error("Error occured while upsert order ", logger.Any("Order().Upsert", err.Error()))
		return err
	}

	return nil

}
