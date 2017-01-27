package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello World")
    t := NewTimer()
    t.Start()
    t.DisplayTime(true)
    time.Sleep(3000 * time.Millisecond)
    t.Reset()

    // time.Sleep(3000 * time.Millisecond)
    // t.Start()
    // time.Sleep(3000 * time.Millisecond)
    // t.Stop()
    for {}
}

// timer data set
type Timer struct {
	seconds, minutes, hours int
	flagChan chan bool
	displayTime bool  // whether display time on console
}


// create a new timer
func NewTimer() Timer {
	t := Timer{
			seconds	: 0, 
			minutes	: 0, 
			hours	: 0, 
			flagChan: make(chan bool),
			displayTime: false}
	return t
}

// a clocktick
func clockTick(t *Timer) {
	for {
		// display time according to setting
		if t.displayTime {
			fmt.Println(t.GetTime())
		}

		// time of 1 second
		time.Sleep(1000 * time.Millisecond)
		
		// calculate new time
		t.seconds++
		if t.seconds == 60 {
			t.minutes++
			t.seconds = 0
		}
		if t.minutes == 60 {
			t.hours++
		}


		// declaring flag
		var isStop bool

		// receive flag status from chan
		select {
			case isStop = <- t.flagChan:
			default:
		}

		// stop timer according to flag status
		if isStop {
			break
		}
	}
}

// start the timer
func (t *Timer) Start() {
	go clockTick(t)
}

// stop the timer
func (t *Timer) Stop() {
	t.flagChan <- true
}

// reset the timer to 00:00:00
func (t *Timer) Reset() {
	t.seconds = 0 
	t.minutes = 0 
	t.hours   = 0 
}

// configure the timer to display the timer on console or not
func (t *Timer) DisplayTime(isDisplay bool) {
	t.displayTime = isDisplay
}

// get the current time string
func (t *Timer) GetTime() string {
	return fmt.Sprintf("%.2d:%.2d:%.2d", t.hours, t.minutes, t.seconds)
}

// print time on console
func (t *Timer) PrintTime() {
	fmt.Printf(t.GetTime())
}

// set a callback function as moment action that will be executed in a defined time
func (t *Timer) SetMomentAction(hour, minute, second int, fn func()) {

}

// set a callback function as a period action that will be executed periodically after a period
func (t *Timer) SetPeriodAction(period int, fn func()) {

}