package mcgotocol

import (
	"bytes"
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state"
	"io"
)

type Connection struct {
	writer         io.Writer
	reader         io.Reader
	packetRegistry *state.PacketRegistry
}

func NewConnection(reader io.Reader, writer io.Writer, packetRegistry *state.PacketRegistry) *Connection {
	return &Connection{
		writer:         writer,
		reader:         reader,
		packetRegistry: packetRegistry,
	}
}

func (c *Connection) SetPacketRegistry(packetRegistry *state.PacketRegistry) {
	c.packetRegistry = packetRegistry
}

func (c *Connection) ReadPacket() (proto.Packet, error) {
	length, err := codec.ReadVarInt(c.reader)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	_, err = io.ReadFull(c.reader, buf)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(buf)

	packetID, err := codec.ReadVarInt(br)
	if err != nil {
		return nil, err
	}

	return c.packetRegistry.Decode(int32(packetID), br)
}

func (c *Connection) WritePacket(p proto.Packet) error {
	buf := &bytes.Buffer{}

	if err := codec.WriteVarInt(buf, codec.VarInt(p.ID())); err != nil {
		return err
	}

	if err := p.Encode(buf); err != nil {
		return err
	}

	packetData := buf.Bytes()
	if err := codec.WriteVarInt(c.writer, codec.VarInt(len(packetData))); err != nil {
		return err
	}

	_, err := c.writer.Write(packetData)
	return err
}
