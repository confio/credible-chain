package queue

import (
	"fmt"
	"sync/atomic"

	app "github.com/confio/credible-chain/app"
	wc "github.com/confio/credible-chain/weaveclient"
	"github.com/confio/credible-chain/x/votes"
)

type Queue struct {
	input chan *Task

	entered  int64
	errored  int64
	finished int64
}

type Stats struct {
	Pending  int64
	Errored  int64
	Finished int64
}

type Task struct {
	Vote     *votes.VoteRecord
	Tx       *app.Tx
	Response chan wc.BroadcastTxResponse
	Error    error
}

func (t *Task) WithError(err error) *Task {
	t.Error = err
	return t
}

func (t *Task) Validate() error {
	if t.Error != nil {
		return t.Error
	}
	return t.Vote.Validate()
}

type Worker interface {
	Run(tasks <-chan *Task) <-chan *Task
}

func NewQueue(queueSize int) *Queue {
	q := &Queue{
		input: make(chan *Task, queueSize),
	}
	return q
}

// Push starts a task
func (q *Queue) Push(task *Task) error {
	if err := task.Validate(); err != nil {
		return err
	}
	fmt.Println("Add entered")
	atomic.AddInt64(&q.entered, 1)
	q.input <- task
	return nil
}

// Run will start a number of go-routines to process everything
// It will block until all work is done and channels are closed
func (q *Queue) Run(worker Worker) {
	out := worker.Run(q.input)
	// read the output of the worker forever, updating stats
	for {
		task, more := <-out
		if !more {
			return
		}
		fmt.Println("Task completed")
		q.done(task)
	}
	fmt.Println("Run finished")
}

// Pending is how many have been added, but not finished
func (q *Queue) Pending() int64 {
	return q.entered - q.errored - q.finished
}

func (q *Queue) Stats() Stats {
	return Stats{
		Pending:  q.Pending(),
		Errored:  q.errored,
		Finished: q.finished,
	}
}

func (q *Queue) done(task *Task) {
	if task.Error != nil {
		fmt.Println("Add error")
		// TODO: mark which one it was for retry
		atomic.AddInt64(&q.errored, 1)
	} else {
		fmt.Println("Add finished")
		atomic.AddInt64(&q.finished, 1)
	}
}
