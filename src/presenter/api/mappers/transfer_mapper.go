package mappers

import (
	"sync"
	"time"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	dto_transfer_request "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/transfer/request"
	dto_transfer_response "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/transfer/response"
)

type STransferMapper struct{}

var TransferMapper_instance *STransferMapper
var oTransferMapper sync.Once

func TransferMapper() *STransferMapper {
	oTransferMapper.Do(func() {
		TransferMapper_instance = &STransferMapper{}
	})
	return TransferMapper_instance
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func (STransferMapper) TransferRequestReceiveToEnty(request *dto_transfer_request.TransferReceiveRequest) *entity.TransferEntity {

	var obj *entity.TransferEntity
	if request != nil {
		obj = &entity.TransferEntity{
			ExternalId:    request.ExternalId,
			AccountFrom:   request.AccountFrom,
			AccountTarget: request.AccountTarget,
			//Amount:         request.Amount,
			ExpirationDate: &[]time.Time{request.ExpirationDate.ToTime()}[0],
			ExpectedOn:     &[]time.Time{request.ExpectedOn.ToTime()}[0],
		}
		if request.Amount != nil {
			obj.Amount = &[]int64{int64((*request.Amount) * 100)}[0]
		}
	}
	return obj
}

/////////////////////////////////////////////////////////////////
/////////////////////////////////////////////////////////////////
func (STransferMapper) TransferEntyToResponseReceive(ent *entity.TransferEntity) *dto_transfer_response.TransferReceiveResponse {

	var obj *dto_transfer_response.TransferReceiveResponse
	if ent != nil {
		obj = &dto_transfer_response.TransferReceiveResponse{
			Id:         &[]string{ent.ID.String()}[0],
			InternalId: ent.InternalId,
			Status:     ent.Status,
		}
	}
	return obj
}
