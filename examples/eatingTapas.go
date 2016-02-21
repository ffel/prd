package main

import (
	"fmt"
	"os"

	. "github.com/ffel/prd"
)

const (
	eatingTapas Proces = iota
	chorizo
	olivas
	bobEating
	charlieEating
)

const (
	morsel Channel = iota
	empty
	allDone
)

func main() {
	PrdStart(24, 5)

	LabelChannel(morsel, "morsel")
	LabelChannel(empty, "empty")
	LabelChannel(allDone, "all done")

	At(0, eatingTapas).Starts("eatin' tapas")
	At(1, eatingTapas).Creates(chorizo, "chorizo")
	At(2, chorizo).WantsToSendOn(morsel, "a bite")
	At(3, eatingTapas).Creates(olivas, "olivas")
	At(4, olivas).WantsToSendOn(morsel, "a bite")

	At(5, eatingTapas).Creates(bobEating, "Bob")
	At(6, eatingTapas).Creates(charlieEating, "Charlie")

	At(7, eatingTapas).WantsToReceiveOn(empty)

	// both chorizo as well as olivas are willing to send
	// so, we have to decide which
	At(8, bobEating).
		AsServedByProces(chorizo).
		WantsToReceiveOn(morsel).
		AndToReceiveOn(allDone)

	At(9, chorizo).
		WantsToSendOn(morsel, "a bite")

	At(11, charlieEating).
		AsServedByProces(olivas).
		WantsToReceiveOn(morsel).
		AndToReceiveOn(allDone)
	At(12, olivas).WantsToSendOn(morsel, "a last bite")

	At(14, bobEating).AsServedByProces(olivas).
		WantsToReceiveOn(morsel).
		AndToReceiveOn(allDone)
	At(15, olivas).WantsToSendOn(empty, "done")
	// At(18, bobEating).WantsToReceiveOn(morsel) //.AndToReceiveOn(allDone)

	At(16, eatingTapas).WantsToReceiveOn(empty)

	At(17, charlieEating).
		WantsToReceiveOn(morsel).AndToReceiveOn(allDone)
	At(18, chorizo).WantsToSendOn(empty, "done")

	// At(18, bobEating).WantsToReceiveOn(morsel).AndToReceiveOn(allDone)

	// closes???, X als mesg?
	At(19, eatingTapas).WantsToSendOn(allDone, "")

	// At(22, charlieEating).WantsToReceiveOn(morsel).AndToReceiveOn(allDone)

	At(24, eatingTapas).Terminates()

	PrdEnd()

	fmt.Fprintln(os.Stderr, Log.String())
	fmt.Println(SVG.String())
}
