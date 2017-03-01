package periodic

import (
	"time"

	l "github.com/vessel/pkg/logger"
)

var logger = l.NewLogger()
var Info = logger.Info
var Warning = logger.Warning
var Error = logger.Error

type PeriodicJob struct {
	Time   time.Duration
	Unit   string
	Finish chan bool
}

func NewPeriodicJob(t time.Duration, unit string, f chan bool) *PeriodicJob {
	return &PeriodicJob{Time: t, Unit: unit, Finish: f}
}

func SleepTime(t time.Duration, unit string) {
	switch unit {
	case "sec":
		time.Sleep(t * time.Second)
	case "min":
		time.Sleep(t * time.Minute)
	}
}

// Create a PeriodicJob that runs every T units of time
func (p PeriodicJob) LanuchPeriodicJob(masterswitch chan bool) {
	go func() {
		for {
			select {
			case <-masterswitch:
				Info.Print("Main Loop has exited")
				return
			case <-p.Finish:
				Info.Print("Exiting Periodic Job")
				return
			default:
				SleepTime(p.Time, p.Unit)
				Info.Print("Performing Database Vacuum")
			}
		}
	}()
}
