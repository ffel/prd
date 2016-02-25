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

func Example_select_1() {
	PrdStart(15, 2)

	LabelChannel(channelA, "a")
	LabelChannel(channelB, "b")

	At(0, procesA).Starts("A")
	At(0, procesB).Starts("B")

	At(1, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)
	At(3, procesB).WantsToSendOn(channelA, "data")
	At(4, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)
	At(6, procesB).WantsToSendOn(channelB, "data")

	PrdEnd()

	// fmt.Fprintln(os.Stderr, SVG.String())
	fmt.Fprintln(os.Stdout, Log.String())

	// output:
	// ** at 0, proces "A" starts
	// ** at 0, proces "B" starts
	// at 1, proces "A" wants to receive on channel "a"
	// - and proces "A" wants to receive on channel "b"
	// at 3, proces "B" wants to send "data" on channel "a"
	// - received by proces "A"
	// at 4, proces "A" wants to receive on channel "a"
	// - and proces "A" wants to receive on channel "b"
	// at 6, proces "B" wants to send "data" on channel "b"
	// - received by proces "A"
}

func Example_select_2() {
	PrdStart(15, 3)

	LabelChannel(channelA, "a")
	LabelChannel(channelB, "b")

	At(0, procesA).Starts("A")
	At(0, procesB).Starts("B")
	At(0, procesC).Starts("C")

	At(1, procesB).WantsToSendOn(channelA, "data")
	At(1, procesC).WantsToSendOn(channelB, "data")

	At(3, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)

	At(5, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)

	PrdEnd()

	// fmt.Fprintln(os.Stderr, SVG.String())
	fmt.Fprintln(os.Stdout, Log.String())

	// output:
	// ** at 0, proces "A" starts
	// ** at 0, proces "B" starts
	// ** at 0, proces "C" starts
	// at 1, proces "B" wants to send "data" on channel "a"
	// at 1, proces "C" wants to send "data" on channel "b"
	// at 3, proces "A" wants to receive on channel "a"
	// - sent by proces "B"
	// - and proces "A" wants to receive on channel "b"
	// at 5, proces "A" wants to receive on channel "a"
	// - and proces "A" wants to receive on channel "b"
	// - sent by proces "C"
}

func Example_select_3() {
	PrdStart(15, 3)

	LabelChannel(channelA, "a")
	LabelChannel(channelB, "b")

	At(0, procesA).Starts("A")
	At(0, procesB).Starts("B")
	At(0, procesC).Starts("C")

	At(1, procesB).WantsToSendOn(channelA, "data")
	At(1, procesC).WantsToSendOn(channelB, "data")

	// by default, procesA will receive from channelA as this is checked first
	// to override, we'll need AsServedByProces
	At(3, procesA).AsServedBy(procesC, channelB).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)
	At(4, procesC).WantsToSendOn(channelB, "data")

	At(5, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)

	At(7, procesA).WantsToReceiveOn(channelA).AndToReceiveOn(channelB)

	PrdEnd()

	// fmt.Fprintln(os.Stderr, SVG.String())
	fmt.Fprintln(os.Stdout, Log.String())

	// output:
	// ** at 0, proces "A" starts
	// ** at 0, proces "B" starts
	// ** at 0, proces "C" starts
	// at 1, proces "B" wants to send "data" on channel "a"
	// at 1, proces "C" wants to send "data" on channel "b"
	// at 3, proces "A" wants to receive on channel "a"
	// - sent by proces "C"
	// - and proces "A" wants to receive on channel "b"
	// at 4, proces "C" wants to send "data" on channel "b"
	// at 5, proces "A" wants to receive on channel "a"
	// - sent by proces "B"
	// - and proces "A" wants to receive on channel "b"
	// at 7, proces "A" wants to receive on channel "a"
	// - and proces "A" wants to receive on channel "b"
	// - sent by proces "C"
}

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
	// fmt.Fprintln(os.Stderr, SVG.String())

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
