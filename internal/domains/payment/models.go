package payment

import "time"

// Payment defines set of fields for payment entity in system
type Payment struct {
	Id         int        `json:"id" gorm:"primaryKey"`
	UserId     int        `json:"user_id"`
	BasketId   int        `json:"basket_id"`
	AddressId  int        `json:"address_id"`
	DiscountId int        `json:"discount_id"`
	GatewayId  int        `json:"gateway_id"`
	PostTypeId int        `json:"post_type_id"`
	TotalPrice float64    `json:"total_price"`
	RefNum     string     `json:"ref_num" gorm:"uniqueIndex"`
	TraceNum   string     `json:"trace_num" gorm:"uniqueIndex"`
	Status     string     `json:"status"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdateAt   *time.Time `json:"update_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}

// CreatePaymentRequest defines set of methods for payment creation
type CreatePaymentRequest struct {
	BasketId   int     `json:"basket_id" validate:"required,min=1"`
	AddressId  int     `json:"address_id" validate:"required,min=1"`
	DiscountId int     `json:"discount_id" validate:"required,min=1"`
	GatewayId  int     `json:"gateway_id" validate:"required,min=1"`
	PostTypeId int     `json:"post_type_id" validate:"required,min=1"`
	TotalPrice float64 `json:"total_price" validate:"required,min=1"`
}

// PaymentVerifyRequest defines set of fields that are required for payment verification
type PaymentVerifyRequest struct {
	PaymentId int    `query:"payment_id" validate:"required,min=1"`
	Authority string `query:"authority" validate:"required,min=36,max=36"`
}

// RequestPaymentResponse defines set of fields in which is returned by payment gateway before starting payment operation and redirecting to bank
type RequestPaymentResponse struct {
	Url             string
	Key             string
	OperationStatus int
}
