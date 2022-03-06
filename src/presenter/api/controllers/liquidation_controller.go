package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	custom_errors "github.com/gabrielnando1/go_service_transfer/src/use_cases/custom_errors"

	dto_liquidation_request "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/liquidation/request"
	dto_liquidation_response "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/liquidation/response"
)

func liquidation_controller(router *gin.RouterGroup) {
	router.PATCH("/paymentOrders", LiquidationController_Receive)
}

// PingExample godoc
// @Summary transfer
// @Schemes
// @Description endpoint for  liquidation
// @Tags Liquidation
// @Accept json
// @Produce json
// @Param data body dto_liquidation_request.LiquidationReceiveRequest true "liquidation request"
// @success 200 {object} dto_liquidation_response.LiquidationReceiveResponse "desc"
// @Router /paymentOrders [patch]
func LiquidationController_Receive(c *gin.Context) {
	var request dto_liquidation_request.LiquidationReceiveRequest
	c.BindJSON(&request)
	if request.Amount == nil || *request.Amount <= 0 {
		Response(c, nil, &custom_errors.BadRequestError{
			Err: errors.New("Invalid data"),
		})
		return
	}

	var response *dto_liquidation_response.LiquidationReceiveResponse = &dto_liquidation_response.LiquidationReceiveResponse{
		InternalId: &[]string{uuid.New().String()}[0],
		Status:     &[]string{"APPROVED"}[0],
	}
	c.IndentedJSON(http.StatusOK, response)
}
