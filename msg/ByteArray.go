package msg

import (
	"encoding/binary"
	"errors"
)

// 缓冲区初始长度
var initSize = 0

type ByteArray struct {
	Bytes                 []byte
	ReadIndex, WriteIndex int
	Capacity              int
	InitSize              int
}

func (b *ByteArray) Remain() int {
	return b.Capacity - b.WriteIndex
}
func (b *ByteArray) Length() int {
	return b.WriteIndex - b.ReadIndex
}

func (b *ByteArray) Resize(size int) {
	if size < b.Length() {
		return
	}
	if size < initSize {
		return
	}
	n := 1
	for n < size {
		n *= 2
	}
	b.Capacity = n
	// Array.Copy(bytes,readIdx,newBytes,0,writeIdx-readIdx)----C#
	newBytes := make([]byte, 0, b.Capacity)
	readIdx := b.ReadIndex
	writeIdx := b.WriteIndex

	copy(newBytes, b.Bytes[readIdx:writeIdx])

	b.Bytes = newBytes
	writeIdx = b.Length()
	readIdx = 0

}

func (b *ByteArray) CheckAndMoveBytes() {
	if b.Length() < 8 {
		b.MoveBytes()
	}
}

func (b *ByteArray) MoveBytes() {
	//b.Bytes = b.Bytes[b.ReadIndex : b.Length()-b.ReadIndex]
	copy(b.Bytes, b.Bytes[b.ReadIndex:b.WriteIndex])
	//b.WriteIndex = b.Length()
	b.WriteIndex -= b.ReadIndex
	b.ReadIndex = 0
}

func (b *ByteArray) ReadInt16() (int16, error) {
	if b.Length() < 2 {
		return 0, errors.New("read int16 fail, buff length < 2")
	}
	// 小端
	ret := int16(binary.LittleEndian.Uint16(b.Bytes[b.ReadIndex:]))
	b.ReadIndex += 2
	b.CheckAndMoveBytes()
	return ret, nil
}
