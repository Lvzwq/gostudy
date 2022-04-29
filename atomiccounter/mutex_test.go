package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(*testing.T) {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100000000; j++ {
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ops)
}

func TestMutex(*testing.T) {
	var ops uint64
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100000000; j++ {
				inc(&mu, &ops)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ops)
}

func inc(mu *sync.Mutex, ops *uint64) {
	mu.Lock()
	defer mu.Unlock()
	*ops++
}
