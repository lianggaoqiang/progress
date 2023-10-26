package main

import (
	"github.com/lianggaoqiang/progress"
	"time"
)

func main() {
	p := progress.Start()

	b := progress.NewBar().Custom(progress.BarSetting{
		PassedText:    progress.ColorText("▇", progress.Green),
		NotPassedText: progress.WhiteText("▇"),
	})
	p.AddBar(b)

	for i := 0; i <= 100; i++ {
		b.Inc()
		time.Sleep(time.Millisecond * 50)
	}
}
