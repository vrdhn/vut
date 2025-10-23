package main

import (
	"strings"
)

func GetFactories() []Factory {

	return []Factory{
		FactoryBrightness{},
	}
}

func GetTools() []Tool {

	var tools []Tool
	for _, f := range GetFactories() {
		tools = append(tools, f.tools()...)
	}
	return tools
}

func GetTool(needle string) []Tool {

	var tools []Tool
	for _, f := range GetFactories() {
		for _, t := range f.tools() {
			fn, _ := t.factoryIdentity()
			tn, _ := t.identity()
			if strings.Contains(fn, needle) || strings.Contains(tn, needle) {
				tools = append(tools, t)
			}
		}
	}
	return tools
}
