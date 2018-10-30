package snappygob_test

import (
	"encoding/gob"
	"testing"

	"github.com/eklitzke/snappygob"
)

// test type to encode
type myTestType struct {
	Foo int
	Bar string
}

// register the type with gob
func init() {
	gob.Register(myTestType{})
}

// test a round trip Serialize -> Deserialize
func TestSerializeRoundTrip(t *testing.T) {
	testVal := myTestType{Foo: 3, Bar: "hello world"}
	enc := snappygob.Encoder{}
	bytes, err := enc.Serialize(testVal)
	if err != nil {
		t.Error(err)
	}

	var roundTripVal myTestType
	err = enc.Deserialize(bytes, &roundTripVal)
	if err != nil {
		t.Error(err)
	}

	if testVal != roundTripVal {
		t.Errorf("orig val = %v, round trip val = %v", testVal, roundTripVal)
	}
}
