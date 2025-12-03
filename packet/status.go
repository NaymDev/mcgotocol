package packet

import (
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"io"
)

type ServerStatusRequest struct{}

var _ proto.Packet = (*ServerStatusRequest)(nil)

func (s *ServerStatusRequest) ID() int32 {
	return 0x00
}

func (s *ServerStatusRequest) Encode(writer io.Writer) error {
	return nil
}

func (s *ServerStatusRequest) Decode(reader io.Reader) error {
	return nil
}

type ClientStatusResponse struct {
	JSONResponse string
}

var _ proto.Packet = (*ClientStatusResponse)(nil)

func (c ClientStatusResponse) ID() int32 {
	return 0x00
}

func (c ClientStatusResponse) Encode(writer io.Writer) error {
	return codec.WriteString(writer, c.JSONResponse)
}

func (c ClientStatusResponse) Decode(reader io.Reader) error {
	var err error
	c.JSONResponse, err = codec.ReadString(reader)
	return err
}

type ServerStatusPing struct {
	Payload int64
}

var _ proto.Packet = (*ServerStatusPing)(nil)

func (s *ServerStatusPing) ID() int32 {
	return 0x01
}

func (s *ServerStatusPing) Encode(writer io.Writer) error {
	return codec.WriteLong(writer, s.Payload)
}

func (s *ServerStatusPing) Decode(reader io.Reader) error {
	var err error
	s.Payload, err = codec.ReadLong(reader)
	return err
}

type ClientStatusPong struct {
	Payload int64
}

var _ proto.Packet = (*ClientStatusPong)(nil)

func (c *ClientStatusPong) ID() int32 {
	return 0x01
}

func (c *ClientStatusPong) Encode(writer io.Writer) error {
	return codec.WriteLong(writer, c.Payload)
}

func (c *ClientStatusPong) Decode(reader io.Reader) error {
	var err error
	c.Payload, err = codec.ReadLong(reader)
	return err
}
