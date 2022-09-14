package main

import (
	"log"
	"sync"
)

type Handler func(data interface{})

type Line struct {
	factory   *Factory
	handler   Handler
	MaxWorker int
	task      chan interface{}
	number    int64
	close     chan struct{}
	wg        sync.WaitGroup
	once      sync.Once
	stopOnce  sync.Once
}

func (l *Line) Wait() {
	l.wg.Wait()
}

func (l *Line) Submit(data interface{}) {
	l.once.Do(func() {
		for i := 0; i < l.MaxWorker; i++ {
			l.factory.wg.Add(1)
			go l.worker()
		}
	})
	l.wg.Add(1)
	select {
	case l.task <- data:
	}
}

func (l *Line) do(task interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
		l.wg.Done()
	}()
	l.handler(task)
}

func (l *Line) worker() {
	var (
		task interface{}
	)
	for {
		select {
		case <-l.close:
			goto CLOSE
		case task = <-l.task:
			l.do(task)
		}
	}
CLOSE:
	l.factory.wg.Done()
}

func (l *Line) stop() {
	l.stopOnce.Do(func() {
		close(l.close)
	})
}
