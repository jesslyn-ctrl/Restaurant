package model

type OrderStatus string

type Order struct {
	Id            string         `gorm:"primaryKey" json:"id"`
	Status        OrderStatus    `json:"status"`
	ProductOrders []ProductOrder `json:"product_orders"`
	TotalPrice    int64          `json:"total_price"`
	ReferenceId   string         `gorm:"unique" json:"reference_id"`
}

type ProductOrderStatus string

type ProductOrder struct {
	Id        string             `gorm:"primaryKey" json:"id"`
	OrderId   string             `json:"order_id"`
	OrderCode string             `json:"order_code"`
	Quantity  int                `json:"quantity"`
	Price     int64              `json:"price"`
	Status    ProductOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceId   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	OrderId string `json:"order_id"`
}
