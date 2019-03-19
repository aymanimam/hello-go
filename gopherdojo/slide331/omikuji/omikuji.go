package omikuji

import (
	"math/rand"
	"sync"
	"time"
)

// Constants
const OMIKUJI_WITH_DAIKICHI_MIN = 0
const OMIKUJI_NO_DAIKICHI_MIN = 1
const OMIKUJI_MAX = 13

// Types
type Omikuji struct {
	text string `json:"omikuji"`
}

type AllOmikujis struct {
	omikujis []Omikuji
}

type Randomizer interface {
	Init()
	GetRandom(min, max int32) interface{}
}

// Singleton all omikujis
var allOmikujis *AllOmikujis
var once sync.Once

// Get all omikujis singleton instance
func GetAllOmikujis() *AllOmikujis {
	// Using once here for thread safety
	// http://marcio.io/2015/07/singleton-pattern-in-go/
	once.Do(func() {
		allOmikujis = &AllOmikujis{[]Omikuji{
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
		}}
	})
	return allOmikujis
}

func (omikujis *AllOmikujis) Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func (omikujis *AllOmikujis) GetRandom(min int, max int) Omikuji {
	randIndex := min + rand.Intn(max-min)
	return omikujis.omikujis[randIndex]
}
