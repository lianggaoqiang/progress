package main

import (
	"github.com/lianggaoqiang/progress"
	"reflect"
	"testing"
	_ "unsafe"
)

func TestInstantiateRepeatedly1(t *testing.T) {
	defer func() {
		if ok := recover(); ok == nil {
			t.Error("no panic when instantiate repeatedly")
		}
	}()
	progress.Start()
	progress.Start()
}

func TestInstantiateRepeatedly2(t *testing.T) {
	defer func() {
		if ok := recover(); ok == nil {
			t.Error("no panic when instantiate repeatedly")
		}
	}()
	progress.Start()
	progress.StartWithFlag(0)
}

func TestAutoStop(t *testing.T) {
	// p := progress.Start()
	// b := progress.NewBar()
	// p.AddBar(b)
	// for i := 0; i <= 100; i++ {
	// 	b.Inc()
	// }
	// if p.Stop() == nil {
	// 	t.Error("auto stopping feature is invalid")
	// }
}

//go:linkname parseSteps github.com/lianggaoqiang/progress.parseSteps
func parseSteps(steps []string)
func TestStepParsed(t *testing.T) {
	s1 := []string{"", ".", "-.", "-."}
	parseSteps(s1)
	isEqual(t, s1, []string{"", ".", "..", "..."})

	s2 := []string{"Loading", ".-", ".-", "-."}
	parseSteps(s2)
	isEqual(t, s2, []string{"Loading", ".Loading", "..Loading", "..Loading."})

	s3 := []string{"Loading", "---", "---", "---"}
	parseSteps(s3)
	isEqual(t, s3, []string{"Loading", "-Loading", "--Loading", "---Loading"})

	s4 := []string{"Loading", "-(--)", "-(--)", "-(--)"}
	parseSteps(s4)
	isEqual(t, s4, []string{"Loading", "Loading-", "Loading--", "Loading---"})

	s5 := []string{"abc", ".-.", ".-.", ".-."}
	parseSteps(s5)
	isEqual(t, s5, []string{"abc", ".abc.", "..abc..", "...abc..."})
}
func isEqual(t *testing.T, s, expect []string) {
	if !reflect.DeepEqual(s, expect) {
		t.Errorf("custom syntax parsing error!\n--- get: %v\n--- expect: %v", s, expect)
	}
}
