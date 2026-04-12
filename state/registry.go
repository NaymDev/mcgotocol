package state

import (
	"io"
	"reflect"

	"github.com/NaymDev/mcgotocol/packets"
	"github.com/NaymDev/mcgotocol/state/states"
)

const MaxPacketID = 0x49

type Registry struct {
	State       states.State
	ClientBound *PacketRegistry
	ServerBound *PacketRegistry
}

func NewRegistry(state states.State) *Registry {
	return &Registry{
		State: state,
		ClientBound: &PacketRegistry{
			State: state.String() + " ClientBound",
		},
		ServerBound: &PacketRegistry{
			State: state.String() + " ServerBound",
		},
	}
}

type PacketRegistry struct {
	State string
	ctors [MaxPacketID]Constructor
}
type Constructor func() packets.Packet

func (r *PacketRegistry) Register(packet packets.Packet) {
	id := packet.ID()
	t := reflect.TypeOf(packet)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	r.ctors[id] = func() packets.Packet {
		v := reflect.New(t)
		return v.Interface().(packets.Packet)
	}
}

func (r *PacketRegistry) Decode(id int32, reader io.Reader) (packets.Packet, error) {
	if id < 0 || int(id) >= MaxPacketID {
		data, _ := io.ReadAll(reader)
		return nil, &UnknownPacketID{
			PacketID: id,
			State:    r.State,
			Data:     data,
		}
	}
	ctor := r.ctors[id]
	if ctor == nil {
		data, _ := io.ReadAll(reader)
		return nil, &UnknownPacketID{
			PacketID: id,
			State:    r.State,
			Data:     data,
		}
	}

	pkt := ctor()
	if err := pkt.Decode(reader); err != nil {
		return nil, err
	}
	return pkt, nil
}
