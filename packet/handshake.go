package packet

import (
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state/states"
)

type HandshakeIntent = codec.VarInt

const (
	StatusHandshakeIntent = HandshakeIntent(states.StatusState)
	LoginHandshakeIntent  = HandshakeIntent(states.LoginState)
)

type ServerHandshake struct {
	ProtocolVersion codec.VarInt
	ServerAddress   string
	ServerPort      uint16
	NextState       codec.VarInt
}

var _ proto.Packet = (*ServerHandshake)(nil)

func (p *ServerHandshake) ID() int32 {
	return 0x00
}

func (p *ServerHandshake) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, p.ProtocolVersion); err != nil {
		return err
	}
	if err := codec.WriteString(writer, p.ServerAddress); err != nil {
		return err
	}
	if err := codec.WriteUShort(writer, p.ServerPort); err != nil {
		return err
	}
	if err := codec.WriteVarInt(writer, p.NextState); err != nil {
		return err
	}
	return nil
}

func (p *ServerHandshake) Decode(reader io.Reader) error {
	var err error
	if p.ProtocolVersion, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if p.ServerAddress, err = codec.ReadString(reader); err != nil {
		return err
	}
	if p.ServerPort, err = codec.ReadUShort(reader); err != nil {
		return err
	}
	if p.NextState, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	return nil
}
