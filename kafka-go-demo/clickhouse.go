package main

import (
	"database/sql"
	"time"

	"github.com/ClickHouse/clickhouse-go"
)

//trying to subscribe clickhouse on kafka topic
func kafkaReader(address string, topic string) {
	connect := connection()

	_, err := connect.Exec(`
		CREATE TABLE queue (
			topic		 String(),
			partition    Int32,
			offset		 Int64,
			key			 String(),
			value   	 String(),
			action_time  DateTime
		) ENGINE = Kafka SETTINGS kafka_broker_list = '` + address + `',
		kafka_topic_list = '` + topic + `',
		kafka_group_name = 'group1',
		kafka_format = 'JSONEachRow',
		kafka_num_consumers = 1;
	`)

	if err != nil {
		ErrorLogger.Fatal(err)
	}
}

func connection() *sql.DB {
	connect, err := sql.Open("clickhouse", "tcp://127.0.0.1:9000?debug=true") //todo env
	if err != nil {
		ErrorLogger.Fatal(err)
	}

	if err := connect.Ping(); err != nil {
		if exception, ok := err.(*clickhouse.Exception); ok {
			InfoLogger.Printf("[%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
		} else {
			InfoLogger.Println(err)
		}
		return nil //todo change behavior?
	}

	return connect
}

func initDB() {
	connect := connection()

	_, err := connect.Exec(`
		CREATE TABLE IF NOT EXISTS example (
			topic		 String(),
			partition    Int32,
			offset		 Int64,
			key			 String(),
			value   	 String(),
			action_time  DateTime
		) engine=Memory
	`)

	if err != nil {
		ErrorLogger.Fatal(err)
	}
}

func writeLog(topic string, partition int, offset int64, key []byte, UUID []byte) error {
	connect := connection()
	var (
		tx, _   = connect.Begin()
		stmt, _ = tx.Prepare("INSERT INTO example (topic, partition, offset, key, value, action_time) VALUES (?, ?, ?, ?, ?, ?)")
	)
	defer stmt.Close()

	if _, err := stmt.Exec(
		topic,
		partition,
		offset,
		string(key),
		string(UUID), //todo clickhouse works with UUID
		time.Now(),
	); err != nil {
		ErrorLogger.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		ErrorLogger.Fatal(err) //todo blocks commit
	}

	return nil
}
