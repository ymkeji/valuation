package watermill

import (
	"context"

	"valuation/internal/conf"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
)

var RabbitConfig *conf.Rabbit

type Option func(*amqp.Config)

func SetExchange(name amqp.QueueNameGenerator, exchangeType string, args map[string]interface{}) Option {
	return func(c *amqp.Config) {
		c.Exchange.GenerateName = name
		c.Exchange.Type = exchangeType
		c.Exchange.Arguments = args
		c.Exchange.Durable = true
		c.Exchange.AutoDeleted = false
		c.Exchange.Internal = false
		c.Exchange.NoWait = false
	}
}

func SetQueue(name amqp.QueueNameGenerator, args map[string]interface{}) Option {
	return func(c *amqp.Config) {
		c.Queue.GenerateName = name
		c.Queue.Arguments = args
		c.Queue.Durable = true
		c.Queue.AutoDelete = false
		c.Queue.NoWait = false
		c.Queue.Exclusive = false
	}
}

func SetQueueBind(routingKey func(topic string) string, args map[string]interface{}) Option {
	return func(c *amqp.Config) {
		c.QueueBind.GenerateRoutingKey = routingKey
		c.QueueBind.Arguments = args
		c.QueueBind.NoWait = false
	}
}

func NewRabbitSub(topic string, opts ...Option) (<-chan *message.Message, error) {
	amqpConfig := amqp.NewDurableQueueConfig(RabbitConfig.AmqpURI)
	for _, opt := range opts {
		opt(&amqpConfig)
	}
	sub, err := amqp.NewSubscriber(
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		return nil, err
	}

	msg, err := sub.Subscribe(context.Background(), topic)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func NewRabbitPub(messages []*message.Message, topic string, opts ...Option) error {
	amqpConfig := amqp.NewDurableQueueConfig(RabbitConfig.AmqpURI)

	for _, opt := range opts {
		opt(&amqpConfig)
	}

	pub, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		return err
	}
	return pub.Publish(topic, messages...)
}
