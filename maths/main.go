package main

import (
	"github.com/reuben-baek/learn-go-with-tests/maths/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
