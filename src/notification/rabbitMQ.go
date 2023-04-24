package notification

import (
	"api/src/config"

	"github.com/streadway/amqp"
)

func SendStockMessage(message []byte) error {
	ConnectRabbit(message, config.RabbitStockRoutingKey)
	return nil
}

func SendFiatMessage(message []byte) error {
	ConnectRabbit(message, config.RabbitFiatRountingKey)
	return nil
}

func ConnectRabbit(message []byte, routingKey string) error {

	// Conecta com o RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	// Cria um canal
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer ch.Close()

	// Configura a exchange
	err = ch.ExchangeDeclare(
		config.RabbitExchange, // exchange name
		"direct",              // exchange type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         message,         //[]byte(message),
		DeliveryMode: amqp.Persistent, // make message persistent
	}

	err = ch.Publish(
		config.RabbitExchange, // exchange
		routingKey,            // routing key
		false,                 // mandatory
		false,                 // immediate
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}
