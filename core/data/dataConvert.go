package data

import (
	"bytes"
	"encoding/binary"
)

func BytesToInt32(bs []byte) int32 {
	var i int32
	buff := bytes.NewBuffer(bs)
	binary.Read(buff, binary.BigEndian, &i)
	return i
}

func BytesToUint8(bs []byte) uint8 {
	var i uint8
	buff := bytes.NewBuffer(bs)
	binary.Read(buff, binary.BigEndian, &i)
	return i
}

func Int32ToBytes(i int32) []byte {
	bs := make([]byte, 0, 4)
	buff := bytes.NewBuffer(bs)
	binary.Write(buff, binary.BigEndian, &i)
	return buff.Bytes()
}

func Uint8ToBytes(i uint8) []byte {
	bs := make([]byte, 0, 1)
	buff := bytes.NewBuffer(bs)
	binary.Write(buff, binary.BigEndian, &i)
	return buff.Bytes()
}
