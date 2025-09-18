package main

import (
	"fmt"
	"time"
)

func data_channel(ch chan string) {
	time.Sleep(1 * time.Second)
	//ch <- "bigbob" // данные
}

func main() {
	var ch = make(chan string)
	go data_channel(ch)
	for {
		select {
		case Data := <-ch:
			fmt.Println("Данные получены", Data)
			return
		default:
			fmt.Println("Данные потеряны")
			time.Sleep(500 * time.Millisecond)
		}
	}

}
