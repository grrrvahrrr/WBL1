package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := workerPullMutex(ctx, 1000)
	if err != nil {
		log.Println(err)
	}
	<-ctx.Done()
}

func workerPullMutex(ctx context.Context, operations int32) error {
	var workers = make(chan struct{}, 10)
	var number number
	var wg sync.WaitGroup

	wg.Add(int(operations))
	for i := 0; i < int(operations); i++ {
		workers <- struct{}{}
		go func() {
			defer func() {
				<-workers
				wg.Done()
			}()
			number.addOne()
		}()
	}
	wg.Wait()
	log.Println("Final result is", number.num)
	return ctx.Err()
}

type number struct {
	num int
	mu  sync.Mutex
}

func (n *number) addOne() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.num++
}
