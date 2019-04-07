package omikuji

import (
	"fmt"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/errors"
	"math/rand"
	"sync"
	"time"
)

// Interfaces
type Randomizer interface {
	GetRandom(min, max int) Omikuji
	GetMax() int
	GetDaikichiMin() int
	GetNoDaikichiMin() int
}

// Types and implementations
type Omikuji struct {
	Text string `json:"omikuji"`
}

type AllOmikujis []Omikuji

func (omikujis *AllOmikujis) GetRandom(min, max int) Omikuji {
	if min < 0 || max > omikujis.GetMax() || min >= max {
		msg := fmt.Sprintf("Invalid arguments: min=%d, max=%d", min, max)
		errors.ThrowOmikujiException(msg, errors.OmikujiErrorCode)
	}

	randIndex := min + rand.Intn(max-min)
	return (*omikujis)[randIndex]
}

func (omikujis *AllOmikujis) GetMax() int {
	return len(*omikujis)
}

func (omikujis *AllOmikujis) GetDaikichiMin() int {
	return 0
}

func (omikujis *AllOmikujis) GetNoDaikichiMin() int {
	return 1
}

// Singleton all omikujis
var allOmikujis AllOmikujis
var once sync.Once

// Get all omikujis singleton instance
func GetOmikujiRandomizer() Randomizer {
	// Using once here for thread safety
	// http://marcio.io/2015/07/singleton-pattern-in-go/
	once.Do(func() {
		// Initialize this var only once
		allOmikujis = []Omikuji{
			{"大吉"},
			{"中吉"},
			{"小吉"},
			{"吉"},
			{"半吉"},
			{"末吉"},
			{"末小吉"},
			{"凶"},
			{"小凶"},
			{"半凶"},
			{"末凶"},
			{"大凶"},
		}
		// Initialize random generator only once
		rand.Seed(time.Now().UTC().UnixNano())
	})
	return &allOmikujis
}
