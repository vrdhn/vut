package main

import (
	// "fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/widget"
	"fmt"
	"os"
)

func usage() {

}
func main() {
	StartLogger()
	args := os.Args[1:]
	switch len(args) {
	case 0:
		fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
			"Device  ", "       Name", "Value", "Range")
		for _, t := range GetTools() {
			fn, _ := t.factoryIdentity()
			tn, _ := t.identity()
			fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
				fn, tn, t.value(), t.domain())
		}
	case 1:
		fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
			"Device  ", "       Name", "Value", "Range")
		for _, t := range GetTool(args[0]) {
			fn, _ := t.factoryIdentity()
			tn, _ := t.identity()
			fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
				fn, tn, t.value(), t.domain())
		}
	case 2:
		fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
			"Device  ", "       Name", "Value", "Range")
		for _, t := range GetTool(args[0]) {
			t.set(args[1])
		}
		for _, t := range GetTool(args[0]) {
			fn, _ := t.factoryIdentity()
			tn, _ := t.identity()
			fmt.Printf("%10s  |  %-20s  | %5s |  %s\n",
				fn, tn, t.value(), t.domain())
		}
	default:
		usage()
	}

	//a := app.New()
	//w := a.NewWindow("Hello")

	//w.SetContent(HelloWindow())
	//w.SetContent(ProgressBarX())

	//w.ShowAndRun()

}

//func HelloWindow() fyne.CanvasObject {
//	hello := widget.NewLabel("Hello Fyne!")
//	return container.NewVBox(
//		hello,
//		widget.NewButton("Hi!", func() {
//			hello.SetText("Welcome :)")
//		}),
//	)
//}
//
//func ProgressBarX() fyne.CanvasObject {
//	my := &MyProgress{name: "Loading", value: 40}
//	w := CreateProgressWidget(my)
//	ret := widget.NewCard(
//		"Volume", // title
//		"Control",
//		w)
//	ret.SetContent(w)
//	return ret
//}
//
//
//
//// Example implementation
//type MyProgress struct {
//	name  string
//	value int
//}
//
//func (p *MyProgress) GetName() string { return p.name }
//func (p *MyProgress) GetValue() int   { return p.value }
//func (p *MyProgress) OnUp() {
//	if p.value < 100 {
//		p.value += 10
//	}
//}
//func (p *MyProgress) OnDown() {
//	if p.value > 0 {
//		p.value -= 10
//	}
//}
//
