// prd - process diagrams, simple diagrams for concurrent processes
package prd

import (
	"bytes"
	"fmt"

	"github.com/ffel/prd/prdsymb"
)

var (
	offsetX int = 250
	offsetY int = 50
	deltaX  int = 50
	deltaY  int = 50
)

type Proces int

type Channel int

type processtate int

const (
	active processtate = iota
	waitingForSend
	waitingForReceive
)

type state struct {
	since  int
	pstate processtate
}

var states map[Proces]state

func x(val int) int {
	return offsetX + (val-1)*deltaX
}

func y(val Proces) int {
	return offsetY + int(val)*deltaY
}

func PrdStart(width, height int) {
	states = make(map[Proces]state)
	prdsymb.Start(width, height)
}

func PrdEnd() *bytes.Buffer {
	return prdsymb.End()
}

func At(time int, proc Proces) Verb {
	fmt.Printf("at %d, process %q", time, proc)
	return Verb{time, proc}
}

type Verb struct {
	time int
	proc Proces
}

func (v Verb) Starts(label string) {
	fmt.Printf(" starts with label %q\n", label)
	states[v.proc] = state{v.time, active}
	prdsymb.Label(x(v.time), y(v.proc), label)
}

func (v Verb) Creates(proc Proces, label string) {
	fmt.Printf(" creates proces %q with label %q\n", proc, label)
	states[v.proc] = state{v.time, active}
	prdsymb.Label(x(v.time), y(proc), label)
	prdsymb.Create(x(v.time), y(v.proc), y(proc))
}

func (v Verb) WantsToReceive() On {
	fmt.Printf(" wants to receive")
	return On{}
}

func (v Verb) WantsToSend(data string) On {
	fmt.Printf(" wants to send %q", data)
	return On{}
}

func (v Verb) Terminates() {
	fmt.Printf(" terminates\n")
}

type On struct{}

func (on On) OnChannel(c Channel) OnOption {
	fmt.Printf(" on channel %q\n", c)

	return OnOption{}
}

type OnOption struct{}

func (o OnOption) HandledBy(proc Proces) {
	fmt.Printf("-  handled by %q\n", proc)
}
