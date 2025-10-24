package core

// Generate one or more Tools
// e.g. Multiple Audio Cards or Monitors
type Factory interface {

	// Get the unique name, and tags of this factory
	Identity() (string, []string)

	// check if required commands/hardware is available
	Check() bool

	// generate the Tools
	Devices() []Device
}

// Tool is a single .. well, tool
type Device interface {

	// Get the unique name, and tags of this tool
	Identity() (string, []string)

	// to get identity of factory if needed
	FactoryIdentity() (string, []string)

	// for cli mode, display the range/domain as string
	Domain() string

	// for cli mode, display the 'value' as string
	Value() string

	// for cli mode, update the 'value', if parsable by tool
	Set(value string) (string, error)
}
