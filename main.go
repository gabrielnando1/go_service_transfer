package main

import (
	api "github.com/gabrielnando1/go_service_transfer/src/presenter/api"
	consumer "github.com/gabrielnando1/go_service_transfer/src/presenter/consumer"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	consumer.StartConsumer()
	api.StartApi()
}
