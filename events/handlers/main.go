package handlers

import (
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage"
)

type EventHandler struct {
	log    logger.Logger
	strg   storage.StoragePgI
	strgCh storage.StorageCh
}

func NewHandler(log logger.Logger, strg storage.StoragePgI, strgCh storage.StorageCh) *EventHandler {
	return &EventHandler{
		log:    log,
		strg:   strg,
		strgCh: strgCh,
	}
}
