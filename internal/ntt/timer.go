package ntt

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	current_time float32
	End          float32
	Finished     bool
	Callback     func()
}

func NewTimer(end float32) Timer {
	return Timer{
		End:      end,
		Finished: true,
	}
}

func (t *Timer) Tick() {
	if !t.Finished {
		t.current_time += rl.GetFrameTime()
		if t.current_time >= t.End {
			t.Stop()
		}
	}
}

func (t *Timer) Start() {
	if t.Finished {
		t.Finished = false
	}
}

func (t *Timer) Stop() {
	if !t.Finished {
		t.current_time = 0
		t.Finished = true
		if t.Callback != nil {
			t.Callback()
		}
	}
}
