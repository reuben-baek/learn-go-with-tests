package simple_clock

import (
	"io"
	"time"
)

type SimpleClock struct {
	Writer io.Writer
}

func (s SimpleClock) Show(t time.Time) {
	s.Writer.Write([]byte(t.String()))
}
