package prd

//go:generate stringer -type=Proces,Channel,processtate

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
