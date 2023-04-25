package repo

import (
	"context"
	"genproto/common"
	"genproto/report_service"
)

type OrderICH interface {
	Upsert(ctx context.Context, entity *common.OrderCopyRequest) error
	GetDashboardAnalytics(ctx context.Context, message *report_service.GetDashboardAnalyticsReq) (res *report_service.GetDashboardAnalyticsRes, err error)
}
