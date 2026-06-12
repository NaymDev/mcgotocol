package codec

import (
	"bytes"
	"io"
	"math"
	"testing"

	mcio "github.com/NaymDev/mcgotocol/io"
	"github.com/NaymDev/mcgotocol/packet"
	"github.com/NaymDev/mcgotocol/profile"
	"github.com/NaymDev/mcgotocol/proto"
	"github.com/google/uuid"
)

func BenchmarkVarIntWrite(b *testing.B) {
	cases := []struct {
		name  string
		value VarInt
	}{
		{"varint(1)", 1},
		{"varint(1_000)", 1_000},
		{"varint(1_000_000)", 1_000_000},
		{"varint(maxint)", math.MaxInt32},
		{"varint(minint)", math.MinInt32},
	}

	benchEncoder(b, WriteVarInt, cases)
}

func BenchmarkVarLongWrite(b *testing.B) {
	cases := []struct {
		name  string
		value VarLong
	}{
		{"varlong(1)", 1},
		{"varlong(1_000)", 1_000},
		{"varlong(1_000_000)", 1_000_000},
		{"varlong(1_000_000_000)", 1_000_000_000},
		{"varlong(maxint)", math.MaxInt64},
		{"varlong(minint)", math.MinInt64},
	}
	benchEncoder(b, WriteVarLong, cases)
}

func BenchmarkBoolWrite(b *testing.B) {
	cases := []struct {
		name  string
		value bool
	}{
		{"bool(true)", true},
		{"bool(false)", false},
	}
	benchEncoder(b, WriteBool, cases)
}

func BenchmarkByteWrite(b *testing.B) {
	cases := []struct {
		name  string
		value int8
	}{
		{"byte(0)", 0x00},
		{"byte(1)", 0x01},
		{"byte(maxbyte)", math.MaxInt8},
		{"byte(minbyte)", math.MinInt8},
	}
	benchEncoder(b, WriteByte, cases)
}

func BenchmarkUByteWrite(b *testing.B) {
	cases := []struct {
		name  string
		value uint8
	}{
		{"ubyte(0)", 0x00},
		{"ubyte(1)", 0x01},
		{"ubyte(maxbyte)", math.MaxUint8},
	}
	benchEncoder(b, WriteUByte, cases)
}

func BenchmarkShortWrite(b *testing.B) {
	cases := []struct {
		name  string
		value int16
	}{
		{"short(0)", 0},
		{"short(100)", 100},
		{"short(10_000)", 10_000},
		{"short(maxint)", math.MaxInt16},
		{"short(minint)", math.MinInt16},
	}
	benchEncoder(b, WriteShort, cases)
}

func BenchmarkUShortWrite(b *testing.B) {
	cases := []struct {
		name  string
		value uint16
	}{
		{"ushort(0)", 0},
		{"ushort(100)", 100},
		{"ushort(10_000)", 10_000},
		{"ushort(maxint)", math.MaxInt16},
	}
	benchEncoder(b, WriteUShort, cases)
}

func BenchmarkIntWrite(b *testing.B) {
	cases := []struct {
		name  string
		value int32
	}{
		{"int(0)", 0},
		{"int(100)", 100},
		{"int(10_000)", 10_000},
		{"int(1_000_000)", 1_000_000},
		{"int(maxint)", math.MaxInt32},
		{"int(minint)", math.MinInt32},
	}
	benchEncoder(b, WriteInt, cases)
}

func BenchmarkLongWrite(b *testing.B) {
	cases := []struct {
		name  string
		value int64
	}{
		{"int(0)", 0},
		{"int(100)", 100},
		{"int(10_000)", 10_000},
		{"int(1_000_000)", 1_000_000},
		{"int(maxint)", math.MaxInt64},
		{"int(minint)", math.MinInt64},
	}
	benchEncoder(b, WriteLong, cases)
}

func BenchmarkFloatWrite(b *testing.B) {
	cases := []struct {
		name  string
		value float32
	}{
		{"float(0)", 0},
		{"float(1.5)", 1.5},
		{"float(-1.5)", -1.5},
		{"float(100)", 100},
		{"float(10_000)", 10_000},
		{"float(1_000_000)", 1_000_000},
		{"float(max)", math.MaxFloat32},
		{"float(min)", -math.MaxFloat32},
		{"float(+inf)", float32(math.Inf(1))},
		{"float(-inf)", float32(math.Inf(-1))},
		{"float(nan)", float32(math.NaN())},
	}

	benchEncoder(b, WriteFloat, cases)
}

func BenchmarkDoubleWrite(b *testing.B) {
	cases := []struct {
		name  string
		value float64
	}{
		{"double(0)", 0},
		{"double(1.5)", 1.5},
		{"double(-1.5)", -1.5},
		{"double(100)", 100},
		{"double(10_000)", 10_000},
		{"double(1_000_000)", 1_000_000},
		{"double(max)", math.MaxFloat64},
		{"double(min)", -math.MaxFloat64},
		{"double(+inf)", math.Inf(1)},
		{"double(-inf)", math.Inf(-1)},
		{"double(nan)", math.NaN()},
	}

	benchEncoder(b, WriteDouble, cases)
}

