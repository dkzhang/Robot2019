package client

import (
	"testing"
	"time"
)

func TestSwitchLaserLight(t *testing.T) {
	SwitchLaserLight(true)

	time.Sleep(time.Second * 10)

	SwitchLaserLight(false)
}
