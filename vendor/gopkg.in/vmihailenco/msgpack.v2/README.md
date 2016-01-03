# MessagePack encoding for Golang [![Build Status](https://travis-ci.org/vmihailenco/msgpack.svg)](https://travis-ci.org/vmihailenco/msgpack)

Supports:
- Primitives, arrays, maps, structs, time.Time and interface{}.
- Appengine *datastore.Key and datastore.Cursor.
- [CustomEncoder](http://godoc.org/gopkg.in/vmihailenco/msgpack.v2#example-CustomEncoder)/CustomDecoder interfaces for custom encoding.
- [Extensions](http://godoc.org/gopkg.in/vmihailenco/msgpack.v2#example-RegisterExt) to encode type information.
- Fields renaming, e.g. `msgpack:"my_field_name"`.
- Fields inlining, e.g. `msgpack:",inline"`.
- Omitempty flag, e.g. `msgpack:",omitempty"`.

API docs: http://godoc.org/gopkg.in/vmihailenco/msgpack.v2.
Examples: http://godoc.org/gopkg.in/vmihailenco/msgpack.v2#pkg-examples.

## Installation

Install:

    go get gopkg.in/vmihailenco/msgpack.v2

## Quickstart

```go
func ExampleMarshal() {
	b, err := msgpack.Marshal(true)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)
	// Output:

	var out bool
	err = msgpack.Unmarshal([]byte{0xc3}, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
	// Output: []byte{0xc3}
	// true
}
```

## Howto

Please go through [examples](http://godoc.org/gopkg.in/vmihailenco/msgpack.v2#pkg-examples) to get an idea how to use this package.
