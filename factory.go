package pipeline

import "sync"

type Factory struct {
	PanicHandler func(err any)
	lines        []*Line
	wg           sync.WaitGroup
}

func (f *Factory) Release() {
	for _, line := range f.lines {
		line.stop()
	}
	f.wg.Wait()
}

func (f *Factory) NewLine(max int, handler Handler) *Line {
	line := &Line{
		factory:   f,
		task:      make(chan interface{}, max*2),
		close:     make(chan struct{}),
		handler:   handler,
		MaxWorker: max,
	}
	f.lines = append(f.lines, line)
	return line
}

func (f *Factory) Wait() {
	for _, line := range f.lines {
		line.Wait()
	}
}
