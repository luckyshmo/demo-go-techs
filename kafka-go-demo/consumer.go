package main

import (
	"context"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

func mainConsumer(address string, topic string) {
	groupID := "0" // ???? os.Getenv("groupID")

	reader := getKafkaReader(address, topic, groupID)

	defer reader.Close()

	InfoLogger.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			ErrorLogger.Fatalln(err)
		}
		InfoLogger.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		err = writeLog(m.Topic, m.Partition, m.Offset, m.Key, m.Value)
		if err != nil {
			ErrorLogger.Println(err)
		}
	}
}
