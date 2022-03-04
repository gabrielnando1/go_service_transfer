package bo_test

import (
	"testing"
	"time"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	bo "github.com/gabrielnando1/go_service_transfer/src/use_cases/bo"
	custom_errors "github.com/gabrielnando1/go_service_transfer/src/use_cases/custom_errors"
)

func TestTransferExpired(t *testing.T) {
	yesterday := time.Now().AddDate(0, 0, -1)
	transfer := &entity.TransferEntity{
		AccountFrom:    &[]string{"teste"}[0],
		AccountTarget:  &[]string{"teste"}[0],
		Amount:         &[]int64{1}[0],
		ExpirationDate: &yesterday,
	}
	transfer, err := bo.TransferBO().Receive(transfer)

	if err == nil {
		t.Errorf("validation expiration date not working.")
	}
	switch err.(type) {
	default:
		t.Errorf("validation expiration date not working.")
	case *custom_errors.BadRequestError:
		error_message := err.Error()
		if error_message != "Invalid data Transfer: Date expirated" {
			t.Errorf("validation expiration date not working.")
		}
	}
}
