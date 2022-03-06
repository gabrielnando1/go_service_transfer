package controllers

import (
	"net/http"

	docs "github.com/gabrielnando1/go_service_transfer/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	dto_common_response "github.com/gabrielnando1/go_service_transfer/src/presenter/api/dto/common/response"
	custom_errors "github.com/gabrielnando1/go_service_transfer/src/use_cases/custom_errors"
)

var (
	Router    = gin.Default()
	RouterApi *gin.RouterGroup
)

func Initialize() {
	docs.SwaggerInfo.BasePath = ""
	RouterApi = Router.Group("")
	{
		RouterApi.Use(gin.Logger())
		RouterApi.Use(gin.Recovery())

		// Register Controllers
		transfer_controller(RouterApi)
		liquidation_controller(RouterApi)

	}
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func Response(c *gin.Context, obj interface{}, err error) {
	var response dto_common_response.GenericResponse
	response.Content = &obj
	if err != nil {
		response.Success = false
		var errror_message string = err.Error()
		response.ErrorMessage = &errror_message
		switch err.(type) {
		default:
			errror_message = "unexpected error"
			response.ErrorMessage = &errror_message
			c.IndentedJSON(http.StatusInternalServerError, response)
		case *custom_errors.BadRequestError:
			errror_message = err.Error()
			response.ErrorMessage = &errror_message
			c.IndentedJSON(http.StatusBadRequest, response)
		case *custom_errors.UnauthorizedRequestError:
			errror_message = err.Error()
			response.ErrorMessage = &errror_message
			c.IndentedJSON(http.StatusUnauthorized, response)
		}
	} else {
		response.Success = true
		c.IndentedJSON(http.StatusOK, response)
	}
}
