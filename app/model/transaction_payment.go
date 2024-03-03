package model

import (
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
)

type TransactionPayment struct {
	Base
	DataOwner
	TransactionPaymentAPI
}

type TransactionPaymentAPI struct {
	TransactionID *uuid.UUID       `json:"transaction_id,omitempty" swaggertype:"string" format:"uuid"`
	PaidAmount    *float64         `json:"paid_amount,omitempty"`
	PaidAt        *strfmt.DateTime `json:"paid_at,omitempty" swaggertype:"string" format:"date-time"`
	Notes         *string          `json:"notes,omitempty"`
	Type          *string          `json:"type,omitempty"`
	Via           *string          `json:"via,omitempty"`
	ReferenceNo   *string          `json:"reference_no,omitempty"`
}
