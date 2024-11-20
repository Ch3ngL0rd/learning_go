package main

import (
	"fmt"
	"math"
)

func main() {
	const MaxByte = math.MaxUint8
	const MaxInt32 = 1 << 31 - 1
	const MaxUInt64 = 1 << 64 - 1

	var b byte = MaxByte
	var smallI int32 = MaxInt32
	var bigI uint64 = MaxUInt64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b, smallI, bigI)
}