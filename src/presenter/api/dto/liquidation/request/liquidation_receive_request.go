package dto_liquidation_request

import (
	dto_common "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/common"
)

type LiquidationReceiveRequest struct {
	ExternalId *string              `json:"externalId"`
	Amount     *int64               `json:"amount"`
	ExpectedOn *dto_common.JSONDATE `json:"expectedOn"`
}
