package tools

import (
	"fmt"
	"strconv"
	"strings"
	"vut/core"
	"vut/parser"
	"vut/utils"
)

const BRIGHTNESSCTL_VER = 0.5

type factoryBrightness struct{}

func BrightnessFactory() factoryBrightness {
	return factoryBrightness{}
}

type brightnessDevice struct {
	Name     string `csv:"0"`
	Type     string `csv:"1"`
	Current  int    `csv:"2"`
	Percent  string `csv:"3"`
	MaxValue int    `csv:"4"`
}

func (factoryBrightness) Identity() (string, []string) {
	return "brightness", []string{"brightness", "display", "monitor"}
}

func (factoryBrightness) Check() []error {
	var errs []error
	ver, err := utils.CommandOutput("brightnessctl --version",
		func(s string) (float64, error) {
			return strconv.ParseFloat(strings.TrimSpace(s), 64)
		})
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	if *ver < BRIGHTNESSCTL_VER {
		errs = append(errs,
			fmt.Errorf("brightnessctl : Expected version >= 0.5, got %d", ver))
		return errs
	}
	return errs
}

func (factoryBrightness) Devices() ([]core.Device, error) {
	devices, err := utils.CommandOutput[[]brightnessDevice](
		"brightnessctl --machine-readable --list",
		parser.ParseCSVIntoStructs)

	if err != nil {
		return nil, err
	}
	var ret []core.Device
	for _, device := range *devices {
		ret = append(ret, device)

	}
	return ret, nil
}

// Get the unique name, and tags of this tool
func (t brightnessDevice) Identity() (string, []string) {
	return t.Name, []string{"brightness", "display", "monitor", t.Type}
}

// to get identity of factory if needed
func (t brightnessDevice) FactoryIdentity() (string, []string) {
	return factoryBrightness{}.Identity()
}

// for cli mode, display the range/domain as string
func (t brightnessDevice) Domain() string {
	return "0 .. " + strconv.Itoa(t.MaxValue)
}

// for cli mode, display the 'value' as string
func (t brightnessDevice) Value() string {
	return strconv.Itoa(t.Current)
}

// for cli mode, update the 'value', if parsable by tool
func (t brightnessDevice) Set(value string) (string, error) {

	devices, err := utils.CommandOutputA[[]brightnessDevice](
		[]string{"brightnessctl",
			"--machine-readable",
			"--device",
			t.Name,
			"set",
			value},
		parser.ParseCSVIntoStructs)
	if err != nil || len(*devices) != 1 {
		return "", err
	}
	return strconv.Itoa((*devices)[0].Current), nil
}
