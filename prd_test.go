package prd

import (
	"fmt"
	"os"
)

const (
	goingForAWalk Proces = iota
	AliceGettingReady
	BobGettingReady
)

const (
	a Channel = iota
	b
	c
)

// go test 2> out.svg
func Example() {
	PrdStart(800, 500)

	At(0, goingForAWalk).Starts("going for a Walk")

	At(1, goingForAWalk).Creates(AliceGettingReady, "Alice getting ready")
	At(2, goingForAWalk).Creates(BobGettingReady, "Bob getting ready")

	// situatie 1: receive moet wachten op een zender

	// overweeg toch een WantsToSend("data", a) en WantsToReceive(a)
	// of WantsToSendOn(a, "data") en WantsToReceiveOn(a)

	// controleer ook even de imports: `import . prd`

	At(3, AliceGettingReady).WantsToReceiveOn(a)
	At(3, BobGettingReady).WantsToReceiveOn(a)

	At(4, goingForAWalk).WantsToSendOn(a, "data").HandledBy(AliceGettingReady)

	// situatie 2: send moet wachten op een ontvanger

	At(4, AliceGettingReady).WantsToSendOn(b, "data")
	At(5, BobGettingReady).WantsToSendOn(b, "data")

	At(7, goingForAWalk).WantsToReceiveOn(b).HandledBy(AliceGettingReady)

	At(10, goingForAWalk).Terminates()
	At(10, BobGettingReady).Terminates()
	At(10, AliceGettingReady).Terminates()

	fmt.Fprintln(os.Stderr, PrdEnd().String())

	// output:
	// at 0, process "goingForAWalk" starts with label "going for a Walk"
	// at 1, process "goingForAWalk" creates proces "AliceGettingReady" with label "Alice getting ready"
	// at 2, process "goingForAWalk" creates proces "BobGettingReady" with label "Bob getting ready"
	// at 3, process "AliceGettingReady" wants to receive on channel "a"
	// at 3, process "BobGettingReady" wants to receive on channel "a"
	// at 4, process "goingForAWalk" wants to send "data" on channel "a"
	// -  handled by "AliceGettingReady"
	// at 4, process "AliceGettingReady" wants to send "data" on channel "b"
	// at 5, process "BobGettingReady" wants to send "data" on channel "b"
	// at 7, process "goingForAWalk" wants to receive on channel "b"
	// -  handled by "AliceGettingReady"
	// at 10, process "goingForAWalk" terminates
	// at 10, process "BobGettingReady" terminates
	// at 10, process "AliceGettingReady" terminates
}
