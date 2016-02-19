Example PRoces Diagrams
=======================

![](goingFor.svg)

is produced by the following Go code:

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

See `goingFor.go`.

Meanwhile, it produces the following log:

	** at 0, proces "going for a walk" starts
	at 1, proces "going for a walk" creates proces "Alice getting ready"
	at 2, proces "going for a walk" creates proces "Bob getting ready"
	at 3, proces "going for a walk" wants to receive on channel "got ready"
	at 6, proces "Bob getting ready" wants to send "true" on channel "got ready"
	- received by proces "going for a walk"
	at 7, proces "going for a walk" wants to receive on channel "got ready"
	at 12, proces "Alice getting ready" wants to send "true" on channel "got ready"
	- received by proces "going for a walk"
	at 20, proces "going for a walk" terminates

