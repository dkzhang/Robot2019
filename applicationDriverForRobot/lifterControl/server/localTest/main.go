package main

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

func UpOrDown(para int64) {
	//use go-rpio
	err := rpio.Open()
	if err != nil {
		return
	}
	defer rpio.Close()

	if para > 0 {
		//上升
		pin := rpio.Pin(5)
		pin.Output() // Output mode
		pin.High()   // Set pin High
		time.Sleep(time.Millisecond * time.Duration(para))
		pin.Low() // Set pin Low
	} else {
		//下降
		pin := rpio.Pin(6)
		pin.Output() // Output mode
		pin.High()   // Set pin High
		time.Sleep(time.Millisecond * time.Duration(-para))
		pin.Low() // Set pin Low
	}
}

func main() {
	UpOrDown(2)
	time.Sleep(time.Second * 5)
	UpOrDown(-2)
}
