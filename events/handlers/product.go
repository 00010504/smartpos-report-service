package handlers

import (
	"context"
	"encoding/json"
	"genproto/common"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/pkg/errors"
)

func (e *EventHandler) UpsertProduct(ctx context.Context, event *kafka.Message) error {

	var req common.CreateProductCopyRequest

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return errors.Wrap(err, "error while unmarshaling product")
	}

	e.log.Info("create product event", logger.Any("event", req))

	err := e.strg.Product().Upsert(ctx, &req)
	if err != nil {
		return err
	}

	return nil

}
