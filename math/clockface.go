package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength = 50
	clockCentreX = 150
	clockCentreY = 150
)

type Point struct {
	X float64
	Y float64
}

type SVG struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y1 float64 `xml:"y1,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func SecondHand(t time.Time) Point {
	p := SecondHandPoint(t)
	p = Point{X: p.X * secondHandLength, Y: p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	return p
}

func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func SecondHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	secondHand(w, t)
	MinuteHand(w, t)
	HourHand(w, t)
	io.WriteString(w, svgEnd)
}

func secondHand(w io.Writer, t time.Time) {
	p := SecondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff; stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func MinutesInRadians(t time.Time) float64 {
	return (SecondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func MinuteHandPoint(t time.Time) Point {
	angle := MinutesInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func MinuteHand(w io.Writer, t time.Time) {
	p := MinuteHandPoint(t)
	p = Point{p.X * minuteHandLength, p.Y * minuteHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func HoursInRadians(t time.Time) float64 {
	return (MinutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour() % 12)))
}

func HourHandPoint(t time.Time) Point{
	angle := HoursInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func HourHand(w io.Writer, t time.Time){
	p := HourHandPoint(t)
	p = Point{p.X * hourHandLength, p.Y * hourHandLength}
	p = Point{p.X, -p.Y}
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}
