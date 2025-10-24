package ui

//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
// "fyne.io/fyne/v2"
// "fyne.io/fyne/v2/app"
// "fyne.io/fyne/v2/container"
// "fyne.io/fyne/v2/widget"

//)
//
//// Define your interface
//type ProgressBar interface {
//	GetName() string
//	GetValue() int // 0..100
//	OnUp()
//	OnDown()
//}
//
//// CreateProgressWidget builds a vertical progress bar with + and - buttons
//func CreateProgressWidget(p ProgressBar) fyne.CanvasObject {
//	// Title label
//	title := widget.NewLabel(p.GetName())
//
//	// Progress bar
//	bar := widget.NewProgressBar()
//	bar.Min = 0
//	bar.Max = 100
//	bar.SetValue(float64(p.GetValue()))
//
//	// Buttons
//	upButton := widget.NewButton("+", func() {
//		p.OnUp()
//		bar.SetValue(float64(p.GetValue()))
//	})
//	downButton := widget.NewButton("-", func() {
//		p.OnDown()
//		bar.SetValue(float64(p.GetValue()))
//	})
//
//	// Layout vertically: + on top, bar in middle, - at bottom
//	return container.NewVBox(
//		title,
//		upButton,
//		bar,
//		downButton,
//	)
//}
//
//a := app.New()
//w := a.NewWindow("Hello")

//w.SetContent(HelloWindow())
//w.SetContent(ProgressBarX())

//w.ShowAndRun()

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
