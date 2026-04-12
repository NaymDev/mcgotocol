package codec

import (
	"fmt"
	"io"
)

type MetadataType uint8

const (
	MetaByte MetadataType = iota
	MetaShort
	MetaInt
	MetaFloat
	MetaString
	MetaSlot
	MetaVector3I
	MetaVector3F
)

type EntityMetadata struct {
	Index uint8
	Type  MetadataType
	Value interface{}
}

func ReadMetadata(r io.Reader) ([]EntityMetadata, error) {
	var result []EntityMetadata
	for {
		item := make([]uint8, 1)
		if _, err := r.Read(item); err != nil {
			return nil, err
		}
		if item[0] == 0x7F {
			break
		}
		metaIndex := item[0] & 0x1F
		metaType := item[0] >> 5

		entry := EntityMetadata{
			Index: metaIndex,
			Type:  MetadataType(metaType),
		}

		var err error
		switch entry.Type {
		case MetaByte:
			entry.Value, err = ReadByte(r)
		case MetaShort:
			entry.Value, err = ReadShort(r)
		case MetaInt:
			entry.Value, err = ReadInt(r)
		case MetaFloat:
			entry.Value, err = ReadFloat(r)
		case MetaString:
			entry.Value, err = ReadString(r)
		case MetaSlot:
			entry.Value, err = ReadSlot(r)
		case MetaVector3I:
			x, err := ReadInt(r)
			if err != nil {
				return nil, err
			}
			y, err := ReadInt(r)
			if err != nil {
				return nil, err
			}
			z, err := ReadInt(r)
			if err != nil {
				return nil, err
			}
			entry.Value = [3]int32{x, y, z}
		case MetaVector3F:
			x, err := ReadFloat(r)
			if err != nil {
				return nil, err
			}
			y, err := ReadFloat(r)
			if err != nil {
				return nil, err
			}
			z, err := ReadFloat(r)
			if err != nil {
				return nil, err
			}
			entry.Value = [3]float32{x, y, z}
		default:
			return nil, fmt.Errorf("unknown metadata type %d", entry.Type)
		}
		if err != nil {
			return nil, err
		}
		result = append(result, entry)
	}
	return result, nil
}

func WriteMetadata(w io.Writer, metadata []EntityMetadata) error {
	for _, entry := range metadata {
		header := (uint8(entry.Index) & 0x1F) | (uint8(entry.Type) << 5)

		if _, err := w.Write([]byte{header}); err != nil {
			return err
		}

		var err error

		switch entry.Type {
		case MetaByte:
			v, ok := entry.Value.(int8)
			if !ok {
				return fmt.Errorf("invalid MetaByte value type")
			}
			err = WriteByte(w, v)

		case MetaShort:
			v, ok := entry.Value.(int16)
			if !ok {
				return fmt.Errorf("invalid MetaShort value type")
			}
			err = WriteShort(w, v)

		case MetaInt:
			v, ok := entry.Value.(int32)
			if !ok {
				return fmt.Errorf("invalid MetaInt value type")
			}
			err = WriteInt(w, v)

		case MetaFloat:
			v, ok := entry.Value.(float32)
			if !ok {
				return fmt.Errorf("invalid MetaFloat value type")
			}
			err = WriteFloat(w, v)

		case MetaString:
			v, ok := entry.Value.(string)
			if !ok {
				return fmt.Errorf("invalid MetaString value type")
			}
			err = WriteString(w, v)

		case MetaSlot:
			v, ok := entry.Value.(ItemSlot)
			if !ok {
				return fmt.Errorf("invalid MetaSlot value type")
			}
			err = WriteSlot(w, v)

		case MetaVector3I:
			v, ok := entry.Value.([3]int32)
			if !ok {
				return fmt.Errorf("invalid MetaVector3I value type")
			}
			if err = WriteInt(w, v[0]); err != nil {
				return err
			}
			if err = WriteInt(w, v[1]); err != nil {
				return err
			}
			if err = WriteInt(w, v[2]); err != nil {
				return err
			}

		case MetaVector3F:
			v, ok := entry.Value.([3]float32)
			if !ok {
				return fmt.Errorf("invalid MetaVector3F value type")
			}
			if err = WriteFloat(w, v[0]); err != nil {
				return err
			}
			if err = WriteFloat(w, v[1]); err != nil {
				return err
			}
			if err = WriteFloat(w, v[2]); err != nil {
				return err
			}

		default:
			return fmt.Errorf("unknown metadata type %d", entry.Type)
		}

		if err != nil {
			return err
		}
	}

	_, err := w.Write([]byte{0x7F})
	return err
}
