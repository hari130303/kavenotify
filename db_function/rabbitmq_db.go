package dbfunction

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
)

var RabbitChannel *amqp.Channel

var Queue1 amqp.Queue

func RabbitQueue() error {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return err
	}
	// defer conn.Close()
	RabbitChannel, err = conn.Channel()
	if err != nil {
		return err
	}
	// defer ch.Close()

	Queue1, err = RabbitChannel.QueueDeclare(
		"loginqueue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	return err
}

func RabbitReciever() {

	msgs, err := RabbitChannel.Consume(
		Queue1.Name, // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		fmt.Println("rabbit mq reciever not listening : ", err)
	}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			Data := bson.M{"user_name": string(d.Body), "entry_type": "LOGIN", "entry_time": time.Now()}
			insertManyResult, err := lmdb.InsertOne(context.TODO(), Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Inserted document with IDs:", insertManyResult.InsertedID)
		}
	}()

	select {}
}
