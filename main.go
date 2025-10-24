package main

import (
	"fmt"
	"os"
	"strings"

	"vut/core"
	"vut/tools"
	"vut/utils"
)

func getFactories() []core.Factory {

	return []core.Factory{
		tools.BrightnessFactory(),
	}
}

func getTools() []core.Device {

	var tools []core.Device
	for _, f := range getFactories() {
		tools = append(tools, f.Devices()...)
	}
	return tools
}

func getTool(needle string) []core.Device {

	var tools []core.Device
	for _, f := range getFactories() {
		for _, t := range f.Devices() {
			fn, _ := t.FactoryIdentity()
			tn, _ := t.Identity()
			if strings.Contains(fn, needle) || strings.Contains(tn, needle) {
				tools = append(tools, t)
			}
		}
	}
	return tools
}

func usage() {

}

func show(devices []core.Device) {
	fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
		"Device  ", "       Name", "Value", "Range")
	for _, t := range devices {
		fn, _ := t.FactoryIdentity()
		tn, _ := t.Identity()
		fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
			fn, tn, t.Value(), t.Domain())
	}

}

func main() {
	utils.StartLogger()
	args := os.Args[1:]
	switch len(args) {
	case 0:
		show(getTools())
	case 1:
		show(getTool(args[0]))
	case 2:
		for _, t := range getTool(args[0]) {
			t.Set(args[1])
		}
		show(getTool(args[0]))
	default:
		usage()
	}

}
