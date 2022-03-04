package repository

import (
	nosql "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/repository/nosql"
)

var (
	NoSQL *nosql.NoSqlRepositoryFactory = nosql.NoSqlRepository()
)
