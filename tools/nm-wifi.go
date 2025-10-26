package tools

import (
	"fmt"
	"strings"
	"vut/core"
)

type nmWifiFactory struct{}

func NMWifiFactory() core.Factory {
	return nmWifiFactory{}
}

// Get the unique name, and tags of this factory
func (nmWifiFactory) Identity() (string, []string) {
	return "nm-wifi", []string{"network", "wifi"}
}

// parser for nmcli commands

// STATE      CONNECTIVITY  WIFI-HW  WIFI     WWAN-HW  WWAN     METERED
// connected  full          enabled  enabled  missing  enabled  no (guessed)
// connected:full:enabled:enabled:missing:enabled:no (guessed)

type nmWifiGeneral struct {
	State        string
	Connectivity string
	WifiHw       string
	Wifi         string
	WWanHw       string
	WWan         string
	Metered      string
}

// Takes simplistic approach of split by ':'
func parseNmcliGeneral(in string) (nmWifiGeneral, error) {
	parts := strings.Split(in, ":")
	if len(parts) != 7 {
		return nmWifiGeneral{}, fmt.Errorf("Can't parse nmcli --terse general output: %s", in)
	}
	return nmWifiGeneral{
		State:        parts[0],
		Connectivity: parts[1],
		WifiHw:       parts[2],
		Wifi:         parts[3],
		WWanHw:       parts[4],
		WWan:         parts[5],
		Metered:      parts[6],
	}, nil
}

// check if required commands/hardware is available
func (nmWifiFactory) Check() []error {

	out, err := CommandOutput("nmcli --terse general",
		csvParser[nmWifiGeneral](':'))

	if err != nil {
		return []error{err}
	}
	if out == nil || len(*out) != 1 {
		return []error{fmt.Errorf("Can't parse output")}
	}
	first := (*out)[0]
	var errs []error
	// connected:full:enabled:enabled:missing:enabled:no (guessed)
	if first.State != "connected" {
		errs = append(errs, fmt.Errorf("State is not connected"))
	}
	if first.Wifi != "enabled" {
		errs = append(errs, fmt.Errorf("Wifi is not eabled"))
	}

	return errs

}

type nmWifiConnection struct {
	Name   string
	Uuid   string
	Type   string
	Device string
}

type nmWifiDevice struct {
	Device  string
	Choices []string
	Active  string
}

// generate the Tools
func (nmWifiFactory) Devices() ([]core.Device, error) {
	out, err := CommandOutput("nmcli --terse connection",
		csvParser[nmWifiConnection](':'))
	if err != nil {
		return nil, err
	}

	var choices []string
	// probably we need pairing of active and device.
	var active string
	var device string
	for _, conn := range *out {
		if conn.Type == "802-11-wireless" {
			choices = append(choices, conn.Name)
			if conn.Device != "" {
				active = conn.Name
				device = conn.Device
			}
		}
	}

	return []core.Device{nmWifiDevice{device, choices, active}}, nil
}

// Get the unique name, and tags of this tool
func (d nmWifiDevice) Identity() (string, []string) {
	return "nm-wifi", []string{"network", "wifi", d.Device, d.Active}
}

// to get identity of factory if needed
func (d nmWifiDevice) FactoryIdentity() (string, []string) {
	return nmWifiFactory{}.Identity()
}

// for cli mode, display the range/domain as string
func (d nmWifiDevice) Domain() string {
	return strings.Join(d.Choices, ", ")
}

// for cli mode, display the 'value' as string
func (d nmWifiDevice) Value() string {
	return d.Active
}

// for cli mode, update the 'value', if parsable by tool
func (d nmWifiDevice) Set(value string) (string, error) {
	var cmd []string
	if value == "" || value == "-" {
		if d.Active != "" {
			cmd = []string{"nmcli", "--terse", "connection", "down", d.Active}
		}
	} else {
		cmd = []string{"nmcli", "--terse", "connection", "up", value}
	}
	if cmd != nil {
		_, err := CommandOutputA(cmd, identity)
		return "", err
	}
	return "", nil
}
