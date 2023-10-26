package progress

import (
	"fmt"
	tw "golang.org/x/text/width"
	"strconv"
)

// get the width of string in terminal
func getWidth(s string) (n int) {
	for _, c := range s {
		switch tw.LookupRune(c).Kind() {
		case tw.EastAsianFullwidth, tw.EastAsianWide:
			n += 2
		case tw.EastAsianHalfwidth, tw.EastAsianNarrow,
			tw.Neutral, tw.EastAsianAmbiguous:
			n += 1
		}
	}
	return
}

// generate ANSI control characters
func esc(suffix ...string) (res string) {
	for _, s := range suffix {
		res += fmt.Sprintf("%c[%s", 033, s)
	}
	return res
}

// ColorText generates string that describe colored text
func ColorText(s string, color uint8) string {
	colorNum := strconv.Itoa(int(color))
	return esc(colorNum+"m") + s + esc("0m")
}

// BlackText generate string that describes black text
func BlackText(s string) string {
	return ColorText(s, Black)
}

// RedText generate string that describes red text
func RedText(s string) string {
	return ColorText(s, Red)
}

// GreenText generate string that describes green text
func GreenText(s string) string {
	return ColorText(s, Green)
}

// YellowText generate string that describes yellow text
func YellowText(s string) string {
	return ColorText(s, Yellow)
}

// BlueText generate string that describes blue text
func BlueText(s string) string {
	return ColorText(s, Blue)
}

// PurpleText generate string that describes purple text
func PurpleText(s string) string {
	return ColorText(s, Purple)
}

// CyanText generate string that describes cyan text
func CyanText(s string) string {
	return ColorText(s, Cyan)
}

// WhiteText generate string that describes white text
func WhiteText(s string) string {
	return ColorText(s, White)
}
