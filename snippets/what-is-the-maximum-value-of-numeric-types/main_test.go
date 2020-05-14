// Numbers: What is the maximum value of numeric types
package main

import (
	"fmt"
	"math"
)

func Example() {
	fmt.Println("Max int8:", math.MaxInt8)
	fmt.Println("Max int16:", math.MaxInt16)
	fmt.Println("Max int32:", math.MaxInt32)
	fmt.Println("Max int64:", math.MaxInt64)
	fmt.Println("Max uint8:", math.MaxUint8)
	fmt.Println("Max uint16:", math.MaxUint16)
	fmt.Println("Max uint32:", math.MaxUint32)
	fmt.Println("Max uint64:", uint64(math.MaxUint64)) // See https://github.com/golang/go/issues/19621
	fmt.Println("Max float32:", math.MaxFloat32)
	fmt.Println("Max float64:", math.MaxFloat64)
	// Output:
	// Max int8: 127
	// Max int16: 32767
	// Max int32: 2147483647
	// Max int64: 9223372036854775807
	// Max uint8: 255
	// Max uint16: 65535
	// Max uint32: 4294967295
	// Max uint64: 18446744073709551615
	// Max float32: 3.4028234663852886e+38
	// Max float64: 1.7976931348623157e+308
}
