package main

import (
	"fmt"
	"time"
	
)

func Request(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Ответ от сервера"
}

func main() {
	var ch  = make(chan string)
	go Request(ch) 
	select {
		case issue := <-ch:
			fmt.Println("Полученный ответ", issue)
		case <- time.After(2 * time.Second):
			fmt.Println("Тайм-аут")
	}
} 