package state

import (
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/NaymDev/mcgotocol/state/states"
	"io"
	"reflect"
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
type Constructor func() proto.Packet

func (r *PacketRegistry) Register(packet proto.Packet) {
	id := packet.ID()
	t := reflect.TypeOf(packet)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	r.ctors[id] = func() proto.Packet {
		v := reflect.New(t)
		return v.Interface().(proto.Packet)
	}
}

func (r *PacketRegistry) Decode(id int32, reader io.Reader) (proto.Packet, error) {
	if id < 0 || int(id) >= MaxPacketID {
		return nil, &UnknownPacketID{
			PacketID: id,
			State:    r.State,
		}
	}
	ctor := r.ctors[id]
	if ctor == nil {
		return nil, &UnknownPacketID{
			PacketID: id,
			State:    r.State,
		}
	}

	pkt := ctor()
	if err := pkt.Decode(reader); err != nil {
		return nil, err
	}
	return pkt, nil
}
