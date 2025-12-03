package codec

import (
	"bytes"
	"github.com/NaymDev/mcgotocol/proto"
)

func MarshalPacket(p proto.Packet) ([]byte, error) {
	buf := &bytes.Buffer{}

	if err := WriteVarInt(buf, VarInt(p.ID())); err != nil {
		return nil, err
	}

	if err := p.Encode(buf); err != nil {
		return nil, err
	}

	packetData := buf.Bytes()

	finalBuf := &bytes.Buffer{}
	if err := WriteVarInt(finalBuf, VarInt(len(packetData))); err != nil {
		return nil, err
	}

	finalBuf.Write(packetData)
	return finalBuf.Bytes(), nil
}
