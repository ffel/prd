package prdsymb

import (
	"bytes"
	"fmt"

	"github.com/ajstarks/svgo"
)

var buffer *bytes.Buffer

var canvas *svg.SVG

type mode int

const (
	wait mode = iota // send and receive
	postponed
	immediately
	active // processes
	asleep
)

const (
	d       = 20  // half symbol size
	size    = 4   // line width
	sw      = 2   // off set
	opacity = 0.4 // opacity
)

// http://tutorials.jenkov.com/svg/svg-and-css.html

// draw a send symbol (is a little bit like an "s")
// de basis kan een soort s zijn (een gespiegelde z) en deze kan worden ingevuld
func Send(m mode, x, y int, color string) {
	if m == wait || m == immediately {
		canvas.Polygon([]int{x - d + sw, x + d + sw, x + d + sw}, []int{y - d - sw, y - d - sw, y + d - sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f",
				color, size, opacity))
	}
	if m == postponed || m == immediately {
		canvas.Polygon([]int{x - d - sw, x - d - sw, x + d - sw}, []int{y - d + sw, y + d + sw, y + d + sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:%s;stroke-opacity:%.2f;fill-opacity:%.2f",
				"none", size, color, opacity, opacity))
	}
}

// Receice draws a receive message symbol
func Receive(m mode, x, y int, color string) {
	if m == wait || m == immediately {
		canvas.Polyline([]int{x - d - sw, x - d - sw, x + d + sw}, []int{y + d + sw, y - d - sw, y - d - sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
				color, size, opacity))
	}
	if m == postponed || m == immediately {
		canvas.Square(x-d+3*sw, y-d+3*sw, 2*d-2*sw,
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:%s;stroke-opacity:%.2f;fill-opacity:%.2f",
				"none", size, color, opacity, opacity))
	}
}

// Process draws a horizontal process line
func Process(m mode, x1, x2, y int) {
	if m == active {
		canvas.Line(x1+d+4*sw, y, x2-d-4*sw, y, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
			"black", size, opacity))
	} else if m == asleep {
		canvas.Line(x1+d+4*sw, y, x2-d-4*sw, y, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round;stroke-dasharray:5, 5",
			"black", size, opacity))
	}
}

func Create(x, y1, y2 int) {
	var north, south int

	if y1 > y2 {
		north = y2
		south = y1
	} else {
		north = y1
		south = y2
	}

	canvas.Line(x, north+d+4*sw, x, south-d-4*sw, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round;stroke-dasharray:5, 5",
		"black", size, opacity))

	canvas.Circle(x, y2, 4*sw, fmt.Sprintf("fill-opacity:%.2f;", opacity))
}

// Channel draws a vertical line to display communication over a channel
func Channel(x, y1, y2 int, color string) {

	var north, south int

	if y1 > y2 {
		north = y2
		south = y1
	} else {
		north = y1
		south = y2
	}

	canvas.Line(x, north+d+4*sw, x, south-d-4*sw, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
		color, size, opacity))

}

// Label writes a process label
func Label(x, y int, txt string) {
	canvas.Text(x-d-4*sw, y+2*sw, txt, fmt.Sprintf("text-anchor:%s;font-size:%dpx;fill-opacity:%.2f;font-family:Verdana",
		"end", 18, opacity))
}

func Start(width, height int) {
	buffer = &bytes.Buffer{}
	canvas = svg.New(buffer)
	canvas.Start(width, height)
	canvas.Rect(0, 0, width, height, "fill:rgb(250,250,250)")
	canvas.Grid(0, 0, width, height, 25, "stroke:lightgray")
}

func End() *bytes.Buffer {
	canvas.End()
	return buffer
}
