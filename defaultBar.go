package progress

import (
	"errors"
)

type DefaultBar struct {
	N, lastN float64
	Setting  BarSetting
}

// BarSetting describes the setting of Bar
type BarSetting struct {
	PercentColor              uint8  // the color of percent value
	LeftSpace, RightSpace     uint16 // the spacings at left and right
	Total                     uint16 // the total count of inner characters of the progress bar
	Hidden                    bool   // if set to true, the bar will be hidden and not rendered
	Inline                    bool   // if set to true, there will not be a line break at the end place of the bar
	HidePercent               bool   // if set to true, there will be a percent value be displayed at the end place of progress bar
	UseFloat                  bool   // if set to true, progress bar will have a precision of 0.01%
	StartText, EndText        string // text at start and end place
	PassedText, NotPassedText string // text passed and not passed
	FirstPassedText           string // the first passed text
}

// enable DefaultBar implement interface Bar
func (b *DefaultBar) kind() string {
	return "default"
}
func (b *DefaultBar) Hide() {
	b.Setting.Hidden = true
}
func (b *DefaultBar) Show() {
	b.Setting.Hidden = false
}
func (b *DefaultBar) IsHidden() bool {
	return b.Setting.Hidden
}

// NewBar returns an instance pointer of Bar
func NewBar() *DefaultBar {
	return (&DefaultBar{}).Custom(BarSetting{})
}

// Custom changes attributes and display form of Bar
func (b *DefaultBar) Custom(setting BarSetting) *DefaultBar {
	setDefaultValueString(&setting.PassedText, "▇")
	setDefaultValueString(&setting.NotPassedText, " ")
	setDefaultValueString(&setting.FirstPassedText, setting.PassedText)
	if setting.Total == 0 {
		setting.Total = 25
	}
	b.Setting = setting
	return b
}

// used to set the default string value of Bar.Setting
func setDefaultValueString(s *string, v string) {
	if *s == "" {
		*s = v
	}
}

// Percent sets the percent of Bar
func (b *DefaultBar) Percent(n float64) (err error) {
	mtx.Lock()
	defer mtx.Unlock()
	if n > 100 {
		if mode&PercentOverflow == 0 {
			n = 100
		}
		err = errors.New("percent can not be grater than 100")
	}
	if n != b.lastN {
		b.N = n
		renderSig <- 1
		<-couldPrint
	}
	return
}

// Inc will increase Bar.N by 1 and then render the progress bar
func (b *DefaultBar) Inc() error {
	return b.Percent(b.N + 1)
}

// Add can add specified value to Bar.N and then render the progress bar
func (b *DefaultBar) Add(n float64) error {
	return b.Percent(n + b.N)
}
