package main

//import (
//	"fyne.io/fyne/v2"
//	"fyne.io/fyne/v2/container"
//	"fyne.io/fyne/v2/widget"
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
