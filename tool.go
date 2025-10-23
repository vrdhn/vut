package main

// Generate one or more Tools
// e.g. Multiple Audio Cards or Monitors
type Factory interface {

	// Get the unique name, and tags of this factory
	identity() (string, []string)

	// check if required commands/hardware is available
	check() bool

	// generate the Tools
	tools() []Tool
}

// Tool is a single .. well, tool
type Tool interface {

	// Get the unique name, and tags of this tool
	identity() (string, []string)

	// to get identity of factory if needed
	factoryIdentity() (string, []string)

	// for cli mode, display the range/domain as string
	domain() string

	// for cli mode, display the 'value' as string
	value() string

	// for cli mode, update the 'value', if parsable by tool
	set(value string) (string, error)
}
