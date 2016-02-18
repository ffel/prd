// prd - process diagrams, simple diagrams for concurrent processes
package prd

import (
	"bytes"
	"fmt"

	"github.com/ffel/prd/prdsymb"
)

var (
	offsetX int = 230
	offsetY int = 50
	deltaX  int = 60
	deltaY  int = 60
)

type Proces int

type Channel int

type processtate int

const color = "red"

const (
	active processtate = iota
	waitingForSend
	waitingForReceive
)

type state struct {
	since   int
	pstate  processtate
	channel Channel // in case of waiting, which channel
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
	states[v.proc] = state{since: v.time, pstate: active}
	prdsymb.Label(x(v.time), y(v.proc), label)
}

func (v Verb) Creates(proc Proces, label string) {
	fmt.Printf(" creates proces %q with label %q\n", proc, label)
	states[proc] = state{since: v.time, pstate: active}
	prdsymb.Label(x(v.time), y(proc), label)
	prdsymb.Create(x(v.time), y(v.proc), y(proc))
}

func (v Verb) WantsToReceiveOn(c Channel) OnOption {
	fmt.Printf(" wants to receive on channel %q\n", c)
	fmt.Printf("-R- since %d, state %q, channel: %q, now %d\n", states[v.proc].since, states[v.proc].pstate, states[v.proc].channel, v.time)

	// draw line
	if states[v.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[v.proc].since), x(v.time), y(v.proc))
	}
	// draw receive symbol
	prdsymb.Receive(prdsymb.Wait, x(v.time), y(v.proc), color)

	// if someone is able to send on c, handle that

	states[v.proc] = state{since: v.time, pstate: waitingForReceive, channel: c}

	// we hebben hier wel een option probleem: we kunnen niet *vooruit* kijken naar een evt
	// handler.  Of omdraaien (eerst optie), of een terminating call

	return OnOption{}
}

// finds a process that currently wants to receive on channel c
// returns false in case there is no such Proces
func findPresentReceiver(c Channel) (Proces, bool) {
	for v, k := range states {
		if k.channel == c && k.pstate == waitingForReceive {
			return v, true
		}
	}

	return 0, false
}

func (v Verb) WantsToSendOn(c Channel, data string) OnOption {
	fmt.Printf(" wants to send %q on channel %q\n", data, c)
	fmt.Printf("-S- since %d, state %q, channel: %q, now %d\n", states[v.proc].since, states[v.proc].pstate, states[v.proc].channel, v.time)

	if states[v.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[v.proc].since), x(v.time), y(v.proc))
	}
	prdsymb.Send(prdsymb.Wait, x(v.time), y(v.proc), color)

	// if someone is able to receive on c, handle that
	if rproc, ok := findPresentReceiver(c); ok {
		fmt.Printf("- received by process %q\n", rproc)
		prdsymb.Receive(prdsymb.Postponed, x(v.time), y(rproc), color)
		prdsymb.Send(prdsymb.Postponed, x(v.time), y(v.proc), color)
		prdsymb.Channel(x(v.time), y(v.proc), y(rproc), color)
		prdsymb.Process(prdsymb.Asleep, x(states[rproc].since), x(v.time), y(rproc))

		// change states
	}

	states[v.proc] = state{since: v.time, pstate: waitingForSend, channel: c}
	return OnOption{}
}

func (v Verb) Terminates() {
	fmt.Printf(" terminates\n")
}

type OnOption struct{}

func (o OnOption) HandledBy(proc Proces) {
	fmt.Printf("-  handled by %q\n", proc)
}
