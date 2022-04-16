package kafka

import (
	"log"
	"github.com/Shopify/sarama"
)
type PusherKafka struct {
	Addr []string // kafka addr
}

func NewKafkaPusher(addr []string) *PusherKafka{
	return &PusherKafka{
		Addr: addr,
	}
}

func (pusher *PusherKafka) Push() error{
	var err error
	producer := pusher.getProducerProducer([]{"localhost:9092"})
	producer.
	return err
}

func (pusher *PusherKafka) getProducerProducer(addr []string) sarama.AsyncProducer{
	config := sarama.NewConfig()
	// config.Producer.Return.Successes = true
	client, err := sarama.NewClient(addr, config)
	if err != nil {
		log.Fatalf("unable to create kafka client: %q", err)
	}

	producer, err := sarama.NewAsyncProducerFromClient(client)
	if err != nil {
		log.Fatalf("unable to create kafka producer: %q", err)
	}
	defer producer.Close()
	return producer
}
