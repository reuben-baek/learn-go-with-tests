package svg_clock

import (
	"math"
	"time"
)

func secondsInRadians(t time.Time) float64 {
	sec := t.Second()
	return math.Pi / (30 / float64(sec))
}

func minutesInRadians(t time.Time) float64 {
	m := t.Minute()
	return (secondsInRadians(t) / 60) + math.Pi/(30/float64(m))
}

func hoursInRadians(t time.Time) float64 {
	h := t.Hour()
	return (minutesInRadians(t) / 12) + math.Pi/(6/float64(h%12))
}
