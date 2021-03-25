package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

func mainProducer(address string, topic string) {

	writer := newKafkaWriter(address, topic)
	defer writer.Close()
	InfoLogger.Println("start producing ... !!")
	for i := 0; ; i++ {
		msg := kafka.Message{
			Key:   []byte(fmt.Sprintf("Key-%d", i)),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			WarningLogger.Println(err)
		}
		time.Sleep(1 * time.Second)
	}
}

func createTopic() {
	// to create topics when auto.create.topics.enable='true'
	_, err := kafka.DialLeader(context.Background(), "tcp", address, topic, 0)
	if err != nil {
		ErrorLogger.Fatal(err)
	}
}
