package codec

import (
	"encoding/binary"
	"errors"
	"github.com/google/uuid"
	"io"
	"math"
)

// ============================
//  VarInt / VarLong
// ============================

type VarInt int32

func ReadVarInt(r io.Reader) (VarInt, error) {
	var num VarInt
	var shift uint
	for {
		var b [1]byte
		if _, err := r.Read(b[:]); err != nil {
			return 0, err
		}
		num |= VarInt(b[0]&0x7F) << shift

		if (b[0] & 0x80) == 0 {
			break
		}
		shift += 7
		if shift > 35 {
			return 0, errors.New("varint too long")
		}
	}
	return num, nil
}

func WriteVarInt(w io.Writer, value VarInt) error {
	u := uint32(value)
	for {
		b := byte(u & 0x7F)
		u >>= 7
		if u != 0 {
			b |= 0x80
		}
		if _, err := w.Write([]byte{b}); err != nil {
			return err
		}
		if u == 0 {
			break
		}
	}
	return nil
}

type VarLong int64

func ReadVarLong(r io.Reader) (VarLong, error) {
	var num VarLong
	var shift uint
	for {
		var b [1]byte
		if _, err := r.Read(b[:]); err != nil {
			return 0, err
		}
		num |= VarLong(b[0]&0x7F) << shift

		if (b[0] & 0x80) == 0 {
			break
		}
		shift += 7
		if shift > 70 {
			return 0, errors.New("varlong too long")
		}
	}
	return num, nil
}

// WriteVarLong TODO: prevent infinit loof when negative
func WriteVarLong(w io.Writer, value VarLong) error {
	for {
		b := byte(value & 0x7F)
		value >>= 7
		if value != 0 {
			b |= 0x80
		}
		if _, err := w.Write([]byte{b}); err != nil {
			return err
		}
		if value == 0 {
			break
		}
	}
	return nil
}

// ============================
//  Primitives (Big-Endian)
// ============================

func ReadBool(r io.Reader) (bool, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0] != 0, err
}

func WriteBool(w io.Writer, v bool) error {
	if v {
		_, err := w.Write([]byte{1})
		return err
	}
	_, err := w.Write([]byte{0})
	return err
}

func ReadByte(r io.Reader) (int8, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return int8(b[0]), err
}
func WriteByte(w io.Writer, v int8) error {
	_, err := w.Write([]byte{byte(v)})
	return err
}

func ReadUByte(r io.Reader) (uint8, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return b[0], err
}
func WriteUByte(w io.Writer, v uint8) error {
	_, err := w.Write([]byte{v})
	return err
}

func ReadShort(r io.Reader) (int16, error) {
	var val int16
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}
func WriteShort(w io.Writer, v int16) error {
	return binary.Write(w, binary.BigEndian, v)
}

func ReadUShort(r io.Reader) (uint16, error) {
	var val uint16
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}
func WriteUShort(w io.Writer, v uint16) error {
	return binary.Write(w, binary.BigEndian, v)
}

func ReadInt(r io.Reader) (int32, error) {
	var val int32
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}
func WriteInt(w io.Writer, v int32) error {
	return binary.Write(w, binary.BigEndian, v)
}

func ReadLong(r io.Reader) (int64, error) {
	var val int64
	err := binary.Read(r, binary.BigEndian, &val)
	return val, err
}
func WriteLong(w io.Writer, v int64) error {
	return binary.Write(w, binary.BigEndian, v)
}

func ReadFloat(r io.Reader) (float32, error) {
	var bits uint32
	err := binary.Read(r, binary.BigEndian, &bits)
	return math.Float32frombits(bits), err
}
func WriteFloat(w io.Writer, v float32) error {
	bits := math.Float32bits(v)
	return binary.Write(w, binary.BigEndian, bits)
}

func ReadDouble(r io.Reader) (float64, error) {
	var bits uint64
	err := binary.Read(r, binary.BigEndian, &bits)
	return math.Float64frombits(bits), err
}
func WriteDouble(w io.Writer, v float64) error {
	bits := math.Float64bits(v)
	return binary.Write(w, binary.BigEndian, bits)
}

// ============================
//  Strings
// ============================

func ReadString(r io.Reader) (string, error) {
	length, err := ReadVarInt(r)
	if err != nil {
		return "", err
	}
	buf := make([]byte, length)
	_, err = io.ReadFull(r, buf)
	return string(buf), err
}

func WriteString(w io.Writer, s string) error {
	if err := WriteVarInt(w, VarInt(len(s))); err != nil {
		return err
	}
	_, err := w.Write([]byte(s))
	return err
}

// ============================
//  Byte Arrays
// ============================

func ReadByteArray(r io.Reader) ([]byte, error) {
	length, err := ReadVarInt(r)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, length)
	_, err = io.ReadFull(r, buf)
	return buf, err
}

func WriteByteArray(w io.Writer, data []byte) error {
	if err := WriteVarInt(w, VarInt(len(data))); err != nil {
		return err
	}
	_, err := w.Write(data)
	return err
}

// ============================
//  UUIDs
// ============================

func ReadUUID(r io.Reader) (uuid.UUID, error) {
	return uuid.NewRandomFromReader(r)
}

func WriteUUID(w io.Writer, uuid uuid.UUID) error {
	_, err := w.Write(uuid[:])
	return err
}

type Angle uint8

func ReadAngle(r io.Reader) (Angle, error) {
	var b [1]byte
	_, err := r.Read(b[:])
	return Angle(b[0]), err
}
func WriteAngle(w io.Writer, v Angle) error {
	_, err := w.Write([]byte{byte(v)})
	return err
}
