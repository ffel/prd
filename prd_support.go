package prd

//go:generate stringer -type=Proces,Channel,processtate

const (
	procesA Proces = iota
	procesB
	procesC
)

const (
	a Channel = iota
	b
	c
)
