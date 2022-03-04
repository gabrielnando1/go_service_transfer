package services

import (

	//entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	service_liquidation_api "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_api"
)

type ServiceFactory struct {
	LiquidationApi service_liquidation_api.ILiquidationApiService
}

func Service() *ServiceFactory {
	return &ServiceFactory{
		LiquidationApi: service_liquidation_api.LiquidationApiService(),
	}
}
