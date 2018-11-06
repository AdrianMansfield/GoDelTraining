package main

import (
	"fmt"
	"sync"
)

func greaterWorld() func() {
	return func() {
		defer fmt.Println("World")
		fmt.Println("Hello")
	}
}

func main() {
	var once sync.Once
	onceBody := greaterWorld()
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}
	<-done
}
