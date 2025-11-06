package tools

//import (
//	"fmt"
//	"vut/core"
//)
//
///// pactl -f json  list sources
///// pactl -f json  list sink
//
//type pulseaudioFactory struct {
//}
//
//type pulseaudioDevice struct {
//}
//
//// Get the unique name, and tags of this factory
//func (pulseaudioFactory) Identity() (string, []string) {
//	return "pulseaudio", []string{"sound", "audio"}
//}
//
//// check if required commands/hardware is available
//func (pulseaudioFactory) Check() []error {
//	json, err := CommandOutput("pactl  -f json info",
//		parseJson)
//	if err != nil {
//		return []error{err}
//	}
//	ver, err := json.GetString("server_version")
//	if err != nil {
//		return []error{err}
//	}
//	if ver != "15.0.0" {
//		if err != nil {
//			return []error{fmt.Errorf("Server version should be 15.0.0, got %s", ver)}
//		}
//	}
//	return nil
//}
//
//// generate the Tools
//func (pulseaudioFactory) Devices() ([]core.Device, error) {
//	json, err := CommandOutput("pactl  -f json list",
//		parseJson)
//	if err != nil {
//		return nil,err
//	}
//	cards, err :=
//
//
//}
//
//// Get the unique name, and tags of this tool
//func (d pulseaudioDevice) Identity() (string, []string) {}
//
//// to get identity of factory if needed
//func (d pulseaudioDevice) FactoryIdentity() (string, []string) {}
//
//// for cli mode, display the range/domain as string
//func (d pulseaudioDevice) Domain() string {}
//
//// for cli mode, display the 'value' as string
//func (d pulseaudioDevice) Value() string {}
//
//// for cli mode, update the 'value', if parsable by tool
//func (d *pulseaudioDevice) Set(value string) (string, error) {}
//
