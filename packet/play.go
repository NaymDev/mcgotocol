package packet

import (
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/google/uuid"
)

type ServerKeepAlive struct {
	KeepAliveID codec.VarInt
}

var _ proto.Packet = (*ServerKeepAlive)(nil)

func (p *ServerKeepAlive) ID() int32 {
	return 0x00
}

func (p *ServerKeepAlive) Encode(writer io.Writer) error {
	return codec.WriteVarInt(writer, p.KeepAliveID)
}

func (p *ServerKeepAlive) Decode(reader io.Reader) error {
	var err error
	p.KeepAliveID, err = codec.ReadVarInt(reader)
	return err
}

type ClientKeepAlive struct {
	KeepAliveID codec.VarInt
}

var _ proto.Packet = (*ClientKeepAlive)(nil)

func (p *ClientKeepAlive) ID() int32 {
	return 0x00
}

func (p *ClientKeepAlive) Encode(writer io.Writer) error {
	return codec.WriteVarInt(writer, p.KeepAliveID)
}

func (p *ClientKeepAlive) Decode(reader io.Reader) error {
	var err error
	p.KeepAliveID, err = codec.ReadVarInt(reader)
	return err
}

type ClientJoinGame struct {
	EntityID         int32
	Gamemode         uint8
	Dimension        int8
	Difficulty       uint8
	MaxPlayers       uint8
	LevelType        string
	ReducedDebugInfo bool
}

var _ proto.Packet = (*ClientJoinGame)(nil)

func (p *ClientJoinGame) ID() int32 {
	return 0x01
}

func (p *ClientJoinGame) Encode(writer io.Writer) error {
	if err := codec.WriteInt(writer, p.EntityID); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, uint8(p.Gamemode)); err != nil {
		return err
	}
	if err := codec.WriteByte(writer, p.Dimension); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, p.Difficulty); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, p.MaxPlayers); err != nil {
		return err
	}
	if err := codec.WriteString(writer, p.LevelType); err != nil {
		return err
	}
	if err := codec.WriteBool(writer, p.ReducedDebugInfo); err != nil {
		return err
	}
	return nil
}

func (p *ClientJoinGame) Decode(reader io.Reader) error {
	var err error
	p.EntityID, err = codec.ReadInt(reader)
	if err != nil {
		return err
	}
	p.Gamemode, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	p.Dimension, err = codec.ReadByte(reader)
	if err != nil {
		return err
	}
	p.Difficulty, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	p.MaxPlayers, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	p.LevelType, err = codec.ReadString(reader)
	if err != nil {
		return err
	}
	p.ReducedDebugInfo, err = codec.ReadBool(reader)
	if err != nil {
		return err
	}
	return nil
}

type ClientChatMessage struct {
	JSONData codec.Chat
	Position int8
}

var _ proto.Packet = (*ClientChatMessage)(nil)

func (p *ClientChatMessage) ID() int32 {
	return 0x02
}

func (p *ClientChatMessage) Encode(writer io.Writer) error {
	if err := codec.WriteChat(writer, p.JSONData); err != nil {
		return err
	}
	if err := codec.WriteByte(writer, p.Position); err != nil {
		return err
	}
	return nil
}

func (p *ClientChatMessage) Decode(reader io.Reader) error {
	var err error
	if p.JSONData, err = codec.ReadChat(reader); err != nil {
		return err
	}
	if p.Position, err = codec.ReadByte(reader); err != nil {
		return err
	}
	return nil
}

type ClientSetSpawnPosition struct {
	X int32
	Y int32
	Z int32
}

var _ proto.Packet = (*ClientSetSpawnPosition)(nil)

func (p *ClientSetSpawnPosition) ID() int32 {
	return 0x05
}

func (p *ClientSetSpawnPosition) Encode(writer io.Writer) error {
	return codec.WritePosition(writer, p.X, p.Y, p.Z)
}

