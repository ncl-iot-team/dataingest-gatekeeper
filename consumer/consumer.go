package consumer

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var wg = sync.WaitGroup{}

// RunConsumer reads data from the queue
func RunConsumer() {

	log.Printf("Running Consumer...")
	var amqpConnDetails AMQPConnDetailsType

	amqpConnDetails.host = viper.GetString("edge-feed.cloud-rabbitmq.host")
	amqpConnDetails.port = viper.GetString("edge-feed.cloud-rabbitmq.port")
	amqpConnDetails.user = viper.GetString("edge-feed.cloud-rabbitmq.user")
	amqpConnDetails.password = viper.GetString("edge-feed.cloud-rabbitmq.password")
	amqpConnDetails.queue = viper.GetString("edge-feed.cloud-rabbitmq.queue")

	deliveries := make(chan string)

	wg.Add(1)

	go NewConsumer(amqpConnDetails, deliveries)

	go handleDeliveries(deliveries)

	wg.Wait()

}

func handleDeliveries(deliveries chan string) {
	doCount := make(chan bool)
	go NewRateCounter(time.Second*10, doCount)
	for range deliveries {
		//log.Printf("Received message: %s", msg)
		//	increment()
		//log.Printf("Incementing Counter: %s", msg)
		doCount <- true
	}

}
