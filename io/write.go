package io

import (
	"bytes"
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
)

func WritePacket(w io.Writer, p proto.Packet) error {
	var buf bytes.Buffer

	if err := codec.WriteVarInt(&buf, codec.VarInt(p.ID())); err != nil {
		return err
	}

	if err := p.Encode(&buf); err != nil {
		return err
	}

	data := buf.Bytes()
	length := len(data)

	if err := codec.WriteVarInt(w, codec.VarInt(length)); err != nil {
		return err
	}

	_, err := w.Write(data)
	return err
}
