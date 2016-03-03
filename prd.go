// prd - process diagrams, simple diagrams for concurrent processes
package prd

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/ffel/prd/prdsymb"
)

var (
	offsetX     int = 150
	offsetY     int = 25
	deltaX      int = 30
	deltaY      int = 60
	deltaSelect int = 4
)

//go:generate stringer -type=processtate

const (
	active processtate = iota
	waiting
	forSend
	forReceive
	terminated
)

// state is the current state a proces is in
type state struct {
	since     int
	pstate    processtate // active, waiting or terminated
	waitstate []wstate    // if waiting
}

type wstate struct {
	state   processtate // forSend, forReceive
	channel Channel
	data    string // in case of send
}

var states map[Proces]state
var procLabels map[Proces]string
var chanLabels map[Channel]string
var timeLabels map[int]bool
var labely int

var Log *bytes.Buffer
var SVG *bytes.Buffer

type Proces int

type Channel int

type processtate int

// PrdStarts initialises drawing the proces diagram, should be first
func PrdStart(totalTime, nrProceses int) {
	states = make(map[Proces]state)
	procLabels = make(map[Proces]string)
	chanLabels = make(map[Channel]string)
	timeLabels = make(map[int]bool)
	labely = nrProceses
	prdsymb.Start(x(totalTime+1), y(Proces(labely+1)))
	Log = new(bytes.Buffer) // hah! a meaningful use of new()
}

// PrdEnd completes drawing the proces diagram, should be last
func PrdEnd() {
	for k, _ := range timeLabels {
		prdsymb.Label(x(k)+2*deltaX/3, y(Proces(labely)), strconv.Itoa(k))
	}

	SVG = prdsymb.End()
}

// LabelChannel is used to give channels a string representation
func LabelChannel(c Channel, lab string) {
	chanLabels[c] = lab
}

// At is the first part of a combined diagram instruction
func At(time int, proc Proces) ProcesInfo {
	timeLabels[time] = true // record time

	if lab, ok := procLabels[proc]; ok {
		fmt.Fprintf(Log, "at %d, proces %q", time, lab)
	} else {
		fmt.Fprintf(Log, "*")
	}

	return ProcesInfo{time: time, proc: proc}
}

// ProcesInfo contains info collected by subsequent diagram instruction
// commands
type ProcesInfo struct {
	time       int     // time set by At
	proc       Proces  // proces set by At
	servedBy   bool    // true if AsServedByProces has been explicitly set
	servedProc Proces  // proces set by AsServedByProces
	servedChan Channel // channel set by AsServedByProces
}

// Starts initiates a proces that is already active
func (info ProcesInfo) Starts(label string) {
	fmt.Fprintf(Log, "* at %d, proces %q starts\n", info.time, label)
	procLabels[info.proc] = label
	states[info.proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(info.proc), label)
}

// Creates indicates a new proces
func (info ProcesInfo) Creates(proc Proces, label string) {
	fmt.Fprintf(Log, " creates proces %q\n", label)
	procLabels[proc] = label
	states[proc] = state{since: info.time, pstate: active}
	prdsymb.Label(x(info.time), y(proc), label)
	prdsymb.Create(x(info.time), y(info.proc), y(proc))
}

// AsSentByProces can specify proces that will serve proces in
// At.  Used in case there are several process that can serve.
func (info ProcesInfo) AsServedBy(proc Proces, channel Channel) ProcesInfo {
	return ProcesInfo{
		time:       info.time,
		proc:       info.proc,
		servedBy:   true,
		servedProc: proc,
		servedChan: channel,
	}
}

