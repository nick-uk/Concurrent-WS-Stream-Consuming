package main

import (
	"log"
	"sync"
	"time"
)

// My worker
func procResponse(id int, wsresponse <-chan string, procresult chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("Worker", id, "is pending...")
	for job := range wsresponse {
		log.Println("Fake processing in progress for", job, "which takes 1 sec & 780 ms from worker", id)
		time.Sleep(time.Second*1 + time.Millisecond*780)
		procresult <- genRandomStr()
	}
}

func readResults() {
	counter := 0
	for {
		counter++
		res := <-procresult
		log.Println("Result:", res, counter)
	}
}
