package service_liquidation_api_request

import (
	dto_common "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/common"
)

type LiquidationApiPaymentOrderSendRequest struct {
	ExternalId *string              `json:"externalId"`
	Amount     *int64               `json:"amount"`
	ExpectedOn *dto_common.JSONDATE `json:"expectedOn"`
}
