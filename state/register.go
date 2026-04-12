package state

import (
	"github.com/NaymDev/mcgotocol/packets"
	"github.com/NaymDev/mcgotocol/state/states"
)

var (
	Handshake = NewRegistry(states.HandshakeState)
	Status    = NewRegistry(states.StatusState)
	Login     = NewRegistry(states.LoginState)
	Play      = NewRegistry(states.PlayState)
)

func InitRegistries() {
	// HANDSHAKE
	Handshake.ServerBound.Register(&packets.ServerHandshake{})

	// STATUS
	Status.ServerBound.Register(&packets.ServerStatusRequest{})
	Status.ServerBound.Register(&packets.ServerStatusPing{})

	Status.ClientBound.Register(&packets.ClientStatusResponse{})
	Status.ClientBound.Register(&packets.ClientStatusPong{})

	// LOGIN
	Login.ServerBound.Register(&packets.ServerLoginStart{})
	Login.ServerBound.Register(&packets.ServerEncryptionResponse{})

	Login.ClientBound.Register(&packets.ClientDisconnect{})
	Login.ClientBound.Register(&packets.ClientEncryptionRequest{})
	Login.ClientBound.Register(&packets.ClientLoginSuccess{})
	Login.ClientBound.Register(&packets.ClientSetCompression{})

	// PLAY
	Play.ServerBound.Register(&packets.ServerKeepAlive{})

	Play.ClientBound.Register(&packets.ClientKeepAlive{})
	Play.ClientBound.Register(&packets.ClientJoinGame{})
	Play.ClientBound.Register(&packets.ClientChatMessage{})
	Play.ClientBound.Register(&packets.ClientSetSpawnPosition{})
	Play.ClientBound.Register(&packets.ClientPlayerPositionAndLook{})
	Play.ClientBound.Register(&packets.ClientPlayerListItem{})
	Play.ClientBound.Register(&packets.ClientSpawnPlayer{})
	Play.ClientBound.Register(&packets.ClientPlayerAbilities{})
}
