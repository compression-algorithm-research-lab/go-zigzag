package zigzag

import (
	"testing"
)

func TestFromToZigZag(t *testing.T) {

}

func TestToZigZag(t *testing.T) {
	zag := ToZigZag(-1)
	t.Log(zag)
}

func TestDecodeSlice(t *testing.T) {

	valueSlice := []int{
		1, -1, -10000, 10086,
	}

	bytes := EncodeSlice(valueSlice)
	t.Log(bytes)

	decodeSlice := DecodeSlice[int](bytes)
	t.Log(decodeSlice)

}
