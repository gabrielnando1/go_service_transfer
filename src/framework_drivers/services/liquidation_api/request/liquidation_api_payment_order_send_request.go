package service_liquidation_api_request

import "time"

type LiquidationApiPaymentOrderSendRequest struct {
	ExternalId *string    `json:"externalId"`
	Amount     *int64     `json:"amount"`
	ExpectedOn *time.Time `json:"expectedOn"`
}
