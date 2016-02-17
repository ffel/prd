// prd - process diagrams, simple diagrams for concurrent processes
package prd

import "fmt"

//go:generate stringer -type=Proces,Channel

type Proces int

const (
	goingForAWalk Proces = iota
	AliceGettingReady
	BobGettingReady
)

type Channel int

const (
	a Channel = iota
	b
	c
)

func At(time int, proc Proces) Verb {
	fmt.Printf("at %d, process %q", time, proc)
	return Verb{}
}

type Verb struct{}

func (v Verb) Starts(label string) {
	fmt.Printf(" starts with label %q\n", label)
}

func (v Verb) Creates(proc Proces, label string) {
	fmt.Printf(" creates proces %q with label %q\n", proc, label)
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