func BenchmarkPacketsWrite(b *testing.B) {
	cases := []struct {
		name  string
		value proto.Packet
	}{
		{
			name: "handshake status",
			value: &packet.ServerHandshake{
				ProtocolVersion: 760,
				ServerAddress:   "localhost",
				ServerPort:      25565,
				NextState:       packet.StatusHandshakeIntent,
			},
		},
		{
			name: "handshake login",
			value: &packet.ServerHandshake{
				ProtocolVersion: 761,
				ServerAddress:   "mc.example.com",
				ServerPort:      25565,
				NextState:       packet.LoginHandshakeIntent,
			},
		},
		{
			name: "login start",
			value: &packet.ServerLoginStart{
				Name: "Steve",
			},
		},
		{
			name: "encryption response",
			value: &packet.ServerEncryptionResponse{
				SharedSecret: []byte{0x01, 0x02, 0x03},
				VerifyToken:  []byte{0x0A, 0x0B},
			},
		},
		{
			name: "disconnect login",
			value: &packet.ClientDisconnect{
				Reason: `{"text":"Disconnected"}`,
			},
		},
		{
			name: "login success",
			value: &packet.ClientLoginSuccess{
				UUID:     "123e4567-e89b-12d3-a456-426614174000",
				Username: "Steve",
			},
		},
		{
			name: "set compression",
			value: &packet.ClientSetCompression{
				Threshold: 256,
			},
		},
		{
			name:  "status request",
			value: &packet.ServerStatusRequest{},
		},
		{
			name: "status response",
			value: &packet.ClientStatusResponse{
				JSONResponse: `{"version":{"name":"1.20.4","protocol":765},"players":{"max":100,"online":5},"description":{"text":"Hello world"}}`,
			},
		},
		{
			name: "status ping",
			value: &packet.ServerStatusPing{
				Payload: 9223372036854775807,
			},
		},
		{
			name: "status pong",
			value: &packet.ClientStatusPong{
				Payload: 42,
			},
		},
		{
			name: "keep alive server",
			value: &packet.ServerKeepAlive{
				KeepAliveID: 12345,
			},
		},
		{
			name: "keep alive client",
			value: &packet.ClientKeepAlive{
				KeepAliveID: 54321,
			},
		},
		{
			name: "join game",
			value: &packet.ClientJoinGame{
				EntityID:         1,
				Gamemode:         packet.GamemodeCreative,
				Dimension:        0,
				Difficulty:       1,
				MaxPlayers:       20,
				LevelType:        "default",
				ReducedDebugInfo: false,
			},
		},
		{
			name: "chat message",
			value: &packet.ClientChatMessage{
				JSONData: `{"text":"Hello world"}`,
				Position: 1,
			},
		},
		{
			name: "set spawn position",
			value: &packet.ClientSetSpawnPosition{
				X: 100,
				Y: 64,
				Z: -100,
			},
		},
		{
			name: "player position and look",
			value: &packet.ClientPlayerPositionAndLook{
				X:     10.5,
				Y:     64.0,
				Z:     -30.25,
				Yaw:   90.0,
				Pitch: 45.0,
				Flags: uint8(packet.X | packet.Y),
			},
		},
		{
			name: "spawn player",
			value: &packet.ClientSpawnPlayer{
				EntityID:    1,
				PlayerUUID:  uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
				X:           100,
				Y:           64,
				Z:           100,
				Yaw:         0,
				Pitch:       0,
				CurrentItem: 0,
				Metadata:    nil,
			},
		},
		{
			name: "player abilities",
			value: &packet.ClientPlayerAbilities{
				Flags:               0x01,
				FlyingSpeed:         0.05,
				FieldOfViewModifier: 1.0,
			},
		},
		{
			name: "player list add",
			value: &packet.ClientPlayerListItem{
				Action: packet.PlayerListItemActionAddPlayer,
				Players: []packet.PlayerListItemEntry{
					{
						UUID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						Name: func() *string { s := "Steve"; return &s }(),
						Properties: []packet.Property{
							packet.FromProfileProperty(profile.Property{
								Name:  "skin",
								Value: "default",
							}),
						},
						Gamemode: func() *VarInt { v := VarInt(packet.GamemodeCreative); return &v }(),
						Ping:     func() *VarInt { v := VarInt(50); return &v }(),
					},
				},
			},
		},
		{
			name: "player list update gamemode",
			value: &packet.ClientPlayerListItem{
				Action: packet.PlayerListItemActionUpdateGamemode,
				Players: []packet.PlayerListItemEntry{
					{
						UUID:     uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
						Gamemode: func() *VarInt { v := VarInt(packet.GamemodeSurvival); return &v }(),
					},
				},
			},
		},
		{
			name: "player list remove",
			value: &packet.ClientPlayerListItem{
				Action: packet.PlayerListItemActionRemovePlayer,
				Players: []packet.PlayerListItemEntry{
					{
						UUID: uuid.MustParse("123e4567-e89b-12d3-a456-426614174000"),
					},
				},
			},
		},
	}

	benchEncoder(b, mcio.WritePacket, cases)
}

func benchEncoder[T any](
	b *testing.B,
	encode func(w io.Writer, value T) error,
	cases []struct {
		name  string
		value T
	},
) {
	writers := []struct {
		name string
		new  func() io.Writer
	}{
		{"bytes.Buffer", func() io.Writer { return &bytes.Buffer{} }},
		{"noop", func() io.Writer { return discardWriter{} }},
	}

	for _, w := range writers {
		b.Run(w.name, func(b *testing.B) {
			for _, tc := range cases {
				b.Run(tc.name, func(b *testing.B) {
					b.ReportAllocs()

					for b.Loop() {
						w := w.new()

						_ = encode(w, tc.value)
					}
				})
			}
		})
	}
}

type discardWriter struct{}

func (d discardWriter) Write(p []byte) (int, error) {
	return len(p), nil
}
