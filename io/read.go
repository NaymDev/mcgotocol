package io

import (
	"bytes"
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/packets"
	"github.com/NaymDev/mcgotocol/state"
)

func ReadPacket(r io.Reader, registry *state.PacketRegistry) (packets.Packet, error) {
	length, err := codec.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return nil, err
	}

	br := bytes.NewReader(buf)

	packetID, err := codec.ReadVarInt(br)
	if err != nil {
		return nil, err
	}

	return registry.Decode(int32(packetID), br)
}
