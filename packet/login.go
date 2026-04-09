package packet

import (
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
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

type ClientEncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

var _ proto.Packet = (*ClientEncryptionRequest)(nil)

func (c *ClientEncryptionRequest) ID() int32 {
	return 0x01
}

func (c *ClientEncryptionRequest) Encode(writer io.Writer) error {
	if err := codec.WriteString(writer, c.ServerID); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, c.PublicKey); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, c.VerifyToken); err != nil {
		return err
	}
	return nil
}

func (c *ClientEncryptionRequest) Decode(reader io.Reader) error {
	var err error
	if c.ServerID, err = codec.ReadString(reader); err != nil {
		return err
	}
	if c.PublicKey, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	if c.VerifyToken, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	return nil
}

type ServerEncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

var _ proto.Packet = (*ServerEncryptionResponse)(nil)

func (s *ServerEncryptionResponse) ID() int32 {
	return 0x01
}

func (s *ServerEncryptionResponse) Encode(writer io.Writer) error {
	if err := codec.WriteByteArray(writer, s.SharedSecret); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, s.VerifyToken); err != nil {
		return err
	}
	return nil
}

func (s *ServerEncryptionResponse) Decode(reader io.Reader) error {
	var err error
	if s.SharedSecret, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	if s.VerifyToken, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	return nil
}
