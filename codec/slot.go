package codec

import "io"

type ItemSlot struct {
	ItemID  VarInt
	Count   int8
	NBTData []byte
}

func ReadSlot(r io.Reader) (ItemSlot, error) {
	var slot ItemSlot
	id, err := ReadVarInt(r)
	if err != nil {
		return slot, err
	}
	slot.ItemID = id

	if id < 0 {
		return slot, nil
	}

	count, err := ReadByte(r)
	if err != nil {
		return slot, err
	}
	slot.Count = count

	nbt, err := ReadByteArray(r)
	if err != nil {
		return slot, err
	}
	if len(nbt) == 0xFF {
		slot.NBTData = nil
	} else {
		slot.NBTData = nbt
	}
	return slot, nil
}

func WriteSlot(w io.Writer, slot ItemSlot) error {
	if err := WriteVarInt(w, slot.ItemID); err != nil {
		return err
	}
	if slot.ItemID < 0 {
		return nil
	}
	if err := WriteByte(w, slot.Count); err != nil {
		return err
	}
	if slot.NBTData == nil {
		if err := WriteVarInt(w, -1); err != nil {
			return err
		}
	} else {
		if err := WriteByteArray(w, slot.NBTData); err != nil {
			return err
		}
	}
	return nil
}
