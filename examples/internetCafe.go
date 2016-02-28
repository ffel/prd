package main

import (
	"io/ioutil"

	. "github.com/ffel/prd"
)

const (
	internetCafe Proces = iota
	reception
	term1
	term2
	exit
)

const (
	toReception Channel = iota
	toTerminal
	toExit
)

func main() {

	PrdStart(25, 5)
	LabelChannel(toReception, "to reception")
	LabelChannel(toTerminal, "to terminal")
	LabelChannel(toExit, "to exit")

	At(0, internetCafe).Starts("internet café")
	At(1, internetCafe).Creates(reception, "reception")
	At(1, internetCafe).Creates(term1, "terminal 1")
	At(1, internetCafe).Creates(term2, "terminal 2")
	At(1, internetCafe).Creates(exit, "exit")

	At(2, reception).WantsToReceiveOn(toReception)
	At(2, term1).WantsToReceiveOn(toTerminal)
	At(2, term2).WantsToReceiveOn(toTerminal)
	At(2, exit).WantsToReceiveOn(toExit)

	At(3, internetCafe).WantsToSendOn(toReception, "al")
	At(4, reception).AsServedBy(term1, toTerminal).
		WantsToSendOn(toTerminal, "al").
		AndDoesNotWait()
	At(14, term1).WantsToSendOn(toExit, "al")
	At(15, exit).WantsToReceiveOn(toExit)

	At(5, reception).WantsToReceiveOn(toReception)

	At(7, internetCafe).WantsToSendOn(toReception, "bob")
	At(8, reception).AsServedBy(term2, toTerminal).
		WantsToSendOn(toTerminal, "bob").
		AndDoesNotWait()
	At(16, term2).WantsToSendOn(toExit, "bob")
	At(17, exit).WantsToReceiveOn(toExit)
	At(17, term2).WantsToReceiveOn(toTerminal)

	At(9, reception).WantsToReceiveOn(toReception)
	At(11, internetCafe).WantsToSendOn(toReception, "chuck")
	At(12, reception).WantsToSendOn(toTerminal, "chuck").
		AndDoesNotWait()
	At(13, reception).WantsToReceiveOn(toReception).AndToSendOn(toTerminal, "chuck")

	At(15, term1).WantsToReceiveOn(toTerminal)
	At(16, reception).WantsToReceiveOn(toReception)

	At(21, term1).WantsToSendOn(toExit, "chuck")
	At(22, term1).WantsToReceiveOn(toTerminal)
	At(22, exit).WantsToReceiveOn(toExit)

	At(25, internetCafe).Terminates()
	At(25, reception).Terminates()
	At(25, term1).Terminates()
	At(25, term2).Terminates()
	At(25, exit).Terminates()

	PrdEnd()
	ioutil.WriteFile("internetCafe.svg", SVG.Bytes(), 0644)
}