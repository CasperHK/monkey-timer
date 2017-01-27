package main

import (
    "fmt"
)

type MomentTask struct {
	id int
	seconds, minutes, hours int
	callback func()
}

func NewMomentTask() MomentTask {
	ta := MomentTask{
				id 		: 0,
				seconds	: 0, 
				minutes	: 0, 
				hours	: 0,
				callback: nil}
	return ta
}

func (t *Timer) SetMomentTask() {

}
