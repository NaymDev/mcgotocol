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

func (p *ServerLoginStart) ID() int32 {
	return 0x00
}

func (p *ServerLoginStart) Encode(writer io.Writer) error {
	return codec.WriteString(writer, p.Name)
}

func (p *ServerLoginStart) Decode(reader io.Reader) error {
	var err error
	p.Name, err = codec.ReadString(reader)
	return err
}

type ServerEncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

var _ proto.Packet = (*ServerEncryptionResponse)(nil)

func (p *ServerEncryptionResponse) ID() int32 {
	return 0x01
}

func (p *ServerEncryptionResponse) Encode(writer io.Writer) error {
	if err := codec.WriteByteArray(writer, p.SharedSecret); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, p.VerifyToken); err != nil {
		return err
	}
	return nil
}

func (p *ServerEncryptionResponse) Decode(reader io.Reader) error {
	var err error
	if p.SharedSecret, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	if p.VerifyToken, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	return nil
}

type ClientDisconnect struct {
	Reason codec.Chat
}

var _ proto.Packet = (*ClientDisconnect)(nil)

func (p *ClientDisconnect) ID() int32 {
	return 0x00
}

func (p *ClientDisconnect) Encode(writer io.Writer) error {
	return codec.WriteChat(writer, p.Reason)
}

func (p *ClientDisconnect) Decode(reader io.Reader) error {
	var err error
	p.Reason, err = codec.ReadChat(reader)
	return err
}

type ClientEncryptionRequest struct {
	ServerID    string
	PublicKey   []byte
	VerifyToken []byte
}

var _ proto.Packet = (*ClientEncryptionRequest)(nil)

func (p *ClientEncryptionRequest) ID() int32 {
	return 0x01
}

func (p *ClientEncryptionRequest) Encode(writer io.Writer) error {
	if err := codec.WriteString(writer, p.ServerID); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, p.PublicKey); err != nil {
		return err
	}
	if err := codec.WriteByteArray(writer, p.VerifyToken); err != nil {
		return err
	}
	return nil
}

func (p *ClientEncryptionRequest) Decode(reader io.Reader) error {
	var err error
	if p.ServerID, err = codec.ReadString(reader); err != nil {
		return err
	}
	if p.PublicKey, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	if p.VerifyToken, err = codec.ReadByteArray(reader); err != nil {
		return err
	}
	return nil
}

type ClientLoginSuccess struct {
	UUID     string
	Username string
}

var _ proto.Packet = (*ClientLoginSuccess)(nil)

func (p *ClientLoginSuccess) ID() int32 {
	return 0x02
}

func (p *ClientLoginSuccess) Encode(writer io.Writer) error {
	if err := codec.WriteString(writer, p.UUID); err != nil {
		return err
	}
	if err := codec.WriteString(writer, p.Username); err != nil {
		return err
	}
	return nil
}

func (p *ClientLoginSuccess) Decode(reader io.Reader) error {
	var err error
	if p.UUID, err = codec.ReadString(reader); err != nil {
		return err
	}
	if p.Username, err = codec.ReadString(reader); err != nil {
		return err
	}
	return nil
}

type ClientSetCompression struct {
	Threshold codec.VarInt
}

var _ proto.Packet = (*ClientSetCompression)(nil)

func (p *ClientSetCompression) ID() int32 {
	return 0x03
}

func (p *ClientSetCompression) Encode(writer io.Writer) error {
	return codec.WriteVarInt(writer, p.Threshold)
}

func (p *ClientSetCompression) Decode(reader io.Reader) error {
	var err error
	p.Threshold, err = codec.ReadVarInt(reader)
	return err
}