func (p *ClientSetSpawnPosition) Decode(reader io.Reader) error {
	var err error
	p.X, p.Y, p.Z, err = codec.ReadPosition(reader)
	return err
}

type ClientPlayerPositionAndLookFlag uint8

const (
	X ClientPlayerPositionAndLookFlag = 1 << iota
	Y
	Z
	YRot
	XRot
)

type ClientPlayerPositionAndLook struct {
	X     float64
	Y     float64
	Z     float64
	Yaw   float32
	Pitch float32
	Flags uint8
}

var _ proto.Packet = (*ClientPlayerPositionAndLook)(nil)

func (p *ClientPlayerPositionAndLook) ID() int32 {
	return 0x08
}

func (p *ClientPlayerPositionAndLook) Encode(writer io.Writer) error {
	if err := codec.WriteDouble(writer, p.X); err != nil {
		return err
	}
	if err := codec.WriteDouble(writer, p.Y); err != nil {
		return err
	}
	if err := codec.WriteDouble(writer, p.Z); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, p.Yaw); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, p.Pitch); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, p.Flags); err != nil {
		return err
	}
	return nil
}

func (p *ClientPlayerPositionAndLook) Decode(reader io.Reader) error {
	var err error
	if p.X, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if p.Y, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if p.Z, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if p.Yaw, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if p.Pitch, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if p.Flags, err = codec.ReadUByte(reader); err != nil {
		return err
	}
	return nil
}

type ClientSpawnPainting struct {
	EntityID  codec.VarInt
	Title     string
	X         int32
	Y         int32
	Z         int32
	Direction codec.Direction
}

var _ proto.Packet = (*ClientSpawnPainting)(nil)

func (p *ClientSpawnPainting) ID() int32 {
	return 0x10
}

func (p *ClientSpawnPainting) Encode(writer io.Writer) error {
	if len(p.Title) > 13 {
		return codec.ErrStringTooLong
	}

	if err := codec.WriteVarInt(writer, p.EntityID); err != nil {
		return err
	}
	if err := codec.WriteString(writer, p.Title); err != nil {
		return err
	}
	if err := codec.WritePosition(writer, p.X, p.Y, p.Z); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, uint8(p.Direction)); err != nil {
		return err
	}
	return nil
}

func (p *ClientSpawnPainting) Decode(reader io.Reader) error {
	var err error
	if p.EntityID, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if p.Title, err = codec.ReadString(reader); err != nil {
		return err
	}
	if len(p.Title) > 13 {
		return codec.ErrStringTooLong
	}
	if p.X, p.Y, p.Z, err = codec.ReadPosition(reader); err != nil {
		return err
	}
	if d, err := codec.ReadUByte(reader); err != nil {
		return err
	} else {
		p.Direction = codec.Direction(d)
	}
	return nil
}

type ClientSpawnPlayer struct {
	EntityID    codec.VarInt
	PlayerUUID  uuid.UUID
	X           int32
	Y           int32
	Z           int32
	Yaw         codec.Angle
	Pitch       codec.Angle
	CurrentItem int16
	Metadata    []codec.EntityMetadata
}

var _ proto.Packet = (*ClientSpawnPlayer)(nil)

func (p *ClientSpawnPlayer) ID() int32 {
	return 0x0C
}

func (p *ClientSpawnPlayer) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, p.EntityID); err != nil {
		return err
	}
	if err := codec.WriteUUID(writer, p.PlayerUUID); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.X); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.Y); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.Z); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, p.Yaw); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, p.Pitch); err != nil {
		return err
	}
	if err := codec.WriteShort(writer, p.CurrentItem); err != nil {
		return err
	}
	if err := codec.WriteMetadata(writer, p.Metadata); err != nil {
		return err
	}
	return nil
}

