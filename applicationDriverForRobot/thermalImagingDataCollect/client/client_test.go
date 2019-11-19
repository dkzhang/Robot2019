package client

import "testing"

func TestCollectThermalImagingData(t *testing.T) {
	/////////////////////////////////
	// Set up a connection to the server.
	//address := "localhost:50061"
	address := "192.168.1.109:50061"

	dataArray, err := CollectThermalImagingData(address)

	if err != nil {
		t.Errorf("CollectThermalImagingData error: %v", err)
	} else {
		t.Logf("CollectThermalImagingData reply: %v", dataArray)
	}
}
