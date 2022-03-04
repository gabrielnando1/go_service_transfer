package liquidation_api_service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	api_request "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_api/request"
	api_response "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_api/response"
)

var (
	HOST = "http://localhost"
)

type ILiquidationApiService interface {
	PaymentOrdersSend(transfer *entity.TransferEntity) (*entity.TransferEntity, error)
	//SendVerifyAccountUserMail(email *string, name *string, accountName *string, token *string) error
}

func LiquidationApiService() ILiquidationApiService {
	return SLiquidationApiService{}
}

type SLiquidationApiService struct {
}

func (SLiquidationApiService) PaymentOrdersSend(transfer *entity.TransferEntity) (*entity.TransferEntity, error) {
	payload, err := json.Marshal(api_request.LiquidationApiPaymentOrderSendRequest{
		ExternalId: transfer.ExternalId,
		Amount:     transfer.Amount,
		ExpectedOn: transfer.ExpectedOn,
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	client := &http.Client{}
	url := fmt.Sprintf("%s%s", HOST, "/paymentOrders")
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var response *api_response.LiquidationApiPaymentOrderSendRespnse = &api_response.LiquidationApiPaymentOrderSendRespnse{}
		json.NewDecoder(resp.Body).Decode(response)

		transfer.InternalId = response.InternalId
		transfer.Status = response.Status

		return transfer, err
	}
	return nil, errors.New("unexpected error")
}
