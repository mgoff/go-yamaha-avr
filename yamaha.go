package yamaha

import (
	"net/http"
	"strconv"
	"strings"
)

// struct used to hold our Yamaha connection
type Yamaha struct {
	address string
	client  *http.Client
}

// open a connection to the Yamaha device
func Connect(address string) (Yamaha, error) {
	return Yamaha{address: address, client: &http.Client{}}, nil
}

// common http post method
func (r Yamaha) Post(data string) error {
	req, err := http.NewRequest("POST", "http://"+r.address+"/YamahaRemoteControl/ctrl", strings.NewReader(data))
	if err != nil {
		return err
	}

	_, err = r.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// power on device
func (y Yamaha) PowerOn() error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Power_Control><Power>On</Power></Power_Control></Main_Zone></YAMAHA_AV>")
}

// power off device, which really just sets it to standby where network functions are still available
func (y Yamaha) PowerOff() error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Power_Control><Power>Standby</Power></Power_Control></Main_Zone></YAMAHA_AV>")
}

// set input to HDMI1
func (y Yamaha) SetInputHDMI1() error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Input><Input_Sel>HDMI1</Input_Sel></Input></Main_Zone></YAMAHA_AV>")
}

// set input to HDMI2
func (y Yamaha) SetInputHDMI2() error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Input><Input_Sel>HDMI2</Input_Sel></Input></Main_Zone></YAMAHA_AV>")
}

// set input to AirPlay
func (y Yamaha) SetInputAirPlay() error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Input><Input_Sel>AirPlay</Input_Sel></Input></Main_Zone></YAMAHA_AV>")
}

// Set volume from a range of 160 (full volume) to -800 (lowest volume before muting)
// which is displayed as +16.0dB to -80.0dB on the receiver. I never set this higher
// than -200 (-20.0dB) which is loud enough for my place. Be careful and don't blow out
// your speaker or ear drums!
func (y Yamaha) SetVolume(volume int) error {
	return y.Post("<YAMAHA_AV cmd=\"PUT\"><Main_Zone><Volume><Lvl><Val>" + strconv.Itoa(volume) + "</Val><Exp>1</Exp><Unit>dB</Unit></Lvl></Volume></Main_Zone></YAMAHA_AV>")
}
