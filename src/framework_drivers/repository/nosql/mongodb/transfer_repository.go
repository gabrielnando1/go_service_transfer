package repository_mongodb

import (
	//"main/app/repository/entity"
	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	interface_nosql "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/repository/nosql/interface"
)

func TransferRepositoryMongoDB() interface_nosql.ITransferRepository {
	return STransferRepositoryMongoDB{}
}

type STransferRepositoryMongoDB struct {
}

func (STransferRepositoryMongoDB) Save(transfer *entity.TransferEntity) (*entity.TransferEntity, error) {
	if transfer.ID != nil && transfer.ID.Hex() != "000000000000000000000000" {
		err := Update(transfer)
		return transfer, err
	} else {
		result, err := Insert(transfer)
		if err == nil {
			transfer = result.(*entity.TransferEntity)
		}
		return transfer, err
	}
}
