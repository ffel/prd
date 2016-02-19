package main

import (
	"fmt"
	"os"

	. "github.com/ffel/prd"
)

const (
	goingForAWalk Proces = iota
	BobGettingReady
	AliceGettingReady
)

const (
	gotReady Channel = iota
)

func main() {
	PrdStart(800, 250)

	LabelChannel(gotReady, "got ready")

	At(0, goingForAWalk).Starts("going for a walk")
	At(1, goingForAWalk).Creates(AliceGettingReady, "Alice getting ready")
	At(2, goingForAWalk).Creates(BobGettingReady, "Bob getting ready")

	At(3, goingForAWalk).WantsToReceiveOn(gotReady)
	At(6, BobGettingReady).WantsToSendOn(gotReady, "true")
	At(7, goingForAWalk).WantsToReceiveOn(gotReady)
	At(12, AliceGettingReady).WantsToSendOn(gotReady, "true")

	At(20, goingForAWalk).Terminates()

	PrdEnd()

	fmt.Fprintln(os.Stderr, Log.String())
	fmt.Println(SVG.String())
}
