package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var numOfWorkers int
	mainch := make(chan struct{})
	var wg sync.WaitGroup

	fmt.Println("Please, Enter number of workers.")

	_, err := fmt.Scan(&numOfWorkers)
	if err != nil {
		return
	}

	wg.Add(numOfWorkers)
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-termChan:
				cancel()
			default:
				mainch <- struct{}{}
			}

		}
	}()

	for i := 0; i < numOfWorkers; i++ {
		go func(i int) {
			for {
				select {
				case <-mainch:
					fmt.Printf("%d got Data: %s\n", i, <-mainch)
				case <-ctx.Done():
					fmt.Printf("%d quiting\n", i)
					wg.Done()
					return
				}
			}
		}(i)
	}

	wg.Wait()
	close(mainch)
	close(termChan)
	fmt.Println("Done!")
}
