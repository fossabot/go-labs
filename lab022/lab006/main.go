package main

import (
	"log"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	//select同时监听两个channel,一个是time.After
	//这样的话 要是超过After的时间，还没有收到c1，那么就执行time.After那个case
	select {
	case res := <-c1:
		log.Println("res=", res)
	case <-time.After(time.Second * 1):
		log.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		log.Println("res=", res)
	case <-time.After(time.Second * 3):
		log.Println("timeout 2")
	}
}
