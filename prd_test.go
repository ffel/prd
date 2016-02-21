package prd

import (
	"fmt"
	"os"
)

const (
	procesA Proces = iota
	procesB
	procesC
)

const (
	channelA Channel = iota
	channelB
)

// go test 2> out.svg
func Example() {
	PrdStart(13, 3)

	LabelChannel(channelA, "a")
	LabelChannel(channelB, "b")

	At(0, procesA).Starts("A")

	At(1, procesA).Creates(procesB, "B")
	At(2, procesA).Creates(procesC, "C")

	At(3, procesA).WantsToReceiveOn(channelA)
	At(5, procesC).WantsToSendOn(channelA, "data")

	At(6, procesA).WantsToReceiveOn(channelA)

	At(8, procesB).WantsToSendOn(channelA, "data")

	At(6, procesC).WantsToSendOn(channelB, "data")
	At(10, procesB).WantsToReceiveOn(channelB)

	At(14, procesA).Terminates()
	At(14, procesB).Terminates()

	PrdEnd()

	fmt.Fprintln(os.Stdout, Log.String())
	fmt.Fprintln(os.Stderr, SVG.String())

	// output:
	// ** at 0, proces "A" starts
	// at 1, proces "A" creates proces "B"
	// at 2, proces "A" creates proces "C"
	// at 3, proces "A" wants to receive on channel "a"
	// at 5, proces "C" wants to send "data" on channel "a"
	// - received by proces "A"
	// at 6, proces "A" wants to receive on channel "a"
	// at 8, proces "B" wants to send "data" on channel "a"
	// - received by proces "A"
	// at 6, proces "C" wants to send "data" on channel "b"
	// at 10, proces "B" wants to receive on channel "b"
	// - sent by proces "C"
	// at 14, proces "A" terminates
	// at 14, proces "B" terminates
}
