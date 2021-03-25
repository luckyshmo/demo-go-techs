package main

import "log"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	go mainSend()
	go mainReceive()
	println("Async call finished")
	for {

	}
}
