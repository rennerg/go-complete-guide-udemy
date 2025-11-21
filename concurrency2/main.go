package main

import (
	"fmt"

	"example.com/concurrency2/channel"
	"example.com/concurrency2/mutex"
)

func main() {
	fmt.Println("Starting mutex examples...")
	mutex.Mutex()
	fmt.Println("Starting channel examples...")
	channel.Channel()
}
