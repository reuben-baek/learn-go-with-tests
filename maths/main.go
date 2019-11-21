package main

import (
	"fmt"
	"github.com/reuben-baek/learn-go-with-tests/maths/clockface/impl/simple_clock"
	"github.com/reuben-baek/learn-go-with-tests/maths/clockface/impl/svg_clock"
	"github.com/reuben-baek/learn-go-with-tests/maths/clockface/model"
	"os"
	"time"
)

func main() {
	var clock model.Clock = svg_clock.SvgClock{os.Stdout}
	clock.Show(time.Now())

	fmt.Println()

	clock = simple_clock.SimpleClock{os.Stdout}
	clock.Show(time.Now())
}
