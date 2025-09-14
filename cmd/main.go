package main

import (
	"fmt"

	msgpack "github.com/codescalersinternships/msgpack-rawan/pkg"
)

func main() {
	bytes, err := msgpack.Serialize(true)
	if err != nil {
		panic(err)
	}

	deserialized, err := msgpack.Deserialize(bytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Deserialized: %v\n", deserialized)
	fmt.Printf("Type: %T\n", deserialized)

}
