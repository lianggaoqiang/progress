package progress

import (
	"errors"
	"fmt"
	slp "github.com/lianggaoqiang/single-line-print"
	"math"
	"strings"
)

type Progress struct {
	bars []Bar
}

// StartWithFlag instantiate an instance of Progress with flags and returns the instance
func StartWithFlag(flag uint8) *Progress {
	mtx.Lock()
	defer mtx.Unlock()

	// duplicate instantiation
	if ins != nil {
		renderSig <- 3
		<-closeDone
		panic("Progress had already been instantiated, you can use Stop to unregister the instance")
	}

	mode = flag
	ins = &Progress{}
	go ins.render()
	return ins
}

// Start returns an instance pointer of Progress
func Start() *Progress {
	return StartWithFlag(AutoStop | HideCursor | DisableInput)
}

// AddBar adds a Bar into Progress
func (p *Progress) AddBar(b Bar) {
	p.bars = append(p.bars, b)

	// if b is TextBar, render once directly
	if _, ok := b.(*TextBar); ok {
		renderSig <- 1
		<-couldPrint
	}
}

// Stop will unregister Progress and close renderSig channel
func (p *Progress) Stop() error {
	mtx.Lock()
	defer mtx.Unlock()
	if ins != nil {
		renderSig <- 3
		<-closeDone
		return nil
	}
	return errors.New("progress is already in closed state")
}

// render the progress bar
func (p *Progress) render() {
	pt := slp.NewPrinterWithFlag(mode)
	defer func() {
		// stop all loading bars
		for _, IBar := range p.bars {
			if bar, ok := IBar.(*LoadingBar); ok {
				if bar.active {
					bar.Stop()
				}
			}
		}
		ins = nil
		pt.Stop()
		closeDone <- true
	}()

	for {
		code := <-renderSig
		if code == 1 || code == 2 {
			var content string
			var isInline, isTextBarInit bool
			for _, iBar := range p.bars {
				// if this bar is hidden, jump into next loop directly
				if iBar.IsHidden() {
					continue
				}

				// add static content directly if bar is a TextBar
				if bar, ok := iBar.(*TextBar); ok {
					content += bar.staticContent
					isInline = bar.Setting.Inline
					if bar.isInit {
						isTextBarInit = true
						bar.isInit = false
					}
				}

				// splice the contents of LoadingBar
				if bar, ok := iBar.(*LoadingBar); ok {
					curStr := bar.loadingText[bar.loadingTextIndex]
					content += bar.Setting.StartText + curStr
					if code == 2 {
						if bar.loadingTextIndex == len(bar.loadingText)-1 {
							bar.loadingTextIndex = 0
						} else {
							bar.loadingTextIndex++
						}
					}
					if bar.Setting.FixedWidth {
						content += strings.Repeat(" ", bar.maxWidth-getWidth(curStr))
					}
					content += bar.Setting.EndText
					isInline = bar.Setting.Inline
				}

				// splice the contents of DefaultBar
				if bar, ok := iBar.(*DefaultBar); ok {
					bar.lastN = bar.N
					barN := math.Min(bar.N, 100)
					doneN := int(barN * float64(bar.Setting.Total) / 100)

					// splice progress bar
					content += strings.Repeat(" ", int(bar.Setting.LeftSpace))
					content += bar.Setting.StartText
					if doneN == 1 {
						content += bar.Setting.FirstPassedText
					} else if doneN > 1 {
						content += strings.Repeat(bar.Setting.PassedText, doneN-1) + bar.Setting.FirstPassedText
					}
					content += strings.Repeat(bar.Setting.NotPassedText, int(bar.Setting.Total)-doneN)
					content += bar.Setting.EndText

					// splice the contents of percent
					if !bar.Setting.HidePercent {
						if bar.Setting.UseFloat {
							percentStr := fmt.Sprintf("%.2f", bar.N) + "%"
							content += " " + ColorText(percentStr, bar.Setting.PercentColor) + strings.Repeat(" ", 7-len(percentStr))
						} else {
							percentStr := fmt.Sprintf("%d", int(bar.N)) + "%"
							content += " " + ColorText(percentStr, bar.Setting.PercentColor) + strings.Repeat(" ", 4-len(percentStr))
						}
					}
					content += strings.Repeat(" ", int(bar.Setting.RightSpace))
					isInline = bar.Setting.Inline
				}

				if !isInline {
					content += "\n"
				}
			}
			pt.Print(content)

			// determine whether to end listening when AutoStop flag is set
			progressShouldBeEnd := true
			if mode&PercentOverflow == 0 && mode&AutoStop != 0 {
				for _, iBar := range p.bars {
					if _, ok := iBar.(*TextBar); ok && isTextBarInit {
						progressShouldBeEnd = false
					}
					if bar, ok := iBar.(*LoadingBar); ok && bar.active {
						progressShouldBeEnd = false
					}
					if bar, ok := iBar.(*DefaultBar); ok && bar.N < 100 {
						progressShouldBeEnd = false
					}
					if !progressShouldBeEnd {
						break
					}
				}
			} else {
				progressShouldBeEnd = false
			}

			couldPrint <- true
			if progressShouldBeEnd {
				return
			}
		} else {
			return
		}
	}
}
