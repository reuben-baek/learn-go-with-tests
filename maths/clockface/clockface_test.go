package clockface

import (
	"math"
	"time"
)

func roughlyEqualFloat64(a, b float64) bool {
	const equalityThreshold = 1e-7
	return math.Abs(a-b) < equalityThreshold
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) &&
		roughlyEqualFloat64(a.Y, b.Y)
}

func testName(tm time.Time) string {
	return tm.Format("15:04:05")
}

func simpleTime(h int, m int, s int) time.Time {
	return time.Date(1, time.January, 0, h, m, s, 0, time.UTC)
}

func containsLine(line Line, lines []Line) bool {
	for _, l := range lines {
		if l == line {
			return true
		}
	}
	return false
}
