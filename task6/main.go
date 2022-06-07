package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	quitch := make(chan struct{})

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-quitch:
				wg.Done()
				fmt.Println("quitch!")
				cancel()
			case <-termChan:
				wg.Done()
				fmt.Println("termChan!")
				cancel()
			default:
				fmt.Println("I am working!")
				time.Sleep(time.Second)
			}
		}
	}()

	go func() {
		time.Sleep(3 * time.Second)
		quitch <- struct{}{}
	}()

	wg.Wait()
	fmt.Println("Program exit.")

}
