package main

import (
	"github.com/lianggaoqiang/progress"
	"time"
)

func main() {
	p := progress.Start()

	// create a custom bar
	b1 := progress.NewBar().Custom(progress.BarSetting{
		Total:           50,
		StartText:       "[",
		EndText:         "]",
		PassedText:      "-",
		FirstPassedText: ">",
		NotPassedText:   "=",
	})

	// create a custom inline bar
	b2 := progress.NewBar().Custom(progress.BarSetting{
		UseFloat:        true,
		Inline:          true,
		StartText:       "|",
		EndText:         "|",
		FirstPassedText: ">",
		PassedText:      "=",
		NotPassedText:   " ",
	})

	// create a custom bar with emoji character
	b3 := progress.NewBar().Custom(progress.BarSetting{
		LeftSpace:     10,
		Total:         10,
		StartText:     "|",
		EndText:       "|",
		PassedText:    "⚡",
		NotPassedText: "  ",
	})

	// add bars in progress
	p.AddBar(b2)
	p.AddBar(b3)
	p.AddBar(b1)

	for i := 0; i <= 100; i++ {
		b1.Inc()
		b2.Add(1.4)
		b3.Percent(float64(i))
		time.Sleep(time.Millisecond * 40)
	}
}
