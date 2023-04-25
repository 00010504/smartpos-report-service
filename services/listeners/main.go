package listeners

import (
	"context"
	"genproto/report_service"

	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage"
)

type reportService struct {
	strgPG storage.StoragePgI
	strgCH storage.StorageCh
	log    logger.Logger
}

type ReportService interface {
	GetDashboardAnalytics(ctx context.Context, message *report_service.GetDashboardAnalyticsReq) (*report_service.GetDashboardAnalyticsRes, error)
}

func NewReportService(strg storage.StoragePgI, strgCH storage.StorageCh, log logger.Logger) ReportService {
	return &reportService{
		strgPG: strg,
		strgCH: strgCH,
		log:    log,
	}
}
