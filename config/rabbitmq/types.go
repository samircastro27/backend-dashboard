package rabbitmq

import "github.com/streadway/amqp"

type AmqpChannel interface {
	QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error)
	Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error)
	Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error
	Close() error
}

type AmqpConnection interface {
	Channel() (AmqpChannel, error)
	Ping() error
}

type AmqpDial func(url string) (AmqpConnection, error)

func AmqpDialWrapper(uri string) (AmqpConnection, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	return AmqpConnectionWrapper{conn}, nil
}

// This is a wrapper for the amqp.Connection type that implements the AmqpConnection interface
type AmqpConnectionWrapper struct {
	conn *amqp.Connection
}

func (w AmqpConnectionWrapper) Channel() (AmqpChannel, error) {
	return w.conn.Channel()
}

func (w AmqpConnectionWrapper) Ping() error {
	ch, err := w.conn.Channel()
	if err != nil {
		return err
	}
	return ch.Close()
}

func (w AmqpConnectionWrapper) Close() error {
	return w.conn.Close()
}
