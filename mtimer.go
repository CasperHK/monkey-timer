package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Hello World")
    t := NewTimer()
    t.Start()
    time.Sleep(3000 * time.Millisecond)
    t.Reset()
    // time.Sleep(3000 * time.Millisecond)
    // t.Start()
    // time.Sleep(3000 * time.Millisecond)
    // t.Stop()
    for {}
}

type Timer struct {
	seconds, minutes, hours int
	flagChan chan bool

}


func NewTimer() Timer {
	t := Timer{
			seconds	: 0, 
			minutes	: 0, 
			hours	: 0, 
			flagChan: make(chan bool)}
	return t
}

func clockTick(t *Timer) {
	for {
			t.PrintTime()
			fmt.Println("")

		time.Sleep(1000 * time.Millisecond)
		
		t.seconds++
		if t.seconds == 60 {
			t.minutes++
			t.seconds = 0
		}
		if t.minutes == 60 {
			t.hours++
		}



		var isStop bool
		select {
			case isStop = <- t.flagChan:
			default:
		}
		if isStop {
			break
		}
	}
}

func (t *Timer) Start() {
	go clockTick(t)
}

func (t *Timer) Stop() {
	t.flagChan <- true
}

func (t *Timer) Reset() {
	t.seconds = 0 
	t.minutes = 0 
	t.hours   = 0 
}

func (t *Timer) GetTime() string {
	return fmt.Sprintf("%.2d:%.2d:%.2d", t.hours, t.minutes, t.seconds)
}

func (t *Timer) PrintTime() {
	fmt.Printf(t.GetTime())
}

func (t *Timer) SetCallback(callback func()) {

}