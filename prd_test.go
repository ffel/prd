package prd

import (
	"fmt"
	"os"
)

// go test 2> out.svg
func Example() {
	PrdStart(800, 500)

	At(0, procesA).Starts("proces A")

	At(1, procesA).Creates(procesB, "proces B")
	At(2, procesA).Creates(procesC, "proces C")

	At(3, procesA).WantsToReceiveOn(a)
	At(5, procesC).WantsToSendOn(a, "data")

	At(6, procesA).WantsToReceiveOn(a)

	At(8, procesB).WantsToSendOn(a, "data")

	At(6, procesC).WantsToSendOn(b, "data")
	At(10, procesB).WantsToReceiveOn(b)

	At(14, procesA).Terminates()
	At(14, procesB).Terminates()

	// send svg to stderr
	fmt.Fprintln(os.Stderr, PrdEnd().String())

	// output:
	// at 0, process "procesA" starts with label "proces A"
	// at 1, process "procesA" creates proces "procesB" with label "proces B"
	// at 2, process "procesA" creates proces "procesC" with label "proces C"
	// at 3, process "procesA" wants to receive on channel "a"
	// at 5, process "procesC" wants to send "data" on channel "a"
	// - received by process "procesA"
	// at 6, process "procesA" wants to receive on channel "a"
	// at 8, process "procesB" wants to send "data" on channel "a"
	// - received by process "procesA"
}
