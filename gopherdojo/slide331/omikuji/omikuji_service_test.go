package omikuji

import (
	"testing"
	"time"
)

// Mock Randomizer
type MockRandomizer struct{}

func (omikujis *MockRandomizer) GetRandom(min, max int) Omikuji {
	if min == 0 {
		return Omikuji{"大吉"}
	} else {
		return Omikuji{"吉"}
	}
}

func (omikujis *MockRandomizer) GetMax() int {
	return 2
}

func (omikujis *MockRandomizer) GetDaikichiMin() int {
	return 0
}

func (omikujis *MockRandomizer) GetNoDaikichiMin() int {
	return 1
}

// ---------------

func TestGetNextOmikujiNilArgs(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestGetNextOmikujiNilConfigObj should have panicked!")
			}
		}()
		// This function should cause a panic
		GetNextOmikuji(nil, nil)
	}()
}

func TestGetNextOmikujiWithDaikichi(t *testing.T) {
	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentDay := currentTime.Day()

	fromDate := PeriodicDate{Month: currentMonth, Day: currentDay}
	toDate := PeriodicDate{Month: currentMonth, Day: currentDay}
	period := &Period{From: fromDate, To: toDate}

	omikuji := GetNextOmikuji(&MockRandomizer{}, period)

	if omikuji.Text != "大吉" {
		t.Error(`Expected "大吉" omikuji! But was [`, omikuji.Text, `]`)
	}
}

func TestGetNextOmikujiNoDaikichi(t *testing.T) {
	futureTime := time.Now().AddDate(0, 1, 0)
	futureMonth := futureTime.Month()
	futureDay := futureTime.Day()

	fromDate := PeriodicDate{Month: futureMonth, Day: futureDay}
	toDate := PeriodicDate{Month: futureMonth, Day: futureDay}
	period := &Period{From: fromDate, To: toDate}

	omikuji := GetNextOmikuji(&MockRandomizer{}, period)

	if omikuji.Text != "吉" {
		t.Error(`Expected "吉" omikuji! But was [`, omikuji.Text, `]`)
	}
}
