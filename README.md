# go-yamaha-avr
A Go library for controlling Yahama Audio/Video receivers with network connections through the Yamaha Control API.

Yamaha Control API
==================
Yamaha has a simple API for controlling their receivers over the local network. This library implments a subset of the available network commands and has been tested with a Yamaha RX-V481 receiver.

Usage
=====

Call `yamaha.Connect()` with the IP address of the Yamaha receiver you'd like control, such as `192.168.1.50`.  Then call various functions to control the receiver as described in [yamaha.go](yamaha.go).

````go
	package main

	import (
		"github.com/mgoff/go-yamaha-avr"
		"log"
		"time"
	)

	func main() {

		// open the connection to the Yamaha device
		device, err := yamaha.Connect("192.168.1.50")
		if err != nil {
			log.Fatal(err)
		}

		// power it on
		err = device.PowerOn()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

		// set input to HDMI1
		err = device.SetInputHDMI1()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

		// set volume to -30dB
		err = device.SetVolume(-300)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

		// set input to HDMI2
		err = device.SetInputHDMI2()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

		// set volume to -40dB
		err = device.SetVolume(-400)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)

		// power it on
		err = device.PowerOff()
		if err != nil {
			log.Fatal(err)
		}
	}
````