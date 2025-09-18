package main

import (
	"fmt"
	"time"
	"sync"
)

var a sync.WaitGroup

func funcional(id int, done chan bool) {
	defer a.Done()
	for {
		select {
		case <- done:
			fmt.Println("Го-рутина", id, "завершила свою работу")
			return
		default :
			fmt.Println("Го-рутина", id, "продолжает свою работу...")
			time.Sleep(time.Second)
		}
	} 
}

func main() {
	var count = 3
	var done = make(chan bool)
	for i:= 0; i < count; i++ {
		a.Add(1)
		go funcional(i + 1, done)

	}
	time.Sleep(5 * time.Second)
	fmt.Println("Го-рутины остановились")

	close(done)
	a.Wait()
	fmt.Println("Го-рутины устали и завершили работу")
}
