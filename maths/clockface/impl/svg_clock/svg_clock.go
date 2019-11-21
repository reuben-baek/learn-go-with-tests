package svg_clock

import (
	"io"
	"time"
)

type SvgClock struct {
	Writer io.Writer
}

func (s SvgClock) Show(t time.Time) {
	SVGWriter(s.Writer, t)
}
