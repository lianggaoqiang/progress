package main

import (
	"github.com/lianggaoqiang/progress"
	"time"
)

// this file is used to test package building process
func main() {
	p := progress.Start()

	bar1 := progress.NewBar()
	p.AddBar(bar1)

	bar2 := progress.NewTextBar("This is a text bar!")
	p.AddBar(bar2)

	bar3 := progress.NewLoadingBar(300, ".", "-.", "-.")
	bar3.Custom(progress.LoadingBarSetting{
		FixedWidth: true,
		StartText:  "Loading",
	}).Hide()
	p.AddBar(bar3)

	go func() {
		<-time.After(time.Millisecond * 1500)
		bar3.Show()
	}()

	for i := 0; i <= 100; i++ {
		bar1.Percent(float64(i))
		time.Sleep(time.Millisecond * 30)
	}
}
