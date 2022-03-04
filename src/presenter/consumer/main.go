package consumer

import (
	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	controllers "github.com/gabrielnando1/go_service_transfer/src/presenter/api/controllers"

	services "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services"
	services_liquidation_queue "github.com/gabrielnando1/go_service_transfer/src/framework_drivers/services/liquidation_queue"

	bo "github.com/gabrielnando1/go_service_transfer/src/use_cases/bo"
)

var (
	router = controllers.Router
)

func QueueServiceConsumer() services_liquidation_queue.ILiquidationQueueServiceConsumer {
	return SLiquidationQueueServiceConsumer{}
}

type SLiquidationQueueServiceConsumer struct {
}

func (SLiquidationQueueServiceConsumer) SendToLiquidation(transfer *entity.TransferEntity) error {
	return bo.TransferBO().SendToLiquidation(transfer)
}

func StartConsumer() {
	go func() {
		services.Service().LiquidationQueue.QueuePaymentOrdersSendConsumer(QueueServiceConsumer())
	}()

}
