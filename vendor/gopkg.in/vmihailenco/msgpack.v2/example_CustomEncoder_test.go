package msgpack_test

import (
	"fmt"

	"gopkg.in/vmihailenco/msgpack.v2"
)

type customStruct struct {
	S string
	N int
}

var (
	_ msgpack.CustomEncoder = &customStruct{}
	_ msgpack.CustomDecoder = &customStruct{}
)

func (s *customStruct) EncodeMsgpack(enc *msgpack.Encoder) error {
	return enc.Encode(s.S, s.N)
}

func (s *customStruct) DecodeMsgpack(dec *msgpack.Decoder) error {
	return dec.Decode(&s.S, &s.N)
}

func ExampleCustomEncoder() {
	b, err := msgpack.Marshal(&customStruct{S: "hello", N: 42})
	if err != nil {
		panic(err)
	}

	var v customStruct
	err = msgpack.Unmarshal(b, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", v)
	// Output: msgpack_test.customStruct{S:"hello", N:42}
}
