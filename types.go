package progress

import (
	"sync"
)

type Bar interface {
	Hide()
	Show()
	kind() string
	IsHidden() bool
}

// Conf includes all customizable configurations
type Conf struct {
	Debug    bool // if set to true, the percent value displayed will be able to exceed 100%
	AutoStop bool // if set to true, this Progress will be set as a stopped state when all Bars' N-property >= 100. Need: Debug=false
}

var (
	mtx        sync.Mutex
	ins        *Progress         // the global unique instance of Progress
	mode       uint8             // the flag will be set when StartWithFlag is called
	renderSig  chan int          // the signal to render contents of progress bar
	couldPrint = make(chan bool) // the signal to ensure printing this time is ended
)

// define flags
const (
	// HideCursor DisableInput ResizeReactively are flags that inherit from single-line-print
	// see more detail: https://github.com/lianggaoqiang/single-line-print
	HideCursor       uint8 = 0x01 // if set, the cursor will be hidden during printing or writing
	DisableInput     uint8 = 0x02 // if set, input will be disabled during printing or writing
	ResizeReactively uint8 = 0x04 // if set, terminal window size will be got before each printing or writing

	PercentOverflow uint8 = 0x10 // if set, the percent value displayed will be able to exceed 100%
	AutoStop        uint8 = 0x08 // if set, Progress will be set as a stopped state automatically
	// when all Bars' N-property >= 100. it requires PercentOverflow flag not be set.
)

// define colors
const (
	Black uint8 = 30 + iota
	Red
	Green
	Yellow
	Blue
	Purple
	Cyan
	White
)
