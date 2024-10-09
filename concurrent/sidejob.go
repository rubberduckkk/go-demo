package main

import (
	"log"
	"time"
)

func SideJob(work func() interface{}) (done <-chan interface{}) {
	workDone := make(chan interface{})
	go func() {
		defer func() {
			if r := recover(); r != nil {
				workDone <- r
			}
			close(workDone)
		}()
		workDone <- work()
	}()
	return workDone
}

func main2() {
	sideDone := SideJob(func() interface{} {
		defer log.Printf("side job done\n")
		log.Printf("side job running\n")
		time.Sleep(time.Second * 5)
		return nil
	})

	log.Printf("demo is running\n")
	<-sideDone
	log.Printf("demo is done\n")
}
