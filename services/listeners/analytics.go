package listeners

import (
	"context"
	"genproto/report_service"

	"github.com/Invan2/invan_report_service/pkg/logger"
)

func (r *reportService) GetDashboardAnalytics(ctx context.Context, message *report_service.GetDashboardAnalyticsReq) (res *report_service.GetDashboardAnalyticsRes, err error) {
	r.log.Info("GetDashboardAnalytics", logger.Any("req", message))
	res, err = r.strgCH.Order().GetDashboardAnalytics(ctx, message)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return &report_service.GetDashboardAnalyticsRes{
				Response: &report_service.GetDashboardAnalyticsResponse{},
			}, nil
		} else {
			r.log.Error("Error occured while execute GetDashboardAnalytics: ", logger.Any("Error", err))
			return nil, err
		}
	}
	return
}
