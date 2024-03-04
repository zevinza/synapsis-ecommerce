package model

import "github.com/google/uuid"

type Transaction struct {
	Base
	DataOwner
	TransactionAPI
	Payments *[]TransactionPayment `json:"payments,omitempty"`
	Details  *[]TransactionDetail  `json:"details,omitempty"`
}

type TransactionAPI struct {
	UserID            *uuid.UUID `json:"user_id,omitempty" swaggertype:"string" format:"uuid"`
	InvoiceNo         *string    `json:"invoice_no,omitempty"`
	TransactionStatus *string    `json:"transaction_status,omitempty"`
	Description       *string    `json:"description,omitempty"`
	Notes             *string    `json:"notes,omitempty"`
	TotalProductPrice *float64   `json:"total_productPrice,omitempty"`
	TotalDiscount     *float64   `json:"total_discount,omitempty"`
	TotalFee          *float64   `json:"total_fee,omitempty"`
	TotalPrice        *float64   `json:"total_price,omitempty"`
	ContactName       *string    `json:"contact_name,omitempty"`
	ContactDetail     *string    `json:"contact_detail,omitempty"`
}

// type TransactionPayload struct {
// 	Data    *TransactionAPI        `json:"data,omitempty"`
// 	Payment *TransactionPaymentAPI `json:"payment,omitempty"`
// 	CartIds *[]uuid.UUID           `json:"cart_ids,omitempty"`
// }

type TransactionPayload struct {
	CartIds     *[]uuid.UUID `json:"cart_ids,omitempty"`
	Notes       *string      `json:"notes,omitempty"`
	InvoiceNo   *string      `json:"invoice_no,omitempty"`
	Description *string      `json:"description,omitempty"`
	Discount    *float64     `json:"discount,omitempty"`
	Fee         *float64     `json:"fee,omitempty"`
}
