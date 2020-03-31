package email

import (
	"encoding/json"
	pb "github.com/AkezhanOb1/diplomaProject/api/proto/email"
	config "github.com/AkezhanOb1/diplomaProject/configs"

	"github.com/streadway/amqp"
	"log"
)

//BusinessRegistrationEmail is a func
func BusinessRegistrationEmail(email string) error{
	conn, err := amqp.Dial(config.RabbitConnection)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"businessRegistration", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Println(err)
		return err
	}

	 req := pb.BusinessOwnerEmailRequest{
	 	Email: &pb.BusinessOwnerEmail{
			Address:              email,
			XXX_NoUnkeyedLiteral: struct{}{},
			XXX_unrecognized:     nil,
			XXX_sizecache:        0,
		},
	 }

	 body, err := json.Marshal(req)
	 if err != nil {
	 	return nil
	 }

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}