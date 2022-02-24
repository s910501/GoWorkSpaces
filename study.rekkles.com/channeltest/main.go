package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job    *job
	Result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func send(s chan<- *job) {
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		s <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func worker(s <-chan *job, resultChan chan<- *result) {
	defer wg.Done()
	for {
		job := <-s
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		newResult := &result{
			job:    job,
			Result: sum,
		}
		resultChan <- newResult
	}

}

func main() {
	wg.Add(1)
	go send(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go worker(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%d, sum%d\n", result.job.value, result.Result)
	}
	wg.Wait()

}
