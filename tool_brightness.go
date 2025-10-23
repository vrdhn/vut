package main

import (
	"fmt"
	"strconv"
	"strings"
)

const BRIGHTNESSCTL_VER = 0.5

type FactoryBrightness struct{}

type Device struct {
	Name     string `csv:"0"`
	Type     string `csv:"1"`
	Value    int    `csv:"2"`
	Percent  string `csv:"3"`
	MaxValue int    `csv:"4"`
}

func (FactoryBrightness) identity() (string, []string) {
	return "brightness", []string{"brightness", "display", "monitor"}
}

func (FactoryBrightness) check() bool {
	out, ok := CommandOutput("brightnessctl", "--version")
	if !ok {
		return false
	}

	ver, err := strconv.ParseFloat(strings.TrimSpace(out), 64)
	if err != nil {
		LogError("brightnessctl" + ": Error finding version :" + out)
		return false
	}
	if ver < BRIGHTNESSCTL_VER {
		LogError("brightnessctl" + ": Expected version >= 0.5, got" + out)
		return false
	}
	return true
}

func (FactoryBrightness) tools() []Tool {
	out, ok := CommandOutput("brightnessctl", "--machine-readable", "--list")
	if !ok {
		return nil
	}
	devices, err := ParseCSVIntoStructs[Device](out)
	if err != nil {
		LogError("Error parsing CSV")
		return nil
	}
	var ret []Tool
	for _, device := range devices {
		ret = append(ret, device)

	}
	return ret
}

// Get the unique name, and tags of this tool
func (t Device) identity() (string, []string) {
	return t.Name, []string{"brightness", "display", "monitor", t.Type}
}

// to get identity of factory if needed
func (t Device) factoryIdentity() (string, []string) {
	return FactoryBrightness{}.identity()
}

// for cli mode, display the range/domain as string
func (t Device) domain() string {
	return "0 .. " + strconv.Itoa(t.MaxValue)
}

// for cli mode, display the 'value' as string
func (t Device) value() string {
	return strconv.Itoa(t.Value)
}

// for cli mode, update the 'value', if parsable by tool
func (t Device) set(value string) (string, error) {
	//_, err := ParseValue(t.MaxValue, value)
	//if err != nil {
	//	LogError("Can't parse input: " + value)
	//	return "", err
	//}

	out, ok := CommandOutput("brightnessctl", "--machine-readable", "--device="+t.Name,
		"set", value)
	if !ok {
		return "", fmt.Errorf("can't set")
	}

	devices, err := ParseCSVIntoStructs[Device](out)
	if err != nil || len(devices) != 1 {
		LogError("Error parsing CSV")
		return "", err
	}
	return string(devices[0].Value), nil
}
