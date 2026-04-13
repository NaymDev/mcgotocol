package packet

import (
	"io"

	"github.com/NaymDev/mcgotocol/codec"
	"github.com/NaymDev/mcgotocol/profile"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/google/uuid"
)

type PlayerListAction int32

const (
	PlayerListItemActionAddPlayer PlayerListAction = iota
	PlayerListItemActionUpdateGamemode
	PlayerListItemActionUpdateLatency
	PlayerListItemActionUpdateDisplayName
	PlayerListItemActionRemovePlayer
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

type PlayerListItemEntry struct {
	UUID        uuid.UUID
	Name        *string
	Properties  []Property
	Gamemode    *codec.VarInt
	Ping        *codec.VarInt
	DisplayName *codec.Chat
}

type ClientPlayerListItem struct {
	Action  PlayerListAction
	Players []PlayerListItemEntry
}

var _ proto.Packet = (*ClientPlayerListItem)(nil)

func (p *ClientPlayerListItem) ID() int32 {
	return 0x38
}

func (p *ClientPlayerListItem) Encode(writer io.Writer) error {
	if err := codec.WriteVarInt(writer, codec.VarInt(p.Action)); err != nil {
		return err
	}
	if err := codec.WriteVarInt(writer, codec.VarInt(len(p.Players))); err != nil {
		return err
	}
	for _, player := range p.Players {
		if err := codec.WriteUUID(writer, player.UUID); err != nil {
			return err
		}
		switch p.Action {
		case PlayerListItemActionAddPlayer:
			if err := codec.WriteString(writer, *player.Name); err != nil {
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
			if err := codec.WriteVarInt(writer, *player.Gamemode); err != nil {
				return err
			}
			if err := codec.WriteVarInt(writer, *player.Ping); err != nil {
				return err
			}
			if err := codec.WriteBool(writer, player.DisplayName != nil); err != nil {
				return err
			}
			if player.DisplayName != nil {
				if err := codec.WriteChat(writer, *player.DisplayName); err != nil {
					return err
				}
			}
		case PlayerListItemActionUpdateGamemode:
			if err := codec.WriteVarInt(writer, *player.Gamemode); err != nil {
				return err
			}
		case PlayerListItemActionUpdateLatency:
			if err := codec.WriteVarInt(writer, *player.Ping); err != nil {
				return err
			}
		case PlayerListItemActionUpdateDisplayName:
			if err := codec.WriteBool(writer, player.DisplayName != nil); err != nil {
				return err
			}
			if player.DisplayName != nil {
				if err := codec.WriteChat(writer, *player.DisplayName); err != nil {
					return err
				}
			}
		case PlayerListItemActionRemovePlayer:
		}
	}
	return nil
}

func (p *ClientPlayerListItem) Decode(reader io.Reader) error {
	actionInt, err := codec.ReadVarInt(reader)
	if err != nil {
		return err
	}
	p.Action = PlayerListAction(actionInt)

	playerCount, err := codec.ReadVarInt(reader)
	if err != nil {
		return err
	}

	p.Players = make([]PlayerListItemEntry, playerCount)

	for i := 0; i < int(playerCount); i++ {
		player := PlayerListItemEntry{}

		u, err := codec.ReadUUID(reader)
		if err != nil {
			return err
		}
		player.UUID = u

		switch p.Action {
		case PlayerListItemActionAddPlayer:
			name, err := codec.ReadString(reader)
			if err != nil {
				return err
			}
			player.Name = &name

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

			gamemode, err := codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
			player.Gamemode = &gamemode

			ping, err := codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
			player.Ping = &ping

			hasDisplayName, err := codec.ReadBool(reader)
			if err != nil {
				return err
			}
			if hasDisplayName {
				displayName, err := codec.ReadChat(reader)
				if err != nil {
					return err
				}
				player.DisplayName = &displayName
			}

		case PlayerListItemActionUpdateGamemode:
			gamemode, err := codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
			player.Gamemode = &gamemode
		case PlayerListItemActionUpdateLatency:
			ping, err := codec.ReadVarInt(reader)
			if err != nil {
				return err
			}
			player.Ping = &ping
		case PlayerListItemActionUpdateDisplayName:
			hasDisplayName, err := codec.ReadBool(reader)
			if err != nil {
				return err
			}
			if hasDisplayName {
				displayName, err := codec.ReadChat(reader)
				if err != nil {
					return err
				}
				player.DisplayName = &displayName
			}
		case PlayerListItemActionRemovePlayer:
		}

		p.Players[i] = player
	}

	return nil
}
