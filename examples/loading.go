package main

import (
	"github.com/lianggaoqiang/progress"
	"time"
)

func main() {
	p := progress.Start()

	// TextBar with green text
	tb := progress.NewTextBar(progress.GreenText("start successfully!"))
	p.AddBar(tb)

	// LoadingBar with RedColor
	lb := progress.NewLoadingBar(300, "Loading", "Loading.", "Loading..", "Loading...")
	p.AddBar(lb.SetColor(progress.Red))

	<-time.After(time.Second * 3)
	lb.Hide()

	p.AddBar(progress.NewTextBar("Done, exiting!"))
	p.Stop()
}
