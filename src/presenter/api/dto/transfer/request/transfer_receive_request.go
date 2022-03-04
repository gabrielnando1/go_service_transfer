package dto_transfer_request

import "time"

type TransferReceiveRequest struct {
	Amount         *float64   `json:"amount"`
	ExpirationDate *time.Time `json:"expiration_date"`
	ExternalId     *string    `json:"external_id"`
	ExpectedOn     *time.Time `json:"expected_on"`
	AccountFrom    *string    `json:"account_from"`
	AccountTarget  *string    `json:"account_target"`
}
