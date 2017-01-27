package main

import (
    "fmt"
)

type MomentAction struct {
	id int
	seconds, minutes, hours int
	callback func()
}

func NewMomentAction() MomentTask {
	ta := MomentAction{
				id 		: 0,
				seconds	: 0, 
				minutes	: 0, 
				hours	: 0,
				callback: nil}
	return ta
}
