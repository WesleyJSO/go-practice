package producer

import (
	"log"

	"github.com/IBM/sarama"
)

func newProducer() sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{"127.0.0.1:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to crate producer: %v", err)
	}
	return producer
}

func SendMsg(topic string, msg []byte) {
	go func() {
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(string(msg)),
		}
		var producer = newProducer()
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		} else {
			log.Printf("Producer message to partition %d at offset %d\n", partition, offset)
		}
		if err := producer.Close(); err != nil {
			log.Fatalf("Failed to close producer: %v", err)
		}
	}()
}
