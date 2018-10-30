package snappygob

import (
	"github.com/golang/snappy"
	"github.com/gorilla/securecookie"
)

// Encoder implements the securecookie.Serializer interface. It wraps the
// GobEncoder implementation with snappy compression.
type Encoder struct{}

// Serialize encodes a value using snappy + gob.
func (e Encoder) Serialize(src interface{}) ([]byte, error) {
	var enc securecookie.GobEncoder
	bytes, err := enc.Serialize(src)
	if err != nil {
		return nil, err
	}
	return snappy.Encode(nil, bytes), nil
}

// Deserialize decodes a value using snappy + gob.
func (e Encoder) Deserialize(src []byte, dst interface{}) error {
	uncompressedBytes, err := snappy.Decode(nil, src)
	if err != nil {
		return err
	}
	var enc securecookie.GobEncoder
	return enc.Deserialize(uncompressedBytes, dst)
}
