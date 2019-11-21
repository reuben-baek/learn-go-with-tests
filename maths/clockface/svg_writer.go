package clockface

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 70
	hourHandLength   = 50
	clockCentreX     = 150
	clockCentreY     = 150
)

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	minuteHand(w, t)
	hourHand(w, t)
	io.WriteString(w, svgEnd)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func secondHand(w io.Writer, t time.Time) {
	p := secondHandPoint(t)
	p = makeHand(secondHandLength, p)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func minuteHand(w io.Writer, t time.Time) {
	p := minuteHandPoint(t)
	p = makeHand(minuteHandLength, p)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func hourHand(w io.Writer, t time.Time) {
	p := hourHandPoint(t)
	p = makeHand(hourHandLength, p)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(length float64, p Point) Point {
	p = Point{p.X * length, p.Y * length}                // scale
	p = Point{p.X, -p.Y}                                 // flip
	return Point{p.X + clockCentreX, p.Y + clockCentreY} // translate
}

func secondHandPoint(tm time.Time) Point {
	angle := secondsInRadians(tm)
	return angleToPoint(angle)
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}

func hourHandPoint(t time.Time) Point {
	angle := hoursInRadians(t)
	return angleToPoint(angle)
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
