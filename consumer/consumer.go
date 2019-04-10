package consumer

import (
	"fmt"
	"sync"
	"time"

	RATECOUNTER "github.com/paulbellamy/ratecounter"
	"github.com/spf13/viper"
)

var wg = sync.WaitGroup{}

// RunConsumer reads data from the queue
func RunConsumer() {

	fmt.Println("Running Consumer...")
	var amqpConnDetails AMQPConnDetailsType

	amqpConnDetails.host = viper.GetString("edge-feed.cloud-rabbitmq.host")
	amqpConnDetails.port = viper.GetString("edge-feed.cloud-rabbitmq.port")
	amqpConnDetails.user = viper.GetString("edge-feed.cloud-rabbitmq.user")
	amqpConnDetails.password = viper.GetString("edge-feed.cloud-rabbitmq.password")
	amqpConnDetails.queue = viper.GetString("edge-feed.cloud-rabbitmq.queue")

	clientid := viper.GetString("device.id")
	deliveries := make(chan string, 409600)
	dataRateReadSeconds := viper.GetInt64("edge-feed.cloud-rabbitmq.dataratereadseconds")

	counter := RATECOUNTER.NewRateCounter(1 * time.Second)

	wg.Add(1)

	go NewConsumer(amqpConnDetails, &deliveries)

	go handleDeliveries(&deliveries, counter)

	//Go routine to print out data sending rate
	go func() {
		for {
			fmt.Printf("%s | Data receive rate at '%s' : %d \t records/sec\n", time.Now().Format(time.RFC3339), clientid, counter.Rate())
			time.Sleep(time.Second * time.Duration(dataRateReadSeconds))
		}
	}()

	wg.Wait()

}

func handleDeliveries(deliveries *chan string, counter *RATECOUNTER.RateCounter) {
	//	doCount := make(chan bool)
	//	go NewRateCounter(time.Second*10, doCount)
	for range *deliveries {
		//log.Printf("Received message: %s", msg)
		//	increment()
		//log.Printf("Incementing Counter: %s", msg)
		//		doCount <- true
		counter.Incr(1)
	}

}
