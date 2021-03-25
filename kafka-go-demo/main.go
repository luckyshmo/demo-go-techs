package main

var address = "localhost:9092"
var topic = "topic"

func main() {
	createTopic()

	initDB()
	go mainProducer(address, topic)
	// go kafkaReader(address, topic) trying to subscribe clickhouse on kafka topic
	go mainConsumer(address, topic)
	InfoLogger.Println("Async call finished")
	for {
		//?
	}
}
