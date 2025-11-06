package tools

import (
	"strings"
	"vut/core"
)

type factoryDarkmode struct{}
type deviceDarkmode struct {
	isDark bool
}

func DarkmodeFactory() factoryDarkmode {
	return factoryDarkmode{}
}

func (factoryDarkmode) Identity() (string, []string) {
	return "darkmode", []string{"mode", "display", "dark", "light"}
}

func (factoryDarkmode) Check() []error {
	return nil
}

func (factoryDarkmode) Devices() ([]core.Device, error) {
	var ret []core.Device
	code, err := CommandOutput(
		"gsettings get org.gnome.desktop.interface color-scheme",
		identity)
	isDark := false
	if err == nil && code != nil && strings.Contains(*code, "prefer-dark") {
		isDark = true
	} else {
		isDark = false
	}

	ret = append(ret, &deviceDarkmode{isDark})
	return ret, nil

}

// Get the unique name, and tags of this tool
func (deviceDarkmode) Identity() (string, []string) {
	return "darkmode", []string{"mode", "display", "dark", "light"}
}

// to get identity of factory if needed
func (deviceDarkmode) FactoryIdentity() (string, []string) {
	return "darkmode", []string{"mode", "display", "dark", "light"}
}

// for cli mode, display the range/domain as string
func (deviceDarkmode) Domain() string {
	return "on/off"
}

// for cli mode, display the 'value' as string
func (d deviceDarkmode) Value() string {
	if d.isDark {
		return "dark"
	} else {
		return "light"
	}
}

// for cli mode, update the 'value', if parsable by tool
func (d *deviceDarkmode) Set(value string) (string, error) {
	val := "default"
	ret := "light"
	mxMode := "leuven"
	footSig := "-USR1"
	d.isDark = false
	if value == "on" || value == "dark" || value == "true" || value == "1" {
		val = "prefer-dark"
		d.isDark = true
		ret = "dark"
		mxMode = "leuven-dark"
		footSig = "-USR2"
	}
	_, err := CommandOutputA(
		[]string{"gsettings", "set", "org.gnome.desktop.interface", "color-scheme", val},
		identity)

	// Handle Emacs
	mxServers, _ := listEmacsServers()
	for _, mx := range mxServers {
		_, _ = CommandOutputA(
			[]string{"emacsclient", "-s", mx, "-n", "-e",
				"(load-theme '" + mxMode + ")"},
			identity)
	}
	foots, _ := CommandOutput("pgrep foot", words)
	for _, ps := range *foots {
		_, _ = CommandOutputA(
			[]string{"kill", footSig, ps}, identity)
	}

	// Handle foot, assumes [colors] is light, and [colors2] is dark

	return ret, err
}

func words(in string) ([]string, error) {
	return strings.Fields(in), nil
}
