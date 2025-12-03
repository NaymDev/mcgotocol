package codec

import (
	"encoding/binary"
	"io"
)

func WritePosition(w io.Writer, x, y, z int32) error {
	ux := uint64(x & 0x3FFFFFF)
	uy := uint64(y & 0xFFF)
	uz := uint64(z & 0x3FFFFFF)

	val := (ux << 38) | (uy << 26) | uz
	return binary.Write(w, binary.BigEndian, val)
}

func ReadPosition(r io.Reader) (x, y, z int32, err error) {
	var val uint64
	if err = binary.Read(r, binary.BigEndian, &val); err != nil {
		return
	}

	x = int32(val >> 38)
	if x >= 1<<25 {
		x -= 1 << 26
	}

	y = int32((val >> 26) & 0xFFF)
	if y >= 1<<11 {
		y -= 1 << 12
	}

	z = int32(val & 0x3FFFFFF)
	if z >= 1<<25 {
		z -= 1 << 26
	}

	return
}
