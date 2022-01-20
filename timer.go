package timer

import (
	"sync"
	"time"
)

type Timer struct {
	tm      *time.Timer
	mtx     *sync.Mutex
	ev_ch   chan struct{}
	running bool
}

func NewTimer() *Timer {
	return &Timer{
		mtx:     &sync.Mutex{},
		running: false,
		ev_ch:   make(chan struct{}, 1),
	}
}

func (t *Timer) Recv() <-chan struct{} {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	return t.ev_ch
}

func (t *Timer) on_time() {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if !t.running {
		return
	}

	select {
	case t.ev_ch <- struct{}{}:
	default:
	}
}

func (t *Timer) Stop() {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if !t.running {
		return
	}
	t.stop()
}

func (t *Timer) Start(d time.Duration) {
	t.mtx.Lock()
	defer t.mtx.Unlock()

	if t.running {
		t.stop()
	}
	t.start(d)
}

func (t *Timer) stop() {
	t.running = false

	t.tm.Stop()
	select {
	case <-t.ev_ch:
	default:
	}
}

func (t *Timer) start(d time.Duration) {
	if t.tm == nil {
		t.tm = time.AfterFunc(d, t.on_time)
	} else {
		t.tm.Reset(d)
	}

	t.running = true
}
