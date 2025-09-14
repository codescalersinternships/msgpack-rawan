# Msgpack Serializer/Deserializer 

This repository implements the Msgpack object serialization/deserialization

**Serialization** is conversion from application objects into MessagePack formats via MessagePack type system.

**Deserialization** is conversion from MessagePack formats into application objects via MessagePack type system.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)


## Installation

1. Clone the repository

   ```bash
   git clone https://github.com/codescalersinternships/msgpack-rawan.git
   ```
## APIs
- `Serialize(object interface{}) ([]byte, error)` : Represents the serialization API, given the object itself, it converts it to its serialized bytes

- `Deserialize(bytes []byte) (interface{}, error)` : Represents the deserialization API, given the serialize bytes, it turns them back to objects

## Usage

```go
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
```