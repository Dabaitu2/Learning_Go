package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func waitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err != nil {
			return nil
		}
		log.Printf("server not responding (%s); retring...", err)
		time.Sleep(time.Second << uint(tries)) //指数增长
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
