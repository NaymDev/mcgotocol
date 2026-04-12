package io

import (
	"bytes"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/packets"
)

func MarshalPacket(p packets.Packet) ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := codec.WriteVarInt(buf, codec.VarInt(p.ID())); err != nil {
		return nil, err
	}

	if err := p.Encode(buf); err != nil {
		return nil, err
	}

	packetData := buf.Bytes()

	finalBuf := &bytes.Buffer{}
	if err := codec.WriteVarInt(finalBuf, codec.VarInt(len(packetData))); err != nil {
		return nil, err
	}

	finalBuf.Write(packetData)
	return finalBuf.Bytes(), nil
}
