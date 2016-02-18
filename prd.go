// prd - process diagrams, simple diagrams for concurrent processes
package prd

import (
	"bytes"
	"fmt"

	"github.com/ffel/prd/prdsymb"
)

var (
	offsetX int = 100
	offsetY int = 25
	deltaX  int = 30
	deltaY  int = 40
)

type Proces int

type Channel int

type processtate int

const (
	active processtate = iota
	waitingForSend
	waitingForReceive
	terminated
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

func At(time int, proc Proces) ProcesInfo {
	fmt.Printf("at %d, process %q", time, proc)
	return ProcesInfo{time, proc}
}

type ProcesInfo struct {
	time int
	proc Proces
}

func (info ProcesInfo) Starts(label string) {
	fmt.Printf(" starts with label %q\n", label)
	states[info.proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(info.proc), label)
}

func (info ProcesInfo) Creates(proc Proces, label string) {
	fmt.Printf(" creates proces %q with label %q\n", proc, label)
	states[proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(proc), label)
	prdsymb.Create(x(info.time), y(info.proc), y(proc))
}

// WantsToReceive marks proces info.proc as to want receive on channel c
// If another proces is to send on c, the receive will actually happen
func (info ProcesInfo) WantsToReceiveOn(c Channel) OnOption {
	fmt.Printf(" wants to receive on channel %q\n", c)
	// fmt.Printf("-R- since %d, state %q, channel: %q, now %d\n", states[info.proc].since, states[info.proc].pstate, states[info.proc].channel, info.time)

	// draw line
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw receive symbol
	prdsymb.Receive(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if sproc, ok := findPresentSender(c); ok {
		fmt.Printf("- sent by proces %q\n", sproc)

		prdsymb.Send(prdsymb.Postponed, x(info.time), y(sproc), channelColor(c))
		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(sproc), y(info.proc), channelColor(c))
		prdsymb.Process(prdsymb.Asleep, x(states[sproc].since), x(info.time), y(sproc))

		states[info.proc] = state{since: info.time, pstate: active}
		states[sproc] = state{since: info.time, pstate: active}
	} else {
		states[info.proc] = state{since: info.time, pstate: waitingForReceive, channel: c}
	}

	return OnOption{}
}

// WantsToSendOn marks proces info.proc as to want send on channel c.
// In another proces is to receive on c, the send will actually happen
func (info ProcesInfo) WantsToSendOn(c Channel, data string) OnOption {
	fmt.Printf(" wants to send %q on channel %q\n", data, c)
	// fmt.Printf("-S- since %d, state %q, channel: %q, now %d\n", states[info.proc].since, states[info.proc].pstate, states[info.proc].channel, info.time)

	// draw proces lone for info.proc
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw send symbol
	prdsymb.Send(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if rproc, ok := findPresentReceiver(c); ok {
		fmt.Printf("- received by proces %q\n", rproc)

		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(rproc), channelColor(c))
		prdsymb.Send(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(info.proc), y(rproc), channelColor(c))
		prdsymb.Process(prdsymb.Asleep, x(states[rproc].since), x(info.time), y(rproc))

		states[info.proc] = state{since: info.time, pstate: active}
		states[rproc] = state{since: info.time, pstate: active}
	} else {
		states[info.proc] = state{since: info.time, pstate: waitingForSend, channel: c}
	}

	return OnOption{}
}

func (info ProcesInfo) Terminates() {
	fmt.Printf(" terminates\n")

	var mode prdsymb.Mode
	mode = prdsymb.Active

	if states[info.proc].pstate != active {
		mode = prdsymb.Asleep
	}

	prdsymb.Process(mode, x(states[info.proc].since), x(info.time), y(info.proc))
	states[info.proc] = state{since: info.time, pstate: terminated}
}

type OnOption struct{}

func (o OnOption) HandledBy(proc Proces) {
	fmt.Printf("-  handled by %q\n", proc)
}

func channelColor(c Channel) string {
	switch int(c) % 3 {
	case 0:
		return "red"
	case 1:
		return "blue"
	case 2:
		return "green"
	}

	return "black"
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

// finds a process that currently wants to send on channel c
// returns false in case there is no such Proces
func findPresentSender(c Channel) (Proces, bool) {
	for v, k := range states {
		if k.channel == c && k.pstate == waitingForSend {
			return v, true
		}
	}

	return 0, false
}
