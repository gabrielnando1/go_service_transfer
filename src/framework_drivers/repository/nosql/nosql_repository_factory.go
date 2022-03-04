package nosql

import (

	//entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	interface_nosql "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/repository/nosql/interface"
	mongodb "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/repository/nosql/mongodb"
)

type NoSqlRepositoryFactory struct {
	TransferRepository interface_nosql.ITransferRepository
}

func NoSqlRepository() *NoSqlRepositoryFactory {
	return &NoSqlRepositoryFactory{
		TransferRepository: mongodb.TransferRepositoryMongoDB(),
	}
}
