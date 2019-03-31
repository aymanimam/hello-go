package omikuji

import (
	"testing"
)

var omikujis = []Omikuji{
	{"大吉"},
	{"吉"},
	{"吉"},
	{"中吉"},
	{"小吉"},
	{"半吉"},
	{"末吉"},
	{"末小吉"},
	{"凶"},
	{"小凶"},
	{"半凶"},
	{"末凶"},
	{"大凶"},
}

func TestGetAllOmikujis(t *testing.T) {
	randomizer := GetOmikujiRandomizer()
	omikujiCount := randomizer.GetMax()
	if omikujiCount != len(omikujis) {
		t.Error(`omikujiCount = `, omikujiCount, `, omikujiCount != 13`)
	}
}

func TestAllOmikujisIsSingleton(t *testing.T) {
	f := func(ch chan Randomizer) {
		randomizer := GetOmikujiRandomizer()
		ch <- randomizer
	}

	randomizerChan1 := make(chan Randomizer)
	randomizerChan2 := make(chan Randomizer)

	go f(randomizerChan1)
	go f(randomizerChan2)

	randomizer1 := <-randomizerChan1
	randomizer2 := <-randomizerChan2

	// Comparing pointers
	if randomizer1 != randomizer2 {
		t.Error(`Two different allOmikujis objects`, randomizer1, randomizer2)
	}
}

func TestGetRandom(t *testing.T) {
	randomizer := GetOmikujiRandomizer()
	omikuji := randomizer.GetRandom(1, 4)
	if !contains(omikujis[1:4], omikuji) {
		t.Error(`This omikuji [`, omikuji, `] is not expected`)
	}
}

func TestGetRandomInvalidArgs(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("TestGetNextOmikujiNilConfigObj should have panicked!")
			}
		}()
		// This function should cause a panic
		randomizer := GetOmikujiRandomizer()
		randomizer.GetRandom(-1, 4)
	}()
}

func contains(s []Omikuji, o Omikuji) bool {
	for _, a := range s {
		if a.Text == o.Text {
			return true
		}
	}
	return false
}