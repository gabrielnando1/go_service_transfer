package controllers

import (
	"errors"

	"github.com/gin-gonic/gin"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
	bo "github.com/gabrielnando1/go_service_transfer/src/use_cases/bo"

	mappers "github.com/gabrielnando1/go_service_transfer/src/presenter/api/mappers"
	custom_errors "github.com/gabrielnando1/go_service_transfer/src/use_cases/custom_errors"

	dto_transfer_request "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/transfer/request"
	dto_transfer_response "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/transfer/response"
)

func transfer_controller(router *gin.RouterGroup) {
	router.POST("/api/transfer/receive", TransferController_Receive)
}

// PingExample godoc
// @Summary transfer
// @Schemes
// @Description endpoint for  transfer
// @Tags Transfer
// @Accept json
// @Produce json
// @Param data body dto_transfer_request.TransferReceiveRequest true "transfer request"
// @success 200 {object} dto_common_response.GenericResponse{content=dto_transfer_response.TransferReceiveResponse} "desc"
// @Router /api/transfer/receive [post]
func TransferController_Receive(c *gin.Context) {
	var request dto_transfer_request.TransferReceiveRequest
	c.BindJSON(&request)
	if request.ExpirationDate == nil || request.Amount == nil || *request.Amount <= 0 {
		Response(c, nil, &custom_errors.BadRequestError{
			Err: errors.New("Invalid data"),
		})
		return
	}
	var transfer *entity.TransferEntity = mappers.TransferMapper().TransferRequestReceiveToEnty(&request)
	var err error
	transfer, err = bo.TransferBO().Receive(transfer)
	if err != nil {
		Response(c, nil, err)
		return
	}

	var response *dto_transfer_response.TransferReceiveResponse = mappers.TransferMapper().TransferEntyToResponseReceive(transfer)
	Response(c, response, err)
}
