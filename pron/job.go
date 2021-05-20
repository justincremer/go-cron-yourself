package pron

import (
	"os/exec"
	"time"
)

var (
	Tab Prontab
)

// Top level pron struct
type Prontab struct {
	ticker *time.Ticker
	jobs   []job
}

// Interface for external and internal jobs
type job interface {
	AddJob()
	RemoveJob()
	Dispatch()
}

// External job rollup w/ time maps
type externalJob struct {
	t      schedule
	action externalAction
}

// Internal job rollup w/ time maps
type internalJob struct {
	t      schedule
	action internalAction
}

// Initializes the tab
func Initialize(t time.Duration) {
	Tab.ticker = time.NewTicker(t)
}

// Emptys the job buffer and stops the clock
func Shutdown() {
	Tab.jobs = nil
	Tab.ticker.Stop()
}

func (p *Prontab) AddJob(schedule string, action action) error {
	// TODO: implement

}

func (p *Prontab) MustAddJob(schedule string, fn *exec.Cmd) {
	if err := p.AddJob(schedule, fn); err != nil {
		panic(err)
	}
}
