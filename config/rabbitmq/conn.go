package rabbitmq

import "os"

var (
	conn AmqpConnection
	err  error
)

func ConnectToRabbitMQ() (AmqpConnection, error) {
	conn, err = AmqpDialWrapper(os.Getenv("RABBIT_URI"))
	if err != nil {
		return nil, err
	}

	return conn, nil
}
