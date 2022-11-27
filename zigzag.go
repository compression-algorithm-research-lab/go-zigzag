package zigzag

import (
	"github.com/compression-algorithm-research-lab/go-varint"
	"github.com/golang-infrastructure/go-gtypes"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// ToZigZag 把整数转为zigzag格式
func ToZigZag[T gtypes.Integer](value T) T {
	return (value << 1) ^ (value >> 31)
}

// FromToZigZag 从zigzag格式恢复整数
func FromToZigZag[T gtypes.Integer](value T) T {
	return (value >> 1) ^ -(value & 1)
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Encode 对整数进行zigzag编码为字节数组
func Encode[T gtypes.Integer](value T) []byte {
	var r uint64 = uint64(ToZigZag(value))
	return varint.Encode(r)
}

// EncodeSlice 从zigzag编码的字节数组恢复整数
func EncodeSlice[T gtypes.Integer](value []T) []byte {
	result := make([]byte, 0)
	for _, x := range value {
		result = append(result, Encode(x)...)
	}
	return result
}

func Decode[T gtypes.Integer](bytes []byte) T {
	value := varint.Decode[uint64](bytes)
	return T(FromToZigZag(value))
}

func DecodeSlice[T gtypes.Integer](bytes []byte) []T {
	uintSlice := varint.DecodeSlice[uint64](bytes)
	result := make([]T, 0)
	for _, x := range uintSlice {
		result = append(result, T(FromToZigZag(x)))
	}
	return result
}

// ------------------------------------------------ ---------------------------------------------------------------------
