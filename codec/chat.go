package codec

import (
	"errors"
	"io"
)

const MaxChatLength = 32767

var ErrStringTooLong = errors.New("string exceeds maximum length")

type Chat string

func ReadChat(r io.Reader) (Chat, error) {
	s, err := ReadString(r)
	if err != nil {
		return "", err
	}
	if len(s) > MaxChatLength {
		return "", ErrStringTooLong
	}
	return Chat(s), nil
}

func WriteChat(w io.Writer, c Chat) error {
	if len(c) > MaxChatLength {
		return ErrStringTooLong
	}
	return WriteString(w, string(c))
}
