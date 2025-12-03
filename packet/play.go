package packet

import (
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/google/uuid"
	"io"
)

type ClientKeepAlive struct {
	KeepAliveID codec.VarInt
}

var _ proto.Packet = (*ClientKeepAlive)(nil)

func (c *ClientKeepAlive) ID() int32 {
	return 0x00
}

func (c *ClientKeepAlive) Encode(writer io.Writer) error {
	return codec.WriteVarInt(writer, c.KeepAliveID)
}

func (c *ClientKeepAlive) Decode(reader io.Reader) error {
	var err error
	c.KeepAliveID, err = codec.ReadVarInt(reader)
	return err
}

type ServerKeepAlive struct {
	KeepAliveID codec.VarInt
}

var _ proto.Packet = (*ServerKeepAlive)(nil)

func (c *ServerKeepAlive) ID() int32 {
	return 0x00
}

func (c *ServerKeepAlive) Encode(writer io.Writer) error {
	return codec.WriteVarInt(writer, c.KeepAliveID)
}

func (c *ServerKeepAlive) Decode(reader io.Reader) error {
	var err error
	c.KeepAliveID, err = codec.ReadVarInt(reader)
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

func (c *ClientJoinGame) ID() int32 {
	return 0x01
}

func (c *ClientJoinGame) Encode(writer io.Writer) error {
	if err := codec.WriteInt(writer, c.EntityID); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, uint8(c.Gamemode)); err != nil {
		return err
	}
	if err := codec.WriteByte(writer, c.Dimension); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, c.Difficulty); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, c.MaxPlayers); err != nil {
		return err
	}
	if err := codec.WriteString(writer, c.LevelType); err != nil {
		return err
	}
	if err := codec.WriteBool(writer, c.ReducedDebugInfo); err != nil {
		return err
	}
	return nil
}

func (c *ClientJoinGame) Decode(reader io.Reader) error {
	var err error
	c.EntityID, err = codec.ReadInt(reader)
	if err != nil {
		return err
	}
	c.Gamemode, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	c.Dimension, err = codec.ReadByte(reader)
	if err != nil {
		return err
	}
	c.Difficulty, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	c.MaxPlayers, err = codec.ReadUByte(reader)
	if err != nil {
		return err
	}
	c.LevelType, err = codec.ReadString(reader)
	if err != nil {
		return err
	}
	c.ReducedDebugInfo, err = codec.ReadBool(reader)
	if err != nil {
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

func (c *ClientSetSpawnPosition) ID() int32 {
	return 0x05
}

func (c *ClientSetSpawnPosition) Encode(writer io.Writer) error {
	return codec.WritePosition(writer, c.X, c.Y, c.Z)
}

func (c *ClientSetSpawnPosition) Decode(reader io.Reader) error {
	var err error
	c.X, c.Y, c.Z, err = codec.ReadPosition(reader)
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

func (c *ClientPlayerPositionAndLook) ID() int32 {
	return 0x08
}

func (c *ClientPlayerPositionAndLook) Encode(writer io.Writer) error {
	if err := codec.WriteDouble(writer, c.X); err != nil {
		return err
	}
	if err := codec.WriteDouble(writer, c.Y); err != nil {
		return err
	}
	if err := codec.WriteDouble(writer, c.Z); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, c.Yaw); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, c.Pitch); err != nil {
		return err
	}
	if err := codec.WriteUByte(writer, c.Flags); err != nil {
		return err
	}
	return nil
}

func (c *ClientPlayerPositionAndLook) Decode(reader io.Reader) error {
	var err error
	if c.X, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if c.Y, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if c.Z, err = codec.ReadDouble(reader); err != nil {
		return err
	}
	if c.Yaw, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if c.Pitch, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if c.Flags, err = codec.ReadUByte(reader); err != nil {
		return err
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

func (c *ClientSpawnPlayer) ID() int32 {
	return 0x0C
}

func (c *ClientSpawnPlayer) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, c.EntityID); err != nil {
		return err
	}
	if err := codec.WriteUUID(writer, c.PlayerUUID); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, c.X); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, c.Y); err != nil {
		return err
	}
	if err := codec.WriteInt(writer, c.Z); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, c.Yaw); err != nil {
		return err
	}
	if err := codec.WriteAngle(writer, c.Pitch); err != nil {
		return err
	}
	if err := codec.WriteShort(writer, c.CurrentItem); err != nil {
		return err
	}
	if err := codec.WriteMetadata(writer, c.Metadata); err != nil {
		return err
	}
	return nil
}

func (c *ClientSpawnPlayer) Decode(reader io.Reader) error {
	var err error
	if c.EntityID, err = codec.ReadVarInt(reader); err != nil {
		return err
	}
	if c.PlayerUUID, err = codec.ReadUUID(reader); err != nil {
		return err
	}
	if c.X, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if c.Y, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if c.Z, err = codec.ReadInt(reader); err != nil {
		return err
	}
	if c.Yaw, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if c.Pitch, err = codec.ReadAngle(reader); err != nil {
		return err
	}
	if c.CurrentItem, err = codec.ReadShort(reader); err != nil {
		return err
	}
	if c.Metadata, err = codec.ReadMetadata(reader); err != nil {
		return err
	}
	return nil
}

type ClientPlayerAbilities struct {
	Flags               int8
	FlyingSpeed         float32
	FieldOfViewModifier float32
}

var _ proto.Packet = (*ClientPlayerAbilities)(nil)

func (c *ClientPlayerAbilities) ID() int32 {
	return 0x39
}

func (c *ClientPlayerAbilities) Encode(writer io.Writer) error {
	if err := codec.WriteByte(writer, c.Flags); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, c.FlyingSpeed); err != nil {
		return err
	}
	if err := codec.WriteFloat(writer, c.FieldOfViewModifier); err != nil {
		return err
	}
	return nil
}

func (c *ClientPlayerAbilities) Decode(reader io.Reader) error {
	var err error
	if c.Flags, err = codec.ReadByte(reader); err != nil {
		return err
	}
	if c.FlyingSpeed, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	if c.FieldOfViewModifier, err = codec.ReadFloat(reader); err != nil {
		return err
	}
	return nil
}
