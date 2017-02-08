package yamaha

import (
	"testing"
	"time"
)

// udpate address with you Yamaha receiver's IP address before running 'go test'
const (
	address = "192.168.1.50"
)

// sleep in between test commands so you can watch the results on your receiver's display
func TestYamahaAvr(t *testing.T) {

	// open the connection to the Yamaha device
	device, err := Connect(address)
	if err != nil {
		t.Error(err)
	}

	// power it on
	if device.PowerOn() != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)

	// set input to HDMI1
	if device.SetInputHDMI1() != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)

	// set volume to -30dB
	if device.SetVolume(-300) != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)

	// set input to HDMI2
	if device.SetInputHDMI2() != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)

	// set volume to -40dB
	if device.SetVolume(-400) != nil {
		t.Error(err)
	}
	time.Sleep(5 * time.Second)

	// power it on
	if device.PowerOff() != nil {
		t.Error(err)
	}
}
