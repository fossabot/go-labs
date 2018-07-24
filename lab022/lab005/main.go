package main

import (
	"log"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(time.Second * 5)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		log.Println("i=", i)
		select {
		case msg1 := <-c1:
			log.Println("received", msg1)
		case msg2 := <-c2:
			log.Println("received", msg2)

			//要是有default，select就不会去等待了，会直接去运行default
			//default:
			//	log.Println("default")
		}
	}
}
