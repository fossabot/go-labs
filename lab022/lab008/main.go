package main

import (
	"log"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs //the more value will be false if jobs has been closed and all values in the channel have already been received
			if more {
				log.Println("received job", j)
			} else { //channel closed and not values in channel
				log.Println("received all jobs")
				time.Sleep(time.Second * 3)
				done <- true
				break
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		log.Println("sent job", j)
		time.Sleep(time.Second)
	}
	close(jobs)
	log.Println("sent all jobs", time.Now().Unix())

	<-done
	log.Println(time.Now().Unix())
}
