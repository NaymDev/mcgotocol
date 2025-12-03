package packet

import (
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"io"
)

type ServerLoginStart struct {
	Name string
}

var _ proto.Packet = (*ServerLoginStart)(nil)

func (s *ServerLoginStart) ID() int32 {
	return 0x00
}

func (s *ServerLoginStart) Encode(writer io.Writer) error {
	return codec.WriteString(writer, s.Name)
}

func (s *ServerLoginStart) Decode(reader io.Reader) error {
	var err error
	s.Name, err = codec.ReadString(reader)
	return err
}

type ClientLoginSuccess struct {
	UUID     string
	Username string
}

var _ proto.Packet = (*ClientLoginSuccess)(nil)

func (c *ClientLoginSuccess) ID() int32 {
	return 0x02
}

func (c *ClientLoginSuccess) Encode(writer io.Writer) error {
	if err := codec.WriteString(writer, c.UUID); err != nil {
		return err
	}
	if err := codec.WriteString(writer, c.Username); err != nil {
		return err
	}
	return nil
}

func (c *ClientLoginSuccess) Decode(reader io.Reader) error {
	var err error
	if c.UUID, err = codec.ReadString(reader); err != nil {
		return err
	}
	if c.Username, err = codec.ReadString(reader); err != nil {
		return err
	}
	return nil
}
