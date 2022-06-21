package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type errorCounter struct {
	mu      sync.Mutex
	counter int
	limit   int
}

func (ec *errorCounter) isEnough() bool {
	defer ec.mu.Unlock()
	ec.mu.Lock()
	return ec.counter > ec.limit
}

func (ec *errorCounter) increase() {
	defer ec.mu.Unlock()
	ec.mu.Lock()
	ec.counter++
}

type Worker struct {
	wg           *sync.WaitGroup
	taskChan     chan Task
	errorCounter *errorCounter
}

func (w Worker) Do() {
	for {
		if w.errorCounter.isEnough() {
			break
		}
		task, ok := <-w.taskChan
		if !ok {
			break
		}
		taskError := task()
		if taskError != nil {
			w.errorCounter.increase()
		}
	}
	w.wg.Done()
}

func putToChannel(tasks []Task, channel chan<- Task) {
	for _, task := range tasks {
		channel <- task
	}
	close(channel)
}

func Run(tasks []Task, n, m int) error {
	if m == 0 || n == 0 {
		return nil
	}
	errorCounter := &errorCounter{
		counter: 0,
		limit:   m,
	}
	taskChan := make(chan Task, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		worker := Worker{
			wg:           &wg,
			taskChan:     taskChan,
			errorCounter: errorCounter,
		}
		go worker.Do()
	}

	go putToChannel(tasks, taskChan)
	wg.Wait()
	if errorCounter.isEnough() {
		return ErrErrorsLimitExceeded
	}
	return nil
}
