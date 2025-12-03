package proto

import "io"

type Packet interface {
	ID() int32
	Encode(io.Writer) error
	Decode(io.Reader) error
}

type Direction uint8

const (
	ClientBound Direction = iota
	ServerBound
)