func (p *ClientSpawnPlayer) Decode(reader io.Reader) error {
	var err error
	if p.EntityID, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if p.PlayerUUID, err = codec.ReadUUID(reader); err != nil {
		return err
	}
	if p.X, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Y, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Z, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Yaw, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if p.Pitch, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if p.CurrentItem, err = codec.ReadShort(reader); err != nil {
		return err
	}
	if p.Metadata, err = codec.ReadMetadata(reader); err != nil {
		return err
	}
	return nil
}

type ObjectType int8

// TODO: make enum from https://minecraft.fandom.com/wiki/Java_Edition_data_values/Pre-flattening/Entity_IDs

type ClientSpawnObject struct {
	EntityID  codec.VarInt
	Type      ObjectType
	X         int32
	Y         int32
	Z         int32
	Pitch     codec.Angle
	Yaw       codec.Angle
	Data      int32
	VelocityX int32
	VelocityY int32
	VelocityZ int32
}

var _ proto.Packet = (*ClientSpawnObject)(nil)

func (p *ClientSpawnObject) ID() int32 {
	return 0x0E
}

func (p *ClientSpawnObject) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, p.EntityID); err != nil {
		return err
	}
	if err := codec.WriteByte(writer, int8(p.Type)); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.X); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.Y); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.Z); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, p.Pitch); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, p.Yaw); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, p.Data); err != nil {
		return err
	}
	if p.Data != 0 {
		if err := codec.WriteInt(writer, p.VelocityX); err != nil {
			return err
		}
		if err := codec.WriteInt(writer, p.VelocityY); err != nil {
			return err
		}
		if err := codec.WriteInt(writer, p.VelocityZ); err != nil {
			return err
		}
	}
	return nil
}

func (p *ClientSpawnObject) Decode(reader io.Reader) error {
	var err error
	if p.EntityID, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if b, err := codec.ReadByte(reader); err != nil {
		return err
	} else {
		p.Type = ObjectType(b)
	}
	if p.X, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Y, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Z, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Pitch, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if p.Yaw, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if p.Data, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if p.Data != 0 {
		if p.VelocityX, err = codec.ReadInt(reader); err != nil {
			return err
		}
		if p.VelocityY, err = codec.ReadInt(reader); err != nil {
			return err
		}
		if p.VelocityZ, err = codec.ReadInt(reader); err != nil {
			return err
		}
	}
	return nil
}

type ClientPlayerAbilities struct {
	Flags               int8
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

var _ proto.Packet = (*ClientPlayerAbilities)(nil)

func (p *ClientPlayerAbilities) ID() int32 {
	return 0x39
}

func (p *ClientPlayerAbilities) Encode(writer io.Writer) error {
	if err := codec.WriteByte(writer, p.Flags); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, p.FlyingSpeed); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, p.FieldOfViewModifier); err != nil {
		return err
	}
	return nil
}

func (p *ClientPlayerAbilities) Decode(reader io.Reader) error {
	var err error
	if p.Flags, err = codec.ReadByte(reader); err != nil {
		return err
	}
	if p.FlyingSpeed, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if p.FieldOfViewModifier, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	return nil
}

type ClientPlayerListHeaderAndFooter struct {
	Header codec.Chat
	Footer codec.Chat
}

var _ proto.Packet = (*ClientPlayerListHeaderAndFooter)(nil)

func (p *ClientPlayerListHeaderAndFooter) ID() int32 {
	return 0x47
}

func (p *ClientPlayerListHeaderAndFooter) Encode(writer io.Writer) error {
	if err := codec.WriteChat(writer, p.Header); err != nil {
		return err
	}
	if err := codec.WriteChat(writer, p.Footer); err != nil {
		return err
	}
	return nil
}

func (p *ClientPlayerListHeaderAndFooter) Decode(reader io.Reader) error {
	var err error
	if p.Header, err = codec.ReadChat(reader); err != nil {
		return err
	}
	if p.Footer, err = codec.ReadChat(reader); err != nil {
		return err
	}
	return nil
}
