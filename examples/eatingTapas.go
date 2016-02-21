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
	quit
)

func main() {
	PrdStart(24, 5)

	LabelChannel(morsel, "morsel")

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
	At(8, bobEating).AsServedByProces(chorizo).WantsToReceiveOn(morsel)
	At(9, chorizo).WantsToSendOn(morsel, "a bite")

	At(11, charlieEating).AsServedByProces(olivas).WantsToReceiveOn(morsel)
	At(12, olivas).WantsToSendOn(morsel, "a last bite")

	At(14, bobEating).AsServedByProces(olivas).WantsToReceiveOn(morsel)
	At(15, olivas).WantsToSendOn(empty, "done")

	At(16, eatingTapas).WantsToReceiveOn(empty)

	At(17, charlieEating).WantsToReceiveOn(morsel)
	At(18, chorizo).WantsToSendOn(empty, "done")

	// closes???, X als mesg?
	At(19, eatingTapas).WantsToSendOn(quit, "")

	At(24, eatingTapas).Terminates()

	PrdEnd()

	fmt.Fprintln(os.Stderr, Log.String())
	fmt.Println(SVG.String())
}
