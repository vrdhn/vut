package tools_test

import (
	"testing"

	"vut/tools"
)

func TestBrightness(t *testing.T) {

	tool := tools.BrightnessFactory()

	checks := tool.Check()
	if len(checks) != 0 {
		t.Fatal("Checks failed")
	}

	devices, err := tool.Devices()
	if err != nil {
		panic(err)
	}
	for _, device := range devices {
		old := device.Value()
		device.Set("0")
		val := device.Value()
		device.Set(old)
		if val != "0" {
			name, _ := device.Identity()
			t.Fatalf("didn't got zero back: %s : %s", name, val)
		}
	}
}
