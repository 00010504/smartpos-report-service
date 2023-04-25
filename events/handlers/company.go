package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"genproto/common"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateCompany(ctx context.Context, event *kafka.Message) error {

	var req common.CompanyCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("create company event", logger.Any("event", req))

	if err := e.strg.Company().Upsert(ctx, &req); err != nil {
		e.log.Info(err.Error(), logger.Any("event", req))

		return err
	}
	e.log.Info("company shop is about to create", logger.Any("event", req))

	if req.Shop == nil {
		return errors.New("nimadir")
	}

	if err := e.strg.Shop().Upsert(req.Shop); err != nil {
		return err
	}

	return nil

}
