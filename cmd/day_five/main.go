package main

import (
	"fmt"
	"sync"
	"time"
)

/*
🕒 Prompt: Factory Worker Timesheet Management System

Please implement a system to track and compute worker hours with the following features:

🧱 Required Methods:
clock_in(worker_id string) error
clock_out(worker_id string) error
get_hours(worker_id string, date time.Time) (float64, error)
calculate_overtime(worker_id string, date time.Time) (float64, error)

📌 Considerations:
	•	Workers may clock in and out multiple times per day.
	•	Workers may be in different time zones (assume UTC for internal logic unless extended).
	•	Your system must guard against bad input (e.g., clocking out before clocking in).
	•	Overtime starts after 8 hours/day unless otherwise specified.
*/

type Shift struct {
	clock_in time.Time
	clock_out time.Time
}

type TimeSheet struct {
	l  sync.Mutex
	ts map[string]map[string][]Shift
}

func (t *TimeSheet) clock_in(worker_id string) error {
	t.l.Lock()
	defer t.l.Unlock()

	today := time.Now().Format("2006-01-02")
	sh := Shift {
		clock_in: time.Now(),
	}

	// first we need to guard against invalid
	// get the date, check if an entry exists. If it does, loop the list of clock_in values
	val, ok := t.ts[worker_id]
	if ok {
		vtime, exists := val[today]
		if exists { // if there is a slice of shifts here, we need to loop to ensure that our clock_in time is not before a clock_out time
			for _, v := range vtime {
				if v.clock_out.IsZero()  {
					return fmt.Errorf("can't clock-in without having clocked out")
				}
				if time.Now().Before(v.clock_out) {
					return fmt.Errorf("invalid clock-in time")
				}
			}
			vtime = append(vtime, sh)
			val[today] = vtime
			t.ts[worker_id] = val
			return nil
		}
	}


	// if no entry exists, we add it
	m := make(map[string][]Shift)
	m[today] = []Shift{sh}
	t.ts[worker_id] = m

	return nil
}


func (ts *TimeSheet) clock_out(worker_id string) error {return nil}
func (ts *TimeSheet) get_hours(worker_id string, date time.Time) (float64, error) { return nil}
func (ts *TimeSheet) calculate_overtime(worker_id string, date time.Time) (float64, error) {return nil}

func main() {
	fmt.Println("here")
}
