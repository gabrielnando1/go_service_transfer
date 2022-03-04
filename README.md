* Alterar as configurações de conexao com banco mongoDB no arquivo:
    - /go_service_transfer/src/framework_drivers/repository/nosql/mongodb/base_repository.go
        variaveis :     username = "username"
                        password = "password"
                        host     = "host"
                        database = "database"
              


* Alterar a varivel HOST para endereço da api de liquidação no arquivo:
    - /go_service_transfer/src/framework_drivers/services/liquidation_api/liquidation_api_service.go


* Alterar a varivel HOST para endereço RabbitMQ no arquivo: 
    - /go_service_transfer/src/framework_drivers/services/liquidation_queue/liquidation_queue.go



* Execução do arquivo de test executar comando:    
    - go test github.com/gabrielnando1/go_service_transfer/tests/use_cases/bo



* Ao rodar api estara respondendo na porta 8080
    - http://localhost:8000/swagger/index.html
