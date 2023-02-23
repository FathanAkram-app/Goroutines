package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func say(s string, repeat int) {
	defer wg.Done()

	for ; 0 < repeat; repeat-- {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
	}

}

func main() {
	wg.Add(1)
	go say("Halo", 3)
	wg.Add(1)
	go say("Semua", 3)
	wg.Wait()

	wg.Add(1)
	go say("Apa kabar?", 1)
	wg.Wait()

}
