package state

import "fmt"

type UnknownPacketID struct {
	PacketID int32
	State    string
}

var _ error = (*UnknownPacketID)(nil)

func (e *UnknownPacketID) Error() string {
	return fmt.Sprintf("unknown packet ID 0x%X (State: %s)", e.PacketID, e.State)
}
