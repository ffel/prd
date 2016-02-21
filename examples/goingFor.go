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
	armAlarm
	BobPutsOnShoes
	AlicePutsOnShoes
)

const (
	gotReady Channel = iota
	putOnShoes
	armedAlarm
)

func main() {
	PrdStart(24, 6)

	LabelChannel(gotReady, "got ready")
	LabelChannel(armedAlarm, "armed alarm")
	LabelChannel(putOnShoes, "put on shoes")

	At(0, goingForAWalk).Starts("going for a walk")
	At(1, goingForAWalk).Creates(BobGettingReady, "Bob getting ready")
	At(2, goingForAWalk).Creates(AliceGettingReady, "Alice getting ready")

	At(3, goingForAWalk).WantsToReceiveOn(gotReady)
	At(6, BobGettingReady).WantsToSendOn(gotReady, "true")
	At(7, goingForAWalk).WantsToReceiveOn(gotReady)
	At(9, AliceGettingReady).WantsToSendOn(gotReady, "true")

	At(10, goingForAWalk).Creates(armAlarm, "arm alarm")
	At(11, goingForAWalk).Creates(BobPutsOnShoes, "Bob puts on shoes")
	At(12, goingForAWalk).Creates(AlicePutsOnShoes, "Alice puts on shoes")

	At(13, goingForAWalk).WantsToReceiveOn(putOnShoes)

	At(16, BobPutsOnShoes).WantsToSendOn(putOnShoes, "true")
	At(17, goingForAWalk).WantsToReceiveOn(putOnShoes)
	At(19, AlicePutsOnShoes).WantsToSendOn(putOnShoes, "true")

	At(20, goingForAWalk).WantsToReceiveOn(armedAlarm)
	At(22, armAlarm).WantsToSendOn(armedAlarm, "true")
	At(24, goingForAWalk).Terminates()

	PrdEnd()

	fmt.Fprintln(os.Stderr, Log.String())
	fmt.Println(SVG.String())
}
