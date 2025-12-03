package codec

import (
	"fmt"
	"io"
)

type MetadataType byte

const (
	MetaByte MetadataType = iota
	MetaVarInt
	MetaFloat
	MetaString
	MetaSlot
	MetaBool
	MetaVector3F
)

type EntityMetadata struct {
	Index byte
	Type  MetadataType
	Value interface{}
}

func ReadMetadata(r io.Reader) ([]EntityMetadata, error) {
	var result []EntityMetadata
	for {
		indexByte := make([]byte, 1)
		if _, err := r.Read(indexByte); err != nil {
			return nil, err
		}
		if indexByte[0] == 0xFF {
			break
		}
		metaTypeByte := make([]byte, 1)
		if _, err := r.Read(metaTypeByte); err != nil {
			return nil, err
		}
		entry := EntityMetadata{
			Index: indexByte[0],
			Type:  MetadataType(metaTypeByte[0]),
		}

		var err error
		switch entry.Type {
		case MetaByte:
			entry.Value, err = ReadByte(r)
		case MetaVarInt:
			entry.Value, err = ReadVarInt(r)
		case MetaFloat:
			entry.Value, err = ReadFloat(r)
		case MetaString:
			entry.Value, err = ReadString(r)
		case MetaSlot:
			entry.Value, err = ReadSlot(r)
		case MetaBool:
			entry.Value, err = ReadBool(r)
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
		if _, err := w.Write([]byte{entry.Index}); err != nil {
			return err
		}
		if _, err := w.Write([]byte{byte(entry.Type)}); err != nil {
			return err
		}

		switch entry.Type {
		case MetaByte:
			if v, ok := entry.Value.(int8); ok {
				if err := WriteByte(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaByte")
			}

		case MetaVarInt:
			if v, ok := entry.Value.(VarInt); ok {
				if err := WriteVarInt(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaVarInt")
			}

		case MetaFloat:
			if v, ok := entry.Value.(float32); ok {
				if err := WriteFloat(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaFloat")
			}

		case MetaString:
			if v, ok := entry.Value.(string); ok {
				if err := WriteString(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaString")
			}

		case MetaSlot:
			if v, ok := entry.Value.(ItemSlot); ok {
				if err := WriteSlot(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaSlot")
			}

		case MetaBool:
			if v, ok := entry.Value.(bool); ok {
				if err := WriteBool(w, v); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaBool")
			}

		case MetaVector3F:
			if v, ok := entry.Value.([3]float32); ok {
				if err := WriteFloat(w, v[0]); err != nil {
					return err
				}
				if err := WriteFloat(w, v[1]); err != nil {
					return err
				}
				if err := WriteFloat(w, v[2]); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("invalid value type for MetaVector3F")
			}

		default:
			return fmt.Errorf("unknown metadata type %d", entry.Type)
		}
	}

	_, err := w.Write([]byte{0xFF})
	return err
}
