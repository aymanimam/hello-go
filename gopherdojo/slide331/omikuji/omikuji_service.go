package omikuji

import (
	"fmt"
	"github.com/aymanimam/hello-go/gopherdojo/slide331/errors"
	"time"
)

// Interfaces
type PeriodChecker interface {
	WithinThePeriod(t time.Time) bool
}

// Types
type PeriodicDate struct {
	Month time.Month
	Day   int
}

type Period struct {
	From PeriodicDate
	To   PeriodicDate
}

func (p *Period) WithinThePeriod(t time.Time) bool {
	m := t.Month()
	d := t.Day()

	if m >= p.From.Month && m <= p.To.Month {
		if d >= p.From.Day && d <= p.To.Day {
			return true
		}
	}

	return false
}

type Dispatcher interface {
	GetNextOmikuji() Omikuji
}

type Service struct {
	PeriodChecker PeriodChecker
	Randomizer    Randomizer
}

func (s *Service) GetNextOmikuji() Omikuji {
	if s.Randomizer == nil || s.PeriodChecker == nil {
		msg := fmt.Sprintf("One or more invalid arguments! [Randomizer: %v][PeriodChecker: %v]",
			s.Randomizer, s.PeriodChecker)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	}

	r := s.Randomizer
	currentTime := time.Now()

	if s.PeriodChecker.WithinThePeriod(currentTime) {
		return r.GetRandom(r.GetDaikichiMin(), r.GetMax())
	} else {
		return r.GetRandom(r.GetNoDaikichiMin(), r.GetMax())
	}
}

// Get PeriodChecker
func GetPeriodChecker(fromDate, toDate PeriodicDate) PeriodChecker {
	if fromDate.Month > toDate.Month {
		msg := fmt.Sprintf("Period checker inputs are invalid [fromDate: %v][toDate: %v]", fromDate, toDate)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	} else if fromDate.Month == toDate.Month {
		if fromDate.Day > toDate.Day {
			msg := fmt.Sprintf("Period checker inputs are invalid [fromDate: %v][toDate: %v]", fromDate, toDate)
			errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
		}
	}
	return &Period{From: fromDate, To: toDate}
}

// Get Dispatcher
func GetOmikujiDispatcher(pc PeriodChecker) Dispatcher {
	return &Service{
		pc,
		GetOmikujiRandomizer(),
	}
}

// Business logic
/*func GetNextOmikuji(r Randomizer, pc PeriodChecker) Omikuji {
	if r == nil || pc == nil {
		msg := fmt.Sprintf("One or more invalid arguments! [Randomizer: %v][PeriodChecker: %v]", r, pc)
		errors.ThrowOmikujiException(msg, errors.OmikujiServiceErrorCode)
	}

	currentTime := time.Now()

	if pc.WithinThePeriod(currentTime) {
		return r.GetRandom(r.GetDaikichiMin(), r.GetMax())
	} else {
		return r.GetRandom(r.GetNoDaikichiMin(), r.GetMax())
	}
}
*/
