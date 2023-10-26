package progress

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

type LoadingBar struct {
	maxWidth         int
	active           bool
	stopSig          chan bool
	Setting          LoadingBarSetting
	loadingTextIndex int      // the index of current displayed string in Bar.loadingText
	loadingText      []string // the strings set of LoadingBar, they will be displayed regularly and orderly
}

// LoadingBarSetting describes the setting of LoadingBar
type LoadingBarSetting struct {
	Hidden             bool
	Inline             bool
	StartText, EndText string // text at start and end place
	FixedWidth         bool   // if set to true, the LoadingBar will use a fixed width (equal to the length of the longest string)
}

// enable LoadingBar implement interface Bar
func (b *LoadingBar) kind() string {
	return "loading"
}
func (b *LoadingBar) Hide() {
	b.Setting.Hidden = true
}
func (b *LoadingBar) Show() {
	b.Setting.Hidden = false
}
func (b *LoadingBar) IsHidden() bool {
	return b.Setting.Hidden
}

// NewLoadingBar returns an instance pointer of a loading bar
func NewLoadingBar(milli time.Duration, steps ...string) *LoadingBar {
	// process the format of steps
	parseSteps(steps)

	ret := &LoadingBar{
		loadingTextIndex: 0,
		active:           true,
		loadingText:      steps,
		stopSig:          make(chan bool),
	}
	for _, v := range steps {
		if ret.maxWidth < getWidth(v) {
			ret.maxWidth = getWidth(v)
		}
	}
	go func() {
		ticker := time.NewTicker(time.Millisecond * milli)
		for {
			select {
			case <-ticker.C:
				renderSig <- 2
				<-couldPrint
			case <-ret.stopSig:
				ticker.Stop()
				return
			}
		}
	}()
	return ret
}

// parse the step strings of LoadingBar base on custom syntax
func parseSteps(steps []string) {
	gap := fmt.Sprintf("%c", 0x0)
	for i := 1; i < len(steps); i++ {
		steps[i] = strings.ReplaceAll(steps[i], "(--)", gap)
		steps[i] = strings.ReplaceAll(steps[i], "--", gap)
		steps[i] = strings.ReplaceAll(steps[i], "-", steps[i-1])
		steps[i] = strings.ReplaceAll(steps[i], gap, "-")
	}
}

// SetColor will change the color of all strings in LoadingBar.loadingText
func (b *LoadingBar) SetColor(color uint8) *LoadingBar {
	for i, v := range b.loadingText {
		b.loadingText[i] = ColorText(v, color)
	}
	return b
}

// Custom changes the setting of LoadingBar
func (b *LoadingBar) Custom(setting LoadingBarSetting) *LoadingBar {
	b.Setting = setting
	return b
}

// Stop clean the timer of LoadingBar
func (b *LoadingBar) Stop() error {
	if b.active {
		b.active = false
		b.stopSig <- true
	} else {
		return errors.New("LoadingBar is already in closed state")
	}
	return nil
}