// WantsToReceive marks proces info.proc as to want receive on channel c
// If another proces is to send on c, the receive will actually happen
func (info ProcesInfo) WantsToReceiveOn(c Channel) AndInfo {
	fmt.Fprintf(Log, " wants to receive on channel %q\n", chanLabels[c])

	// draw line
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw receive symbol
	prdsymb.Receive(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if info.servedBy {
		connectProces(info, info.servedChan, info.servedProc, forSend, "")
		// set state to active
	} else if sproc, ok := findWaiting(forSend, c, info.time); ok {
		connectProces(info, c, sproc, forSend, "")
	} else {
		states[info.proc] = state{
			since:     info.time,
			pstate:    waiting,
			waitstate: []wstate{{state: forReceive, channel: c}},
		}
	}

	return AndInfo{info, 0}
}

// WantsToSendOn marks proces info.proc as to want send on channel c.
// In another proces is to receive on c, the send will actually happen
func (info ProcesInfo) WantsToSendOn(c Channel, data string) AndInfo {
	fmt.Fprintf(Log, " wants to send %q on channel %q\n", data, chanLabels[c])

	// draw proces line for info.proc
	if states[info.proc].pstate == active {
		prdsymb.Process(prdsymb.Active, x(states[info.proc].since), x(info.time), y(info.proc))
	}
	// draw send symbol
	prdsymb.Send(prdsymb.Wait, x(info.time), y(info.proc), channelColor(c))

	if info.servedBy {
		connectProces(info, info.servedChan, info.servedProc, forReceive, data)
	} else if rproc, ok := findWaiting(forReceive, c, info.time); ok {
		connectProces(info, c, rproc, forReceive, data)
	} else {
		states[info.proc] = state{
			since:     info.time,
			pstate:    waiting,
			waitstate: []wstate{{state: forSend, channel: c, data: data}},
		}
	}

	return AndInfo{info, 0}
}

// AndInfo is used for combined receive or send instructions
type AndInfo struct {
	ProcesInfo
	nr int
}

// AndToReceiveOn adds another receiver at a point that already
// has receivers or senders
func (and AndInfo) AndToReceiveOn(c Channel) AndInfo {
	and.nr++

	fmt.Fprintf(Log, "- and proces %q wants to receive on channel %q\n",
		procLabels[and.proc], chanLabels[c])

	prdsymb.Receive(prdsymb.Wait,
		x(and.time)-and.nr*deltaSelect,
		y(and.proc)-and.nr*deltaSelect, channelColor(c))

	if sproc, ok := findWaiting(forSend, c, and.time); ok && states[and.proc].pstate == waiting {
		connectProces(and.ProcesInfo, c, sproc, forSend, "")
	}

	addWaitState(and.proc, forReceive, c, "")

	return and
}

func (and AndInfo) AndToSendOn(c Channel, data string) AndInfo {
	and.nr++

	fmt.Fprintf(Log, "- and proces %q wants to send %q on channel %q\n",
		procLabels[and.proc], data, chanLabels[c])

	prdsymb.Send(prdsymb.Wait,
		x(and.time)-and.nr*deltaSelect,
		y(and.proc)-and.nr*deltaSelect, channelColor(c))

	if rproc, ok := findWaiting(forReceive, c, and.time); ok && states[and.proc].pstate == waiting {
		connectProces(and.ProcesInfo, c, rproc, forReceive, data)
	}

	addWaitState(and.proc, forSend, c, data)

	return and
}

// DoesNotWait adds does not wait symbol
func (and AndInfo) AndDoesNotWait() {
	fmt.Fprintf(Log, " does not want to wait\n")

	prdsymb.Else(prdsymb.Wait,
		x(and.time), y(and.proc))

	if states[and.proc].pstate == waiting {
		prdsymb.Else(prdsymb.Immediately,
			x(and.time), y(and.proc))

		states[and.proc] = state{
			since:     and.time,
			pstate:    active,
			waitstate: []wstate{},
		}
	} else {
		prdsymb.Else(prdsymb.Wait,
			x(and.time), y(and.proc))
	}
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

// connectProces finds a proces that is willing to serve the present
// proces that wants to send or receive
func connectProces(info ProcesInfo, c Channel, rsProces Proces, pstate processtate, data string) {
	verb := "sent"

	if pstate == forReceive {
		verb = "received"
	}

	fmt.Fprintf(Log, "- %s by proces %q\n", verb, procLabels[rsProces])

	if pstate == forReceive {
		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(rsProces), channelColor(c))
		prdsymb.Send(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(rsProces), y(info.proc), channelColor(c), data)
	} else if pstate == forSend {
		prdsymb.Send(prdsymb.Postponed, x(info.time), y(rsProces), channelColor(c))
		prdsymb.Receive(prdsymb.Postponed, x(info.time), y(info.proc), channelColor(c))
		prdsymb.Channel(x(info.time), y(info.proc), y(rsProces), channelColor(c),
			states[rsProces].waitstate[len(states[rsProces].waitstate)-1].data)
	}

	prdsymb.Process(prdsymb.Asleep, x(states[rsProces].since), x(info.time), y(rsProces))

	// kan ik hier bij het verstuurde bericht?

	states[info.proc] = state{since: info.time, pstate: active}
	states[rsProces] = state{since: info.time, pstate: active}
}

// checkWaitingProces asserts a proces for waiting to send or receive
func checkWaitingProces(proc Proces, ch Channel, state processtate) bool {
	if states[proc].pstate != waiting {
		return false
	}
	for _, w := range states[proc].waitstate {
		if w.channel == ch && w.state == state {
			return true
		}
	}
	return false
}

// addWaitState expands a proces state with another receiver or sender
func addWaitState(proc Proces, nstate processtate, channel Channel, data string) {

	wstates := append(states[proc].waitstate,
		wstate{state: nstate, channel: channel, data: data})

	update := state{
		since:     states[proc].since,
		pstate:    states[proc].pstate,
		waitstate: wstates,
	}

	states[proc] = update
}

// finds a process that currently wants to receive or send on channel c
// returns false in case there is no such Proces
func findWaiting(pstate processtate, c Channel, now int) (Proces, bool) {
	for k, v := range states {
		if now >= v.since && v.pstate == waiting {
			for _, w := range v.waitstate {
				if w.state == pstate && w.channel == c {
					return k, true
				}
			}
		}
	}

	return 0, false
}

// x calculates x pos based upon time value
func x(val int) int {
	return offsetX + (val-1)*deltaX
}

// y calculates y pos based upon proces number
func y(val Proces) int {
	return offsetY + int(val)*deltaY
}

func channelColor(c Channel) string {
	switch int(c) % 4 {
	case 0:
		return "red"
	case 1:
		return "blue"
	case 2:
		return "green"
	case 3:
		return "darkmagenta"
	}

	return "black"
}
