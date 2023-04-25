package models

import (
	"genproto/common"
)

type CreateOrderItemReq struct {
	Id          string  `json:"id"`
	OrderId     string  `json:"order_id"`
	CompanyId   string  `json:"company_id"`
	ShopId      string  `json:"shop_id"`
	ShopName    string  `json:"shop_name"`
	ProductId   string  `json:"product_id"`
	Quantity    float32 `json:"quantity"`
	SupplyPrice float32 `json:"supply_price"`
	SalePrice   float32 `json:"sale_price"`
	ProductName string  `json:"product_name"`
	CreatedAt   string  `json:"created_at"`
}

type GetDashboardAnalyticsResponse struct {
	GrossSale     float64         `json:"gross_sale"`
	Refund        float64         `json:"refund"`
	NetSale       float64         `json:"net_sale"`
	GrossProfit   float64         `json:"gross_profit"`
	TopProducts   []*TopProducts  `json:"top_products"`
	Payments      []*Payments     `json:"payments"`
	ShopAnalytics []*SaleAnalitcs `json:"shop_analytics"`
}

type TopProducts struct {
	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id    string  `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Price float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}

type Payments struct {
	Name  string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}
type SaleAnalitcs struct {
	Shop  *common.ShortShop `protobuf:"bytes,1,opt,name=shop,proto3" json:"shop,omitempty"`
	Date  string            `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Price float32           `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
}
