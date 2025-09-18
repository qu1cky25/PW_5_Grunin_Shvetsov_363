package main

import (
	"fmt"
	"time"	
)

const (
	hp = iota + 1
	lp
)

func processing_task(hc, lc chan string) {
	for {
		select {
		case task := <-hc:
			fmt.Println("Задача с высоким приоритетом: ", task)
		case task := <-lc:
			fmt.Println("Задача с низким приоритетом: ", task)
		}
	}
}

func main() {
	var hc = make(chan string)
	var lc = make(chan string)
	go processing_task(hc, lc)
	hc <- "Важная задача"
	lc <- "Обычная задача"
	time.Sleep(time.Second)
} 