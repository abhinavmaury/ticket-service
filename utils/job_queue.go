package utils

import (
	"context"
	"fmt"
	"sync"
)

type JobFunc func(ctx context.Context, data interface{}) error

type Job struct {
	Name string
	Fun  JobFunc
	Data interface{}
}

type JobQueue struct {
	queue     []Job
	qInput    chan Job
	work      chan Job
	batchChan chan Job
}

var jobQueue *JobQueue
var once sync.Once

var pool sync.Map

func (b *JobQueue) processTasks(ctx context.Context, n Job) {
	b.batchChan <- n
	fmt.Println("JOB_RUNNING", n.Name)
	err := n.Fun(ctx, n.Data)
	if err != nil {
		fmt.Println("JOB_FAILED", n.Name+"\n"+err.Error())
	}
	fmt.Println("JOB_COMPLETED", n.Name)
	<-b.batchChan
}

func (b *JobQueue) worker() {
	for {
		select {
		case n := <-b.work:
			_ctx := context.Background()
			ctx := context.WithValue(_ctx, 1,
				map[string]interface{}{
					"requestID": "job_queue_" + n.Name,
					"sessionID": "session_" + n.Name,
				})
			fmt.Println("JOB_WAITING", n.Name)
			go b.processTasks(ctx, n)
		}
	}
}

func (b *JobQueue) scheduler() {
	for {
		if len(b.queue) == 0 {
			select {
			case i := <-b.qInput:
				b.queue = append(b.queue, i)
			}
		} else {
			select {
			case i := <-b.qInput:
				b.queue = append(b.queue, i)
			case b.work <- b.queue[0]:
				b.queue = b.queue[1:]
			}
		}
	}
}

func (b *JobQueue) Insert(id Job) {
	b.qInput <- id
}

func GetJobQueue(agents int) *JobQueue {
	if agents == 0 {
		agents = 4
	}
	once.Do(func() {
		jobQueue = &JobQueue{
			qInput: make(chan Job), work: make(chan Job),
			batchChan: make(chan Job, agents)}
		go jobQueue.scheduler()
		go jobQueue.worker()
	})
	return jobQueue
}