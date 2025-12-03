package mcgotocol

import (
	"bytes"
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state"
	"io"
	"net"
)

type Connection struct {
	serverBoundPacketRegistry *state.PacketRegistry
	conn                      io.ReadWriter
}

func NewConnection(conn io.ReadWriter, registry *state.Registry) *Connection {
	return &Connection{
		conn:                      conn,
		serverBoundPacketRegistry: registry.ServerBound,
	}
}

func (c *Connection) SetState(registry *state.Registry) {
	c.serverBoundPacketRegistry = registry.ServerBound
}

func (c *Connection) ReadPacket() (proto.Packet, error) {
	length, err := codec.ReadVarInt(c.conn)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, length)
	_, err = io.ReadFull(c.conn, buf)
	if err != nil {
		return nil, err
	}

	br := bytes.NewReader(buf)

	packetID, err := codec.ReadVarInt(br)
	if err != nil {
		return nil, err
	}

	return c.serverBoundPacketRegistry.Decode(int32(packetID), br)
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
	if err := codec.WriteVarInt(c.conn, codec.VarInt(len(packetData))); err != nil {
		return err
	}

	_, err := c.conn.Write(packetData)
	return err
}

func (c *Connection) RawConn() io.ReadWriter {
	return c.conn
}

func (c *Connection) Close() error {
	if closer, ok := c.conn.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (c *Connection) RemoteAddr() string {
	if addrProvider, ok := c.conn.(net.Conn); ok {
		return addrProvider.RemoteAddr().String()
	}
	return "unknown"
}

func (c *Connection) State() string {
	return c.serverBoundPacketRegistry.State
}
