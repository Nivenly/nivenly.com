
package main

import (
	"time"
	"fmt"
)

// Will call the getChannel factory function, passing in different sleep times for each channel
func main() {
	ch_a := getChannel(1)
	ch_b := getChannel(2)
	ch_c := getChannel(5)
	switch {
	case <-ch_a:
		fmt.Println("Channel A")
		break
	case <-ch_b:
		fmt.Println("Channel B")
		break
	case <-ch_c:
		fmt.Println("Channel C")
		break
	}
}

// Will generate a new channel, and concurrently run a sleep based on the input
// Will return true after the sleep is over
func getChannel(N int) chan bool {
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * time.Duration(N))
		ch <- true
	}()
	return ch
}