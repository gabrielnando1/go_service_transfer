package bo

import (
	"errors"
	"sync"
	"time"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	repository "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/repository"
	services "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services"
	custom_errors "github.com/gabrielnando1/go_service_transfer/src/use_cases/custom_errors"
)

type STransferBO struct{}

var TransferBO_instance *STransferBO
var oTransferBO sync.Once

func TransferBO() *STransferBO {
	oTransferBO.Do(func() {
		TransferBO_instance = &STransferBO{}
	})
	return TransferBO_instance
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func (STransferBO) Receive(transfer *entity.TransferEntity) (*entity.TransferEntity, error) {
	if transfer == nil || transfer.AccountFrom == nil || transfer.AccountTarget == nil || transfer.ExpirationDate == nil || transfer.Amount == nil {
		return nil, &custom_errors.BadRequestError{
			Err: errors.New("Invalid data Transfer"),
		}
	}

	if *transfer.Amount == 0 {
		return nil, &custom_errors.BadRequestError{
			Err: errors.New("Amount can't be zero (0)"),
		}
	}

	if (*transfer.ExpirationDate).Before(time.Now()) {
		return nil, &custom_errors.BadRequestError{
			Err: errors.New("Invalid data Transfer: Date expirated"),
		}
	}

	transfer.Status = &[]string{"PROCESSING"}[0]
	var err error
	transfer, err = repository.NoSQL.TransferRepository.Save(transfer)
	if err != nil {
		return nil, errors.New("unexpected error")
	}

	/*
		transfer, err = services.Service().LiquidationApi.PaymentOrdersSend(transfer)
	*/
	transfer, err = services.Service().LiquidationQueue.QueuePaymentOrdersSendPublisher(transfer)
	if err != nil {
		return nil, err
	}

	return transfer, nil
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func (STransferBO) SendToLiquidation(transfer *entity.TransferEntity) error {
	if transfer == nil || transfer.AccountFrom == nil || transfer.AccountTarget == nil || transfer.ExpirationDate == nil || transfer.Amount == nil {
		return &custom_errors.BadRequestError{
			Err: errors.New("Invalid data Transfer"),
		}
	}

	if (*transfer.ExpirationDate).Before(time.Now()) {
		return &custom_errors.BadRequestError{
			Err: errors.New("Invalid data Transfer: Date expirated"),
		}
	}

	var err error
	transfer, err = services.Service().LiquidationApi.PaymentOrdersSend(transfer)
	if err != nil {
		return errors.New("unexpected error")
	}
	transfer, err = repository.NoSQL.TransferRepository.Save(transfer)
	if err != nil {
		return err
	}

	return nil
}
