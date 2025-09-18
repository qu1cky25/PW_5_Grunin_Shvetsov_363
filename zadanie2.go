package main

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"sync"	
)
var a sync.WaitGroup
func Loop(done chan struct{}) {
	defer a.Done()
	for {
		select {
		case <- done:
			return
		default:
			fmt.Println("Продолжение работы....")
			time.Sleep(time.Second)
		}
	}
} 

func main() {
	var done = make(chan struct{})
	a.Add(1)
	
	go Loop(done)
	var signals = make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	<- signals
	close(done)

	a.Wait()
} 