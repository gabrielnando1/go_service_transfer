package interface_nosql

import (
	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
)

type ITransferRepository interface {
	Save(transfer *entity.TransferEntity) (*entity.TransferEntity, error)
}
