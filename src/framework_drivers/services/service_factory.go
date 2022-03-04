package services

import (

	//entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	service_liquidation_api "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_api"
	service_liquidation_queue "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_queue"
)

type ServiceFactory struct {
	LiquidationApi   service_liquidation_api.ILiquidationApiService
	LiquidationQueue service_liquidation_queue.ILiquidationQueueService
}

func Service() *ServiceFactory {
	return &ServiceFactory{
		LiquidationApi:   service_liquidation_api.LiquidationApiService(),
		LiquidationQueue: service_liquidation_queue.LiquidationQueueService(),
	}
}
