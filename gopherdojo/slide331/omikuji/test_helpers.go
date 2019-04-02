package omikuji

import (
	"testing"
)

// Is the given array has the given omikuji
func Contains(arr []Omikuji, o Omikuji) bool {
	for _, e := range arr {
		if e.Text == o.Text {
			return true
		}
	}
	return false
}

// Function under test
type AssertPanicFunc func()

func AssertPanic(t *testing.T, failMessage string, function AssertPanicFunc) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf(failMessage)
			}
		}()
		// This function should cause a panic
		function()
	}()
}
