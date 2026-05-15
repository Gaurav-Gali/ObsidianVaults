package main

import (
	"fmt"
	"sync"
	"time"
)

func countNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 5 {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func countLetters(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 'a'; i <= 'e'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go countNumbers(&wg)
	go countLetters(&wg)

	wg.Wait()

	fmt.Println("All go routines are done")
}
