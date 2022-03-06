package dto_transfer_request

import (
	dto_common "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/common"
)

type TransferReceiveRequest struct {
	Amount         *float64             `json:"amount"`
	ExpirationDate *dto_common.JSONDATE `json:"expiration_date"`
	ExternalId     *string              `json:"external_id"`
	ExpectedOn     *dto_common.JSONDATE `json:"expected_on"`
	AccountFrom    *string              `json:"account_from"`
	AccountTarget  *string              `json:"account_target"`
}
