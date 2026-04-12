package packets

import (
	"io"

	"github.com/NaymDev/mcgotocol/codec"
)

type ServerStatusRequest struct{}

var _ Packet = (*ServerStatusRequest)(nil)

func (p *ServerStatusRequest) ID() int32 {
	return 0x00
}

func (p *ServerStatusRequest) Encode(writer io.Writer) error {
	return nil
}

func (p *ServerStatusRequest) Decode(reader io.Reader) error {
	return nil
}

type ServerStatusPing struct {
	Payload int64
}

var _ Packet = (*ServerStatusPing)(nil)

func (p *ServerStatusPing) ID() int32 {
	return 0x01
}

func (p *ServerStatusPing) Encode(writer io.Writer) error {
	return codec.WriteLong(writer, p.Payload)
}

func (p *ServerStatusPing) Decode(reader io.Reader) error {
	var err error
	p.Payload, err = codec.ReadLong(reader)
	return err
}

type ClientStatusResponse struct {
	JSONResponse string
}

var _ Packet = (*ClientStatusResponse)(nil)

func (p ClientStatusResponse) ID() int32 {
	return 0x00
}

func (p ClientStatusResponse) Encode(writer io.Writer) error {
	return codec.WriteString(writer, p.JSONResponse)
}

func (p ClientStatusResponse) Decode(reader io.Reader) error {
	var err error
	p.JSONResponse, err = codec.ReadString(reader)
	return err
}

type ClientStatusPong struct {
	Payload int64
}

var _ Packet = (*ClientStatusPong)(nil)

func (p *ClientStatusPong) ID() int32 {
	return 0x01
}

func (p *ClientStatusPong) Encode(writer io.Writer) error {
	return codec.WriteLong(writer, p.Payload)
}

func (p *ClientStatusPong) Decode(reader io.Reader) error {
	var err error
	p.Payload, err = codec.ReadLong(reader)
	return err
}
