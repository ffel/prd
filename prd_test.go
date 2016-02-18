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

	// situatie 1: receive moet wachten op een zender

	// overweeg toch een WantsToSend("data", a) en WantsToReceive(a)
	// of WantsToSendOn(a, "data") en WantsToReceiveOn(a)

	// controleer ook even de imports: `import . prd`

	// At(3, procesB).WantsToReceiveOn(a)
	// At(4, procesC).WantsToReceiveOn(a)

	// At(4, procesA).WantsToSendOn(a, "data").HandledBy(procesB)

	// // situatie 2: send moet wachten op een ontvanger

	// At(4, procesB).WantsToSendOn(b, "data")
	// At(5, procesC).WantsToSendOn(b, "data")

	// At(7, procesA).WantsToReceiveOn(b).HandledBy(procesB)

	// At(10, procesA).Terminates()
	// At(10, procesC).Terminates()
	// At(10, procesB).Terminates()

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
