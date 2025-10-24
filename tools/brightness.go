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

func (factoryBrightness) Check() bool {
	out, ok := utils.CommandOutput("brightnessctl", "--version")
	if !ok {
		return false
	}

	ver, err := strconv.ParseFloat(strings.TrimSpace(out), 64)
	if err != nil {
		utils.LogError("brightnessctl" + ": Error finding version :" + out)
		return false
	}
	if ver < BRIGHTNESSCTL_VER {
		utils.LogError("brightnessctl" + ": Expected version >= 0.5, got" + out)
		return false
	}
	return true
}

func (factoryBrightness) Devices() []core.Device {
	out, ok := utils.CommandOutput("brightnessctl", "--machine-readable", "--list")
	if !ok {
		return nil
	}
	devices, err := parser.ParseCSVIntoStructs[brightnessDevice](out)
	if err != nil {
		utils.LogError("Error parsing CSV")
		return nil
	}
	var ret []core.Device
	for _, device := range devices {
		ret = append(ret, device)

	}
	return ret
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
	//_, err := ParseValue(t.MaxValue, value)
	//if err != nil {
	//	LogError("Can't parse input: " + value)
	//	return "", err
	//}

	out, ok := utils.CommandOutput("brightnessctl", "--machine-readable", "--device="+t.Name,
		"set", value)
	if !ok {
		return "", fmt.Errorf("can't set")
	}

	devices, err := parser.ParseCSVIntoStructs[brightnessDevice](out)
	if err != nil || len(devices) != 1 {
		utils.LogError("Error parsing CSV")
		return "", err
	}
	return string(devices[0].Current), nil
}
