// PRoces Diagram svg symbols, lines etc
package prdsymb

import (
	"fmt"
	"os"
)

func Example() {
	Start(400, 400)

	Send(Wait, 50, 50, "red")
	Send(Postponed, 100, 50, "red")
	Send(Immediately, 150, 50, "red")

	Receive(Wait, 50, 100, "blue")
	Receive(Postponed, 100, 100, "blue")
	Receive(Immediately, 150, 100, "blue")

	Else(Wait, 50, 150)
	Else(Postponed, 100, 150)
	Else(Immediately, 150, 150)

	Else(Wait, 50, 200)
	Send(Wait, 50-4, 200-4, "red")
	Receive(Wait, 50-8, 200-8, "blue")

	Else(Wait, 100, 200)
	Send(Wait, 100-4, 200-4, "red")
	Receive(Wait, 100-8, 200-8, "blue")
	Send(Postponed, 100, 200, "red")

	Else(Wait, 150, 200)
	Send(Wait, 150-4, 200-4, "red")
	Receive(Wait, 150-8, 200-8, "blue")
	Receive(Postponed, 150, 200, "blue")

	Else(Wait, 200, 200)
	Send(Wait, 200-4, 200-4, "red")
	Receive(Wait, 200-8, 200-8, "blue")
	Else(Postponed, 200, 200)

	buff := End()

	fmt.Fprintln(os.Stdout, buff.String())

	// output:
	// <?xml version="1.0"?>
	// <!-- Generated by SVGo -->
	// <svg width="400" height="400"
	//      xmlns="http://www.w3.org/2000/svg"
	//      xmlns:xlink="http://www.w3.org/1999/xlink">
	// <polygon points="39,41 39,61 59,61" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polygon points="91,39 111,39 111,59" style="stroke:none;stroke-width:2;fill:red;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <polygon points="139,41 139,61 159,61" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polygon points="141,39 161,39 161,59" style="stroke:none;stroke-width:2;fill:red;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <polyline points="39,111 39,89 61,89" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <rect x="93" y="93" width="18" height="18" style="stroke:none;stroke-width:2;fill:blue;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <polyline points="139,111 139,89 161,89" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <rect x="143" y="93" width="18" height="18" style="stroke:none;stroke-width:2;fill:blue;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <line x1="38" y1="162" x2="62" y2="138" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <circle cx="100" cy="150" r="6" style="fill-opacity:0.30;"/>
	// <line x1="138" y1="162" x2="162" y2="138" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <circle cx="150" cy="150" r="6" style="fill-opacity:0.30;"/>
	// <line x1="38" y1="212" x2="62" y2="188" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <polygon points="35,187 35,207 55,207" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polyline points="31,203 31,181 53,181" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <line x1="88" y1="212" x2="112" y2="188" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <polygon points="85,187 85,207 105,207" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polyline points="81,203 81,181 103,181" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <polygon points="91,189 111,189 111,209" style="stroke:none;stroke-width:2;fill:red;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <line x1="138" y1="212" x2="162" y2="188" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <polygon points="135,187 135,207 155,207" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polyline points="131,203 131,181 153,181" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <rect x="143" y="193" width="18" height="18" style="stroke:none;stroke-width:2;fill:blue;stroke-opacity:0.30;fill-opacity:0.30"/>
	// <line x1="188" y1="212" x2="212" y2="188" style="stroke:black;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <polygon points="185,187 185,207 205,207" style="stroke:red;stroke-width:2;fill:none;stroke-opacity:0.30"/>
	// <polyline points="181,203 181,181 203,181" style="stroke:blue;stroke-width:2;fill:none;stroke-opacity:0.30;stroke-linecap:round"/>
	// <circle cx="200" cy="200" r="6" style="fill-opacity:0.30;"/>
	// </svg>

}