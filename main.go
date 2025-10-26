package main

import (
	"fmt"
	"os"
	"strings"

	"vut/core"
	"vut/tools"
)

func getFactories() []core.Factory {

	return []core.Factory{
		tools.BrightnessFactory(),
		tools.NMWifiFactory(),
	}
}

func getTools() ([]core.Device, []error) {

	var errs []error
	var tools []core.Device
	for _, f := range getFactories() {
		devices, err := f.Devices()
		if err == nil {
			tools = append(tools, devices...)
		} else {
			errs = append(errs, err)
		}
	}
	return tools, errs
}

func getTool(needle string) ([]core.Device, []error) {

	var tools []core.Device
	var errs []error

	for _, f := range getFactories() {
		devices, err := f.Devices()
		if err == nil {
			for _, t := range devices {
				fn, _ := t.FactoryIdentity()
				tn, _ := t.Identity()
				if strings.Contains(fn, needle) || strings.Contains(tn, needle) {
					tools = append(tools, t)
				}
			}
		} else {
			errs = append(errs, err)
		}
	}
	return tools, errs
}

func usage() {

}

func show(devices []core.Device, errs []error) {
	if len(devices) > 0 {
		fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
			"Tool   ", "     Device", "Value", "Range")
		for _, t := range devices {
			fn, _ := t.FactoryIdentity()
			tn, _ := t.Identity()
			fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
				fn, tn, t.Value(), t.Domain())
		}
	}
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Printf("** %s\n", e.Error())
		}
	}
}

func main() {
	args := os.Args[1:]
	switch len(args) {
	case 0:
		show(getTools())
	case 1:
		show(getTool(args[0]))
	case 2:
		tools, errs := getTool(args[0])
		if len(errs) > 0 {
			for _, e := range errs {
				fmt.Printf("** %s\n", e.Error())
			}
		}

		for _, t := range tools {
			t.Set(args[1])
		}
		show(getTool(args[0]))
	default:
		usage()
	}

}
