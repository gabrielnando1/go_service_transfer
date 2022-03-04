package main

import (
	api "github.com/gabrielnando1/go_service_transfer/src/presenter/api"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	api.StartApi()
}
