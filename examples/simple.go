package main

import (
	"github.com/lianggaoqiang/progress"
	"time"
)

func main() {
	p := progress.Start()
	b := progress.NewBar()
	p.AddBar(b)

	for i := 0; i <= 100; i++ {
		b.Inc()
		time.Sleep(time.Millisecond * 10)
	}
}
