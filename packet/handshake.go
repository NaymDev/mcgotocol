package packet

import (
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state/states"
	"io"
)

type HandshakeIntent codec.VarInt

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

func (s *ServerHandshake) ID() int32 {
	return 0x00
}

func (s *ServerHandshake) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, s.ProtocolVersion); err != nil {
		return err
	}
	if err := codec.WriteString(writer, s.ServerAddress); err != nil {
		return err
	}
	if err := codec.WriteUShort(writer, s.ServerPort); err != nil {
		return err
	}
	if err := codec.WriteVarInt(writer, s.NextState); err != nil {
		return err
	}
	return nil
}

func (s *ServerHandshake) Decode(reader io.Reader) error {
	var err error
	if s.ProtocolVersion, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if s.ServerAddress, err = codec.ReadString(reader); err != nil {
		return err
	}
	if s.ServerPort, err = codec.ReadUShort(reader); err != nil {
		return err
	}
	if s.NextState, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	return nil
}
