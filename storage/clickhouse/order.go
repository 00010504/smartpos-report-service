package clickhouse

import (
	"context"
	"fmt"
	"genproto/common"
	"genproto/report_service"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/Invan2/invan_report_service/config"
	"github.com/Invan2/invan_report_service/models"
	"github.com/Invan2/invan_report_service/pkg/helper"
	"github.com/Invan2/invan_report_service/pkg/logger"
	"github.com/Invan2/invan_report_service/storage/repo"
	"github.com/google/uuid"
)

type orderRepo struct {
	db  clickhouse.Conn
	log logger.Logger
}

func NewOrderCHRepo(db clickhouse.Conn, log logger.Logger) repo.OrderICH {
	return orderRepo{
		db:  db,
		log: log,
	}
}

func (o orderRepo) Upsert(ctx context.Context, req *common.OrderCopyRequest) error {
	max_chunk_size := 50_000
	for len(req.GetItems()) > 0 {
		var chunk []*common.CreateOrderItemCopyRequest
		switch {
		case len(req.GetItems()) <= max_chunk_size:
			chunk = req.GetItems()[:] //[:len(documents)]
		default:
			chunk = req.GetItems()[:max_chunk_size]
		}

		batch, err := o.db.PrepareBatch(ctx, `INSERT INTO "order_item"`)
		if err != nil {
			return err
		}

		for _, item := range chunk {
			if err = batch.Append(
				uuid.New().String(),
				req.ShopId,
				req.OrderId,
				req.CompanyId,
				req.ShopName,
				item.ProductId,
				item.Value,
				item.SupplyPrice,
				item.SalePrice,
				item.Name,
				time.Now(),
			); err != nil {
				return err
			}
		}

		if err = batch.Send(); err != nil {
			return err
		}

		switch {
		case len(req.GetItems()) >= max_chunk_size:
			req.Items = req.Items[max_chunk_size:]
		default:
			req.Items = req.Items[len(req.Items):]
		}

		chunk = nil
	}

	return nil
}

func (o orderRepo) GetDashboardAnalytics(ctx context.Context, req *report_service.GetDashboardAnalyticsReq) (*report_service.GetDashboardAnalyticsRes, error) {
	var (
		filter        string
		options       string
		groupByPeriod string
		args          []interface{}
		tem           report_service.GetDashboardAnalyticsResponse
		res           models.GetDashboardAnalyticsResponse
		shops         []string
		top_products  []*report_service.TopProducts
		sale_analitcs []*report_service.SaleAnalitcs
	)
	

	query := `
	SELECT
		sum(sale_price * quantity) as gross_sale,
		sum(sale_price * quantity) as net_sale,
		sum((sale_price - supply_price) * quantity) as gross_profit
	FROM "order_item" WHERE company_id = @company_id
	`
	options += " GROUP BY company_id"
	args = append(args, clickhouse.Named("company_id", req.Request.CompanyId))

	if len(req.DateFrom) > 0 {
		if bDate, err := helper.ParseDate(req.DateFrom); err == nil {
			if req.DateFrom == req.DateTo {
				bDate = bDate.Add(-time.Hour * 24)
			}
			filter += ` AND created_at >= @begin_date`
			args = append(args, clickhouse.Named("begin_date", bDate.UTC().Format(config.DateTimeFormat)))
		} else if err != nil {
			return nil, err
		}
	}
	if len(req.DateTo) > 0 {
		if eDate, err := helper.ParseDate(req.DateTo); err == nil {
			filter += ` AND created_at <= @end_date`
			args = append(args, clickhouse.Named("end_date", eDate.UTC().Format(config.DateTimeFormat)))
		} else if err != nil {
			return nil, err
		}
	}

	err := o.db.QueryRow(ctx, query+filter+options, args...).Scan(
		&res.GrossSale,
		&res.NetSale,
		&res.GrossProfit)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return nil, err
		}
	}

	// Select top products
	queryTopProducts := `
	SELECT
		oi.product_name,
		oi.sale_price
	FROM "order_item" AS oi 
	WHERE company_id = @company_id 
	`
	options = `	GROUP BY oi.product_name, oi.sale_price
	ORDER BY count(product_name) DESC LIMIT 10`

	rows, err := o.db.Query(ctx, queryTopProducts+filter+options, args...)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return nil, err
		}
	}

	for rows.Next() {
		var topProduct report_service.TopProducts
		if err := rows.Scan(&topProduct.Name, &topProduct.Price); err != nil {
			return nil, err
		}
		top_products = append(top_products, &topProduct)
	}

	// Transactions
	queryTransactions := `
	SELECT
		count(product_id) as products
	FROM "order_item" WHERE company_id = @company_id 
	`
	options = " GROUP BY company_id"

	var (
		transactions  report_service.Transactions
		countProducts uint64
	)
	err = o.db.QueryRow(ctx, queryTransactions+filter+options, args...).Scan(
		&countProducts)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return nil, err
		}
	}
	transactions.Products = int32(countProducts)

	// Select sale graphics
	if len(req.Shops) > 0 {
		req.Shops = fmt.Sprintf("'%s'", req.Shops)
		req.Shops = strings.ReplaceAll(req.Shops, ",", "','")
		shops = strings.Split(req.Shops, ",")
		filter = fmt.Sprintf("%s AND (", filter)
		for i, v := range shops {
			if i == 0 {
				filter = fmt.Sprintf("%s shop_id = %s", filter, v)
			} else {
				filter = fmt.Sprintf("%s OR shop_id = %s", filter, v)
			}
		}
		filter = fmt.Sprintf("%s )", filter)
	}

	switch req.GetAnalyticsDateType() {
	case "hour":
		groupByPeriod = `CONCAT(toString(toDate(oi.created_at)), ' ', toString(toStartOfHour(oi.created_at)))`
		break
	case "week":
		groupByPeriod = `toString(toDateTime(toStartOfWeek(oi.created_at)))`
		break
	case "month":
		groupByPeriod = `toString(toStartOfMonth(toStartOfWeek(oi.created_at)))`
		break
	default:
		// for a day
		groupByPeriod = `toString(toDate(oi.created_at))`
	}

	querySaleGraph := `
	SELECT
		COALESCE(sum(oi.sale_price * oi.quantity), 0),
		` + groupByPeriod + `,
		oi.shop_name,
		oi.shop_id
	FROM "order_item" AS oi 
	WHERE company_id = @company_id 
	`
	options = `
	GROUP BY ` + groupByPeriod + `, oi.shop_name, oi.shop_id
	ORDER BY ` + groupByPeriod

	row1, err := o.db.Query(ctx, querySaleGraph+filter+options, args...)
	if err != nil {
		if err.Error() == "no rows in result set" {
			err = nil
		} else {
			return nil, err
		}
	}

	var price float64
	for row1.Next() {
		var saleA report_service.SaleAnalitcs
		var shop common.ShortShop
		if err := row1.Scan(&price, &saleA.Date, &shop.Name, &shop.Id); err != nil {
			return nil, err
		}
		saleA.Price = float32(price)
		saleA.Shop = &shop
		sale_analitcs = append(sale_analitcs, &saleA)
	}

	tem.GrossProfit = float32(res.GrossProfit)
	tem.GrossSale = float32(res.GrossSale)
	tem.NetSale = float32(res.NetSale)
	tem.TopProducts = top_products
	tem.ShopAnalytics = sale_analitcs
	tem.Transactions = &transactions
	return &report_service.GetDashboardAnalyticsRes{
		Response: &tem,
	}, nil
}
