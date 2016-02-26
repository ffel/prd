package main

import (
	"io/ioutil"

	. "github.com/ffel/prd"
)

const (
	internetCafe Proces = iota
	reception
	waiting
	term1
	term2
	exit
)

const (
	toReception Channel = iota
	toQueue
	toTerminal
	toExit
)

func main() {
	image_1()
	image_2()
}

func image_1() {
	PrdStart(25, 6)
	LabelChannel(toReception, "to reception")
	LabelChannel(toQueue, "to queue")
	LabelChannel(toTerminal, "to terminal")
	LabelChannel(toExit, "to exit")

	/*
		Wanneer een klant aankomt bij de receptie wordt deze aan de wachtrij
		toegevoegd wanneer de wachtrij niet leeg is.

		Wanneer er geen wachtrij is, wordt gekeken of er een terminal vrij is.

		Als er een terminal vrij is (als de toTerminal channel de tourist
		accepteert) dan gaat de tourist naar een terminal, anders gaat de
		tourist alsnog naar de wachtrij.

		Touristen in de wachtrij worden sowieso op toTerminal aangeboden.
	*/

	At(0, internetCafe).Starts("internet café")
	At(1, internetCafe).Creates(reception, "reception")
	// At(1, internetCafe).Creates(waiting, "waiting")
	At(1, internetCafe).Creates(term1, "terminal 1")
	At(1, internetCafe).Creates(term2, "terminal 2")
	At(1, internetCafe).Creates(exit, "exit")

	At(2, reception).WantsToReceiveOn(toReception)
	// At(2, waiting).WantsToReceiveOn(toQueue)
	At(2, term1).WantsToReceiveOn(toTerminal)
	At(2, term2).WantsToReceiveOn(toTerminal)
	At(2, exit).WantsToReceiveOn(toExit)

	At(3, internetCafe).WantsToSendOn(toReception, "al")
	At(4, reception).AsServedBy(term1, toTerminal).WantsToSendOn(toTerminal, "al")
	At(13, term1).WantsToSendOn(toExit, "al")
	At(14, exit).WantsToReceiveOn(toExit)

	At(5, reception).WantsToReceiveOn(toReception)

	At(6, internetCafe).WantsToSendOn(toReception, "bob")
	At(7, reception).AsServedBy(term2, toTerminal).WantsToSendOn(toTerminal, "bob")
	At(15, term2).WantsToSendOn(toExit, "bob")
	At(16, exit).WantsToReceiveOn(toExit)

	// nu moet iets gedaan worden wanneer op 9 de derde toerist komt
	At(8, reception).WantsToReceiveOn(toReception)
	At(9, internetCafe).WantsToSendOn(toReception, "chuck")
	At(10, reception).WantsToSendOn(toTerminal, "chuck") // no terminal available

	At(25, internetCafe).Terminates()

	PrdEnd()
	ioutil.WriteFile("internetCafe_1.svg", SVG.Bytes(), 0644)
}

func image_2() {
	PrdStart(25, 6)
	LabelChannel(toReception, "to reception")
	LabelChannel(toQueue, "to queue")
	LabelChannel(toTerminal, "to terminal")
	LabelChannel(toExit, "to exit")

	/*
		Wanneer een klant aankomt bij de receptie wordt deze aan de wachtrij
		toegevoegd wanneer de wachtrij niet leeg is.

		Wanneer er geen wachtrij is, wordt gekeken of er een terminal vrij is.

		Als er een terminal vrij is (als de toTerminal channel de tourist
		accepteert) dan gaat de tourist naar een terminal, anders gaat de
		tourist alsnog naar de wachtrij.

		Touristen in de wachtrij worden sowieso op toTerminal aangeboden.
	*/

	At(0, internetCafe).Starts("internet café")
	At(1, internetCafe).Creates(reception, "reception")
	At(1, internetCafe).Creates(waiting, "waiting")
	At(1, internetCafe).Creates(term1, "terminal 1")
	At(1, internetCafe).Creates(term2, "terminal 2")
	At(1, internetCafe).Creates(exit, "exit")

	At(2, reception).WantsToReceiveOn(toReception)
	At(2, waiting).WantsToReceiveOn(toQueue)
	At(2, term1).WantsToReceiveOn(toTerminal)
	At(2, term2).WantsToReceiveOn(toTerminal)
	At(2, exit).WantsToReceiveOn(toExit)

	At(3, internetCafe).WantsToSendOn(toReception, "al")
	At(4, reception).AsServedBy(term2, toTerminal).DoesNotWait().AndToSendOn(toTerminal, "al")
	// At(13, term1).WantsToSendOn(toExit, "al")
	// At(14, exit).WantsToReceiveOn(toExit)

	// At(5, reception).WantsToReceiveOn(toReception)

	// At(6, internetCafe).WantsToSendOn(toReception, "bob")
	// At(7, reception).AsServedBy(term2, toTerminal).WantsToSendOn(toTerminal, "bob")
	// At(15, term2).WantsToSendOn(toExit, "bob")
	// At(16, exit).WantsToReceiveOn(toExit)

	// // nu moet iets gedaan worden wanneer op 9 de derde toerist komt
	// At(8, reception).WantsToReceiveOn(toReception)
	// At(9, internetCafe).WantsToSendOn(toReception, "chuck")
	// At(10, reception).WantsToSendOn(toTerminal, "chuck") // no terminal available

	At(25, internetCafe).Terminates()

	PrdEnd()
	ioutil.WriteFile("internetCafe_2.svg", SVG.Bytes(), 0644)
}
