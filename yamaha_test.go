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
func TestYamahaAvr (t *testing.T) {

	// open the connection to the Yamaha device
	device, err := Connect(address)
	if err != nil {
		t.Error(err)
		return
	}

	// power it on
	err = device.PowerOn()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)

	// set input to HDMI1
	err = device.SetInputHDMI1()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)

	// set volume to -30dB
	err = device.SetVolume(-300)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)

	// set input to HDMI2
	err = device.SetInputHDMI2()
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)

	// set volume to -40dB
	err = device.SetVolume(-400)
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(5 * time.Second)

	// power it on
	err = device.PowerOff()
	if err != nil {
		t.Error(err)
		return
	}
}