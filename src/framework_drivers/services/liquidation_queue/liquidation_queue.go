package liquidation_queue_service

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	entity "github.com/gabrielnando1/go_service_transfer/src/entity"
)

var (
	HOST       = "amqp://user:pass@localhost:port"
	QUEUE_NAME = "liquidation_queue"
)

type ILiquidationQueueService interface {
	QueuePaymentOrdersSendPublisher(transfer *entity.TransferEntity) (*entity.TransferEntity, error)
	QueuePaymentOrdersSendConsumer(consumer ILiquidationQueueServiceConsumer) error
	//SendVerifyAccountUserMail(email *string, name *string, accountName *string, token *string) error
}

func LiquidationQueueService() ILiquidationQueueService {
	return SLiquidationQueueService{}
}

type SLiquidationQueueService struct {
}

func (SLiquidationQueueService) QueuePaymentOrdersSendPublisher(transfer *entity.TransferEntity) (*entity.TransferEntity, error) {
	conn, err := amqp.Dial(HOST)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer ch.Close()
	q, err := ch.QueueDeclare(
		QUEUE_NAME,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	fmt.Println(q)

	payload, err := json.Marshal(*transfer)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	err = ch.Publish(
		"",
		QUEUE_NAME,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return transfer, nil
}

type ILiquidationQueueServiceConsumer interface {
	SendToLiquidation(transfer *entity.TransferEntity) error
}

func (SLiquidationQueueService) QueuePaymentOrdersSendConsumer(consumer ILiquidationQueueServiceConsumer) error {
	conn, err := amqp.Dial(HOST)
	if err != nil {
		fmt.Println("Failed Initializing Broker Connection")
		return err
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	msgs, err := ch.Consume(
		QUEUE_NAME,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved Message: %s\n", d.Body)
			var obj entity.TransferEntity
			if err := json.Unmarshal(d.Body, &obj); err != nil {
				d.Nack(false, true)
				fmt.Println(err)
			}
			err = consumer.SendToLiquidation(&obj)
			if err != nil {
				d.Nack(false, true)
				fmt.Println(err)
			}
			if err == nil {
				d.Ack(false)
				fmt.Printf("Success Consume Message: %s\n", d.Body)
			}
		}
	}()

	fmt.Println("Successfully Connected to our RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever

	return nil
}
