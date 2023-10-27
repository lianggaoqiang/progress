# progress [![](https://camo.githubusercontent.com/315a8800fc96d3c5b32e13227b10500ef850688793cc6664418d018980eb3cb4/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f676f737572692f75696c6976653f7374617475732e737667)](https://pkg.go.dev/github.com/lianggaoqiang/progress) [![https://github.com/lianggaoqiang/progress/blob/main/LICENSE](https://img.shields.io/badge/license-MIT-red.svg)](https://github.com/lianggaoqiang/progress/blob/main/LICENSE) [![](https://github.com/lianggaoqiang/progress/actions/workflows/ci.yml/badge.svg)](https://github.com/lianggaoqiang/progress/actions/workflows/ci.yml)

A thread-safe progress bar printing program, which is designed for printing multiple bars simultaneously with a highly efficient. This package is powered by [single-line-print](https://github.com/lianggaoqiang/single-line-print)

This Golang package is cross-platform which works perfectly on Linux, Windows and MacOS. It provides many very simple apis to render the progress bar with diverse styles that you want, in addition, you can also combine these different styles of progress bars in any format.
ssss

<br>

## Install
```shell
go get github.com/lianggaoqiang/progress
```

<br>

## Basic usage

```go
import "github.com/lianggaoqiang/progress"

func main(){
	p := progress.Start()
	bar := progress.NewBar()
	p.AddBar(bar)
	
	for i := 0; i <= 100; i++ {
		bar.Inc()
		time.Sleep(time.Millisecond * 10)
	}
}
```

<img src="https://github.com/lianggaoqiang/progress/blob/main/doc/simple-bar.gif" style="width:70%" />

<br>

## Customization

If the example above is not the style you want, you can use method `Custom` to get progress bar with custom style.

```go
import "github.com/lianggaoqiang/progress"

func main() {
	p := progress.Start()
	b := progress.NewBar().Custom(
		progress.BarSetting{
			StartText:       "[",
			EndText:         "]",
			PassedText:      "-",
			FirstPassedText: ">",
			NotPassedText:   "=",
		},
	)
	p.AddBar(b)

	for i := 0; i <= 100; i++ {
		b.Percent(float64(i))
		time.Sleep(time.Millisecond * 30)
	}
}
```

<img src="https://github.com/lianggaoqiang/progress/blob/main/doc/custom-bar.gif" style="width:70%;border-radius:6px;" />

The `BarSetting` structure contains rich customization related properties to modify the style of progress bar:

- LeftSpace, RightSpace(string): the spacings at left and right
- Total(uint16): the total count of inner characters of the progress bar
- Hidden(bool): if set to true, the bar will be hidden and not rendered
- Inline(bool): if set to true, there will not be a line break at the end place of the bar
- HidePercent(bool): if set to true, there will be a percent value be displayed at the end place of progress bar
- UseFloat(bool): if set to true, progress bar will have a precision of 0.01%
- StartText, EndText(bool): text at start and end place
- PassedText, NotPassedText(string): text passed and not passed
- FirstPassedText(string): the first passed text
- PercentColor(uint8): the color of percent value, accept a value like progress.Green, you can see more details at [Change the color](#change-the-color)

<br>

## Change the color

If you want to change the color of the progress bar, you can use built-in color realted method to generate ANSI color characters:

```go
import "github.com/lianggaoqiang/progress"

// Both progress.ColorText(str, progress.Xxx) and progress.XxxText(str) is okay
// Xxx may be Black, Red, Green, Yellow, Blue, Purple, Cyan, White
func main() {
	p := progress.Start()
	b := progress.NewBar().Custom(progress.BarSetting{
		NotPassedText: progress.WhiteText("▇"),
		PassedText:    progress.ColorText("▇", progress.Green),
	})
	p.AddBar(b)

	for i := 0; i <= 100; i++ {
		b.Inc()
		time.Sleep(time.Millisecond * 50)
	}
}
```

<img src="https://github.com/lianggaoqiang/progress/blob/main/doc/color-bar.gif" style="width:70%" />

<br>

## LoadingBar, TextBar

The above examples have demonstrated the method of how to utilize the default progress bar. However this package includes not only default progress bar but also two additional UI formats: TextBar and LoadingBar.

```go
import "github.com/lianggaoqiang/progress"

func main(){
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
```

<img src="https://github.com/lianggaoqiang/progress/blob/main/doc/loading-bar.gif" style="width:70%" />

The first parameter of NewLoadingBar is the interval of each render, and the remaining parameters are the texts will be printed at each render. You may have noticed that when we creating a LoadingBar, we need write "Loading" three times. While this is fine for short text, if the text is longer, it can cause significant inconvenience. At this point, we can use hyphen(-) to represent the previous text in steps parameter, just like the following(see more parsing rules of custom syntax at [FAQ](#faq)):

```go
lb := progress.NewLoadingBar(300, "Loading", "-.", "-.", "-.")
```

It should be noted that `TextBarSetting` and `LoadingBarSetting` have fewer properties than `BarSetting`:

+ properties of TextBarSetting:
   - Inline, Hidden (they has the same effect as properties in BarSetting)
+ properties of LoadingBarSetting:
   - Hide, Inline, StartText, EndText (they has the same effect as properties in BarSetting)
   - FixedWidth(bool): if set to true, the LoadingBar will use a fixed width (equal to the length of the longest string)


<br>

## Combine multiple bars

As mentioned at the beginning, you can combine different styles of bar in any format, the following is an example for combining the bars:

```go
import "github.com/lianggaoqiang/progress"

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
```

<img src="https://github.com/lianggaoqiang/progress/blob/main/doc/combine.gif" style="width:70%" />

<br>

## FAQ

1. How could i use actual "-" in NewLoadingBar?
   <br>If you want to use actual "-", use "--" instead. But "---" will be parsed as "-[the previous string]", if you want the last two hyphens be parsed as "-", you could use "()" to change priority, likes following:"-(--)" will be parsed as "[the previous string]-"

2. How to disable auto-stopping feature?
   <br>You could use `progress.StartWithFlag` method to initialize progress. Three are five flags can be choosed(the fist three inherit from [single-line-print](https://github.com/lianggaoqiang/single-line-print)):
   - HideCursor: if set, the cursor will be hidden during printing or writing
   - DisableInput: if set, input will be disabled during printing or writing
   - ResizeReactively: if set, terminal window size will be got before each printing or writing

   - PercentOverflow: if set, the percent value displayed will be able to exceed 100%
   - AutoStop: if set, Progress will automatically stopped when all DefaultBars' percent value >= 100 and all LoadingBars are stopped. it requires PercentOverflow flag not be set.
    ```go
     p := progress.StartWithFlag(progress.HideCursor | progress.DisableInput | progress.AutoStop)
     ```
