package client

import (
	"testing"
	"time"
)

func TestLifterControl(t *testing.T) {
	LifterControl(1000 * 10)
	time.Sleep(time.Second * 3)
	LifterControl(-1000 * 11)
}
