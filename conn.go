package mcgotocol

import (
	"bytes"
	"compress/zlib"
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state"
	"io"
)

type Connection struct {
	writer io.Writer
	reader io.Reader

	packetRegistry *state.PacketRegistry

	compressionThreshold int
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

func (c *Connection) SetCompressionThreshold(threshold int) {
	c.compressionThreshold = threshold
}

func (c *Connection) ReadPacket() (proto.Packet, error) {
	length, err := codec.ReadVarInt(c.reader)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	if _, err := io.ReadFull(c.reader, buf); err != nil {
		return nil, err
	}

	br := bytes.NewReader(buf)

	var dataReader io.Reader = br

	if c.compressionThreshold >= 0 {
		dataLength, err := codec.ReadVarInt(br)
		if err != nil {
			return nil, err
		}

		if dataLength != 0 {
			zr, err := zlib.NewReader(br)
			if err != nil {
				return nil, err
			}
			defer zr.Close()

			decompressed, err := io.ReadAll(zr)
			if err != nil {
				return nil, err
			}

			dataReader = bytes.NewReader(decompressed)
		}
	}

	packetID, err := codec.ReadVarInt(dataReader)
	if err != nil {
		return nil, err
	}

	return c.packetRegistry.Decode(int32(packetID), dataReader)
}

func (c *Connection) WritePacket(p proto.Packet) error {
	var buf bytes.Buffer

	if err := codec.WriteVarInt(&buf, codec.VarInt(p.ID())); err != nil {
		return err
	}

	if err := p.Encode(&buf); err != nil {
		return err
	}

	packetData := buf.Bytes()

	var dataBuf bytes.Buffer

	if c.compressionThreshold >= 0 {
		if len(packetData) >= c.compressionThreshold {
			if err := codec.WriteVarInt(&dataBuf, codec.VarInt(len(packetData))); err != nil {
				return err
			}

			zw := zlib.NewWriter(&dataBuf)
			if _, err := zw.Write(packetData); err != nil {
				return err
			}
			zw.Close()
		} else {
			if err := codec.WriteVarInt(&dataBuf, 0); err != nil {
				return err
			}
			dataBuf.Write(packetData)
		}
	} else {
		dataBuf.Write(packetData)
	}

	if err := codec.WriteVarInt(c.writer, codec.VarInt(dataBuf.Len())); err != nil {
		return err
	}

	_, err := c.writer.Write(dataBuf.Bytes())
	return err
}
