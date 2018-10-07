package consumer

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

//var count int64

// NewRateCounter initiates a new counter
func NewRateCounter(window time.Duration, countChannel chan bool) {

	var count int64
	deviceid := viper.GetString("device.id")
	// Function generates ticks at the end of the configured window
	go func() {
		for {
			time.Sleep(window)
			recordCount(window, &count, deviceid)
		}
	}()

	for range countChannel {
		count++
	}

}

//func increment() {
//	log.Printf("Increment Count: %d", count)
//	count
//}

func recordCount(window time.Duration, count *int64, deviceid string) {
	//	log.Printf("Receive Count: %d", *count)
	fmt.Printf("%s | %s Receive Count: %d\n", time.Now().Format(time.RFC3339), deviceid, *count)
	*count = 0

}
