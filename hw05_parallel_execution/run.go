package hw05parallelexecution

import (
	"errors"
	"sync"
)

var (
	ErrErrorsLimitExceeded = errors.New("errors limit exceeded")
	ErrWrongParams         = errors.New("wrong params passed")
)

type (
	Task      func() error
	errBuffer struct {
		sync.RWMutex
		errLimit int
	}
)

func newErrBuffer(errLimit int) errBuffer {
	return errBuffer{errLimit: errLimit}
}

func (b *errBuffer) add() {
	b.Lock()
	b.errLimit--
	b.Unlock()
}

func (b *errBuffer) isFilled() bool {
	b.RLock()
	defer b.RUnlock()

	return b.errLimit > 1
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if len(tasks) < 1 || n < 1 || m < 1 {
		return ErrWrongParams
	}

	wg := sync.WaitGroup{}
	errBuff := newErrBuffer(m)

	chanToDo := make(chan Task, n)

	go func() {
		defer close(chanToDo)

		for _, t := range tasks {
			if errBuff.isFilled() {
				return
			}
			chanToDo <- t
		}
	}()

	for i := n; i > 0; i-- {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for t := range chanToDo {
				if errBuff.isFilled() {
					return
				}

				err := t()
				if err != nil {
					errBuff.add()
				}
			}
		}()
	}

	wg.Wait()

	if errBuff.isFilled() {
		return ErrErrorsLimitExceeded
	}

	return nil
}
