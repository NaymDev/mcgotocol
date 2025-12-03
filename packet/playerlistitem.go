package packet

import (
	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/profile"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/google/uuid"
	"io"
)

type PlayerListAction int32

const (
	AddPlayer PlayerListAction = iota
	UpdateGamemode
	UpdateLatency
	UpdateDisplayName
	RemovePlayer
)

type Property struct {
	profile.Property
	IsSigned  bool
	Signature string
}

func FromProfileProperty(p profile.Property) Property {
	var sign string
	if p.Signature != nil {
		sign = *p.Signature
	} else {
		sign = ""
	}
	return Property{
		Property:  p,
		IsSigned:  p.Signature != nil,
		Signature: sign,
	}
}

func (p *Property) Encode(w io.Writer) error {
	if err := codec.WriteString(w, p.Name); err != nil {
		return err
	}
	if err := codec.WriteString(w, p.Value); err != nil {
		return err
	}
	if err := codec.WriteBool(w, p.IsSigned); err != nil {
		return err
	}
	if p.IsSigned {
		if err := codec.WriteString(w, p.Signature); err != nil {
			return err
		}
	}
	return nil
}

type PlayerProfile struct {
	UUID           uuid.UUID
	Name           string
	Properties     []Property
	Gamemode       codec.VarInt
	Ping           codec.VarInt
	HasDisplayName bool
	DisplayName    *codec.Chat
}

type ClientPlayerListItem struct {
	Action  PlayerListAction
	Players []PlayerProfile
}

var _ proto.Packet = (*ClientPlayerListItem)(nil)

func (c *ClientPlayerListItem) ID() int32 {
	return 0x38
}

func (c *ClientPlayerListItem) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, codec.VarInt(c.Action)); err != nil {
		return err
	}
	if err := codec.WriteVarInt(writer, codec.VarInt(len(c.Players))); err != nil {
		return err
	}
	for _, player := range c.Players {
		if err := codec.WriteUUID(writer, player.UUID); err != nil {
			return err
		}
		switch c.Action {
		case AddPlayer:
			if err := codec.WriteString(writer, player.Name); err != nil {
				return err
			}
			if err := codec.WriteVarInt(writer, codec.VarInt(len(player.Properties))); err != nil {
				return err
			}
			for _, prop := range player.Properties {
				if err := prop.Encode(writer); err != nil {
					return err
				}
			}
			if err := codec.WriteVarInt(writer, player.Gamemode); err != nil {
				return err
			}
			if err := codec.WriteVarInt(writer, player.Ping); err != nil {
				return err
			}
			if err := codec.WriteBool(writer, player.HasDisplayName); err != nil {
				return err
			}
			if player.HasDisplayName {
				if err := codec.WriteChat(writer, *player.DisplayName); err != nil {
					return err
				}
			}
		case UpdateGamemode:
			if err := codec.WriteVarInt(writer, player.Gamemode); err != nil {
				return err
			}
		case UpdateLatency:
			if err := codec.WriteVarInt(writer, player.Ping); err != nil {
				return err
			}
		case UpdateDisplayName:
			if err := codec.WriteBool(writer, player.HasDisplayName); err != nil {
				return err
			}
			if player.HasDisplayName {
				if err := codec.WriteChat(writer, *player.DisplayName); err != nil {
					return err
				}
			}
		case RemovePlayer:
		}
	}
	return nil
}

func (c *ClientPlayerListItem) Decode(reader io.Reader) error {
	actionInt, err := codec.ReadVarInt(reader)
	if err != nil {
		return err
	}
	c.Action = PlayerListAction(actionInt)

	playerCount, err := codec.ReadVarInt(reader)
	if err != nil {
		return err
	}

	c.Players = make([]PlayerProfile, playerCount)

	for i := 0; i < int(playerCount); i++ {
		player := PlayerProfile{}

		u, err := codec.ReadUUID(reader)
		if err != nil {
			return err
		}
		player.UUID = u

		switch c.Action {
		case AddPlayer:
			name, err := codec.ReadString(reader)
			if err != nil {
				return err
			}
			player.Name = name

			propCount, err := codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
			player.Properties = make([]Property, propCount)
			for j := 0; j < int(propCount); j++ {
				prop := Property{}
				prop.Name, err = codec.ReadString(reader)
				if err != nil {
					return err
				}
				prop.Value, err = codec.ReadString(reader)
				if err != nil {
					return err
				}
				prop.IsSigned, err = codec.ReadBool(reader)
				if err != nil {
					return err
				}
				if prop.IsSigned {
					prop.Signature, err = codec.ReadString(reader)
					if err != nil {
						return err
					}
				}
				player.Properties[j] = prop
			}

			player.Gamemode, err = codec.ReadVarInt(reader)
			if err != nil {
				return err
			}

			player.Ping, err = codec.ReadVarInt(reader)
			if err != nil {
				return err
			}

			player.HasDisplayName, err = codec.ReadBool(reader)
			if err != nil {
				return err
			}
			if player.HasDisplayName {
				displayName, err := codec.ReadChat(reader)
				if err != nil {
					return err
				}
				player.DisplayName = &displayName
			}

		case UpdateGamemode:
			player.Gamemode, err = codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
		case UpdateLatency:
			player.Ping, err = codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
		case UpdateDisplayName:
			player.HasDisplayName, err = codec.ReadBool(reader)
			if err != nil {
				return err
			}
			if player.HasDisplayName {
				displayName, err := codec.ReadChat(reader)
				if err != nil {
					return err
				}
				player.DisplayName = &displayName
			}
		case RemovePlayer:
		}

		c.Players[i] = player
	}

	return nil
}
