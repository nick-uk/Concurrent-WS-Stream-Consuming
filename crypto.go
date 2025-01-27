package main

import (
	"crypto/rand"
	"fmt"
	"log"
)

// This is for demo only. There are better random crypto solutions of-course
func genRandomStr() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
