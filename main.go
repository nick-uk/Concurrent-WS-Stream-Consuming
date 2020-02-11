package main

import (
	"flag"
	"log"
	"runtime"
	"sync"
)

var wsresponse = make(chan string)
var procresult = make(chan string)
var wg sync.WaitGroup

// package init
func init() {
	// Assuming we have to do CPU costly calculations, we're creating dynamically one worker per CPU core for max stream processing performance.
	// Magically Go will distribute correctly our light go routines to the cores
	for worker := 1; worker <= runtime.NumCPU(); worker++ {
		wg.Add(1)
		go procResponse(worker, wsresponse, procresult, &wg)
	}
}

func main() {

	defer func() {
		log.Println("Stream is over & processing is done. Cheers")
	}()

	host := flag.String("host", "ws://127.0.0.1:12345", "Set destination WS host stream producer")

	// Create async WS session and let it go ...
	go func() {
		if err := wsClient(*host, wsresponse); err != nil {
			log.Println("Oops, WS error", err)
		}
	}()

	// Read infinitely results in async
	go readResults()

	// Now wait them to complete
	wg.Wait()
}
