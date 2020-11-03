package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Spawns a worker that adds values received from numsChan.
// It will increment the WaitGroup on start, and calls Done() when it
// detects the numsChan is closed.
func spawnWorker(id int, numsChan <-chan int, results chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	sum := 0
	count := 0
	for num := range numsChan {
		fmt.Printf("Worker.%d adding %d\n", id, num)

		// sleep a while to add some unevenness between the threads
		randSleepMillis := rand.Intn(10) * 100
		time.Sleep(time.Duration(randSleepMillis) * time.Millisecond)

		sum += num
		count++
	}
	fmt.Printf("Worker.%d sum=%d count=%d\n", id, sum, count)
	results <- sum
	wg.Done()
}

func WorkerPool() {

	const numWorkers = 5
	numsChan := make(chan int, numWorkers)
	results := make(chan int, numWorkers)
	var wg sync.WaitGroup

	// Create 5 workers
	for i := 0; i < numWorkers; i++ {
		go spawnWorker(i, numsChan, results, &wg)
	}

	// Send some numbers on the channel
	for i := 1; i <= numWorkers * 5; i++ {
		numsChan <- i
	}

	// signal the work is done
	close(numsChan)

	// wait for the threads to finish
	wg.Wait()

	total := 0
	for i := 0; i < numWorkers; i++ {
		total += <-results
	}

	fmt.Printf("Total=%d\n", total)
}
