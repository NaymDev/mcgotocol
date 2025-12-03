package states

import "fmt"

type State int

const (
	HandshakeState State = 0
	StatusState    State = 1
	LoginState     State = 2
	PlayState      State = 3
)

func (s State) String() string {
	switch s {
	case HandshakeState:
		return "Handshake"
	case StatusState:
		return "Status"
	case LoginState:
		return "Login"
	case PlayState:
		return "Play"
	}
	return "UnknownState"
}

var _ fmt.Stringer = (*State)(nil)
