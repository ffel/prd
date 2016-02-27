package prdsymb

import (
	"bytes"
	"fmt"

	"github.com/ajstarks/svgo"
)

var buffer *bytes.Buffer

var canvas *svg.SVG

type Mode int

const (
	Wait Mode = iota // send and receive
	Postponed
	Immediately
	Active // processes
	Asleep
)

const (
	d        = 10  // half symbol size
	size     = 2   // line width
	sw       = 1   // off set
	opacity  = 0.3 // opacity
	fontsize = 11  // font size
)

// draw a send symbol (is a little bit like an "s")
func Send(m Mode, x, y int, color string) {
	if m == Wait || m == Immediately {
		canvas.Polygon([]int{x - d - sw, x - d - sw, x + d - sw}, []int{y - d + sw, y + d + sw, y + d + sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f",
				color, size, opacity))
	}
	if m == Postponed || m == Immediately {
		canvas.Polygon([]int{x - d + sw, x + d + sw, x + d + sw}, []int{y - d - sw, y - d - sw, y + d - sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:%s;stroke-opacity:%.2f;fill-opacity:%.2f",
				"none", size, color, opacity, opacity))
	}
}

// Receice draws a receive message symbol, shaped like an "r"
func Receive(m Mode, x, y int, color string) {
	if m == Wait || m == Immediately {
		canvas.Polyline([]int{x - d - sw, x - d - sw, x + d + sw}, []int{y + d + sw, y - d - sw, y - d - sw},
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
				color, size, opacity))
	}
	if m == Postponed || m == Immediately {
		canvas.Square(x-d+3*sw, y-d+3*sw, 2*d-2*sw,
			fmt.Sprintf("stroke:%s;stroke-width:%d;fill:%s;stroke-opacity:%.2f;fill-opacity:%.2f",
				"none", size, color, opacity, opacity))
	}
}

// Else draws the else symbol, a double line in line with channel
// communication
func Else(m Mode, x, y int) {
	if m == Wait || m == Immediately {
		canvas.Line(x-2*sw, y+7*d/5, x-2*sw, y-7*d/5, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
			"black", size, opacity))
		canvas.Line(x+2*sw, y+7*d/5, x+2*sw, y-7*d/5, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
			"black", size, opacity))
	}
	if m == Postponed || m == Immediately {
		canvas.Circle(x, y, 6*sw, fmt.Sprintf("fill-opacity:%.2f;", opacity))
	}
}

// Process draws a horizontal process line
func Process(m Mode, x1, x2, y int) {
	if m == Active {
		canvas.Line(x1+d+4*sw, y, x2-d-4*sw, y, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round",
			"black", size, opacity))
	} else if m == Asleep {
		canvas.Line(x1+d+4*sw, y, x2-d-4*sw, y, fmt.Sprintf("stroke:%s;stroke-width:%d;fill:none;stroke-opacity:%.2f;stroke-linecap:round;stroke-dasharray:1, 5",
			"black", size, opacity))
	}
}

// Create draws a proces creation line from y1 to y2
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
		"end", fontsize, opacity))
}

func Start(width, height int) {
	buffer = &bytes.Buffer{}
	canvas = svg.New(buffer)
	canvas.Start(width, height)
	// canvas.Rect(0, 0, width, height, "fill:rgb(253,253,253)")
	// canvas.Grid(0, 0, width, height, 25, "stroke:lightgray")
}

func End() *bytes.Buffer {
	canvas.End()
	return buffer
}
