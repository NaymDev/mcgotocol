package state

import (
	"github.com/NaymDev/mcgotocol/packet"
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
	Handshake.ServerBound.Register(&packet.ServerHandshake{})

	// STATUS
	Status.ServerBound.Register(&packet.ServerStatusRequest{})
	Status.ServerBound.Register(&packet.ServerStatusPing{})

	Status.ClientBound.Register(&packet.ClientStatusResponse{})
	Status.ClientBound.Register(&packet.ClientStatusPong{})

	// LOGIN
	Login.ServerBound.Register(&packet.ServerLoginStart{})
	Login.ServerBound.Register(&packet.ServerEncryptionResponse{})

	Login.ClientBound.Register(&packet.ClientDisconnect{})
	Login.ClientBound.Register(&packet.ClientEncryptionRequest{})
	Login.ClientBound.Register(&packet.ClientLoginSuccess{})
	Login.ClientBound.Register(&packet.ClientSetCompression{})

	// PLAY
	Play.ServerBound.Register(&packet.ServerKeepAlive{})

	Play.ClientBound.Register(&packet.ClientKeepAlive{})
	Play.ClientBound.Register(&packet.ClientJoinGame{})
	Play.ClientBound.Register(&packet.ClientChatMessage{})
	Play.ClientBound.Register(&packet.ClientSetSpawnPosition{})
	Play.ClientBound.Register(&packet.ClientPlayerPositionAndLook{})
	Play.ClientBound.Register(&packet.ClientPlayerListItem{})
	Play.ClientBound.Register(&packet.ClientSpawnPlayer{})
	Play.ClientBound.Register(&packet.ClientPlayerAbilities{})
	Play.ClientBound.Register(&packet.ClientPlayerListHeaderAndFooter{})
}
