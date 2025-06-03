package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numWorkers  = 3
	numMessages = 10
)

type Job struct {
	Message string
}

func main() {
	rand.Seed(time.Now().UnixNano())

	jobs := make(chan Job)
	var wg sync.WaitGroup

	// Start fixed number of workers
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for i := 1; i <= numMessages; i++ {
		job := Job{Message: fmt.Sprintf("Message %d", i)}
		wg.Add(1)
		jobs <- job
		time.Sleep(100 * time.Millisecond) // simulate message arrival
	}

	close(jobs) // No more jobs
	wg.Wait()   // Wait for all jobs to finish
}

func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	for job := range jobs {
		processJob(job, id)
		wg.Done()
	}
}

func processJob(job Job, workerID int) {
	delay := time.Duration(500+rand.Intn(1000)) * time.Millisecond // 500ms to 1500ms
	fmt.Printf("[Worker %d] Processing %s (will take %v)\n", workerID, job.Message, delay)
	time.Sleep(delay)
	fmt.Printf("[Worker %d] Finished: %s\n", workerID, job.Message)
}
