package ants

import (
	"errors"
	"time"
)

var (
	// errQueueIsFull 在工作队列已满时返回
	// errQueueIsFull will be returned when the worker queue is full.
	errQueueIsFull = errors.New("the queue is full")

	// errQueueIsReleased 尝试将项目插入已释放的工作队列时将返回
	// errQueueIsReleased will be returned when trying to insert item to a released worker queue.
	errQueueIsReleased = errors.New("the queue length is zero")
)

type workerArray interface {
	len() int
	isEmpty() bool
	insert(worker *goWorker) error
	detach() *goWorker
	retrieveExpiry(duration time.Duration) []*goWorker
	reset()
}

type arrayType int

const (
	// 堆栈类型
	stackType arrayType = 1 << iota
	// 循环队列类型
	loopQueueType
)

// 创建一个工作队列
func newWorkerArray(aType arrayType, size int) workerArray {
	switch aType {
	case stackType:
		return newWorkerStack(size)
	case loopQueueType:
		return newWorkerLoopQueue(size)
	default:
		return newWorkerStack(size)
	}
}
