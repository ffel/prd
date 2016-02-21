// prd - process diagrams, simple diagrams for concurrent processes
package prd

import (
	"bytes"
	"fmt"

	"github.com/ffel/prd/prdsymb"
)

var (
	offsetX int = 150
	offsetY int = 25
	deltaX  int = 30
	deltaY  int = 40
)

var Log *bytes.Buffer
var SVG *bytes.Buffer

type Proces int

type Channel int

type processtate int

//go:generate stringer -type=processtate

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
var procLabels map[Proces]string
var chanLabels map[Channel]string

func x(val int) int {
	return offsetX + (val-1)*deltaX
}

func y(val Proces) int {
	return offsetY + int(val)*deltaY
}

func PrdStart(totalTime, nrProceses int) {
	states = make(map[Proces]state)
	procLabels = make(map[Proces]string)
	chanLabels = make(map[Channel]string)
	prdsymb.Start(x(totalTime+1), y(Proces(nrProceses)+1))
	Log = new(bytes.Buffer) // hah! a meaningful use of new()
}

// func PrdEnd() *bytes.Buffer {
// 	return prdsymb.End()
// }
func PrdEnd() {
	SVG = prdsymb.End()
}

func LabelChannel(c Channel, lab string) {
	chanLabels[c] = lab
}

func At(time int, proc Proces) ProcesInfo {
	if lab, ok := procLabels[proc]; ok {
		fmt.Fprintf(Log, "at %d, proces %q", time, lab)
	} else {
		fmt.Fprintf(Log, "*")
	}

	return ProcesInfo{time: time, proc: proc}
}

type ProcesInfo struct {
	time       int
	proc       Proces
	servedby   bool   // true if servedproc has been explicitly set
	servedproc Proces // in case there are multiple helpers
}

func (info ProcesInfo) Starts(label string) {
	fmt.Fprintf(Log, "* at %d, proces %q starts\n", info.time, label)
	procLabels[info.proc] = label
	states[info.proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(info.proc), label)
}

func (info ProcesInfo) Creates(proc Proces, label string) {
	fmt.Fprintf(Log, " creates proces %q\n", label)
	procLabels[proc] = label
	states[proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(proc), label)
	prdsymb.Create(x(info.time), y(info.proc), y(proc))
}

// WantsToReceive marks proces info.proc as to want receive on channel c
// If another proces is to send on c, the receive will actually happen
func (info ProcesInfo) WantsToReceiveOn(c Channel) {
	fmt.Fprintf(Log, " wants to receive on channel %q\n", chanLabels[c])
	// fmt.Fprintf(Log,"-R- since %d, state %q, channel: %q, now %d\n", states[info.proc].since, states[info.proc].pstate, states[info.proc].channel, info.time)

	// draw line
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw receive symbol
	prdsymb.Receive(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if sproc, ok := findPresentSender(c); ok {
		// check if AsServedByProces is added
		if info.servedby {
			sproc = info.servedproc

			if states[sproc].pstate != waitingForSend {
				fmt.Fprintf(Log, "WARNING: proces %s not waiting to send (%q)\n",
					procLabels[sproc], states[sproc].pstate)
			}
		}

		fmt.Fprintf(Log, "- sent by proces %q\n", procLabels[sproc])

		prdsymb.Send(prdsymb.Postponed, x(info.time), y(sproc), channelColor(c))
		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(sproc), y(info.proc), channelColor(c))
		prdsymb.Process(prdsymb.Asleep, x(states[sproc].since), x(info.time), y(sproc))

		states[info.proc] = state{since: info.time, pstate: active}
		states[sproc] = state{since: info.time, pstate: active}
	} else {
		states[info.proc] = state{since: info.time, pstate: waitingForReceive, channel: c}
	}
}

// WantsToSendOn marks proces info.proc as to want send on channel c.
// In another proces is to receive on c, the send will actually happen
func (info ProcesInfo) WantsToSendOn(c Channel, data string) {
	fmt.Fprintf(Log, " wants to send %q on channel %q\n", data, chanLabels[c])
	// fmt.Fprintf(Log,"-S- since %d, state %q, channel: %q, now %d\n", states[info.proc].since, states[info.proc].pstate, states[info.proc].channel, info.time)

	// draw proces line for info.proc
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw send symbol
	prdsymb.Send(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if rproc, ok := findPresentReceiver(c); ok {
		// check if AsServedByProces is added
		if info.servedby {
			rproc = info.servedproc

			if states[rproc].pstate != waitingForReceive {
				fmt.Fprintf(Log, "WARNING: proces %s not waiting to reveive (%q)\n",
					procLabels[rproc], states[rproc].pstate)
			}
		}

		fmt.Fprintf(Log, "- received by proces %q\n", procLabels[rproc])

		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(rproc), channelColor(c))
		prdsymb.Send(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(info.proc), y(rproc), channelColor(c))
		prdsymb.Process(prdsymb.Asleep, x(states[rproc].since), x(info.time), y(rproc))

		states[info.proc] = state{since: info.time, pstate: active}
		states[rproc] = state{since: info.time, pstate: active}
	} else {
		states[info.proc] = state{since: info.time, pstate: waitingForSend, channel: c}
	}
}

// AsServedByProces can specify proces that will serve proces in
// At.
func (info ProcesInfo) AsServedByProces(proc Proces) ProcesInfo {
	return ProcesInfo{time: info.time, proc: info.proc, servedby: true, servedproc: proc}
}

func (info ProcesInfo) Terminates() {
	fmt.Fprintf(Log, " terminates\n")

	var mode prdsymb.Mode
	mode = prdsymb.Active

	if states[info.proc].pstate != active {
		mode = prdsymb.Asleep
	}

	prdsymb.Process(mode, x(states[info.proc].since), x(info.time), y(info.proc))
	states[info.proc] = state{since: info.time, pstate: terminated}
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
