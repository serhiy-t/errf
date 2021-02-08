package errflow

import (
	"bufio"
	"io"
)

// Primitive types

// Ints

func Int(value int, err error) int {
	ImplementCheck(recover(), err)
	return value
}

func IntSlice(value []int, err error) []int {
	ImplementCheck(recover(), err)
	return value
}

func Int8(value int8, err error) int8 {
	ImplementCheck(recover(), err)
	return value
}

func Int8Slice(value []int8, err error) []int8 {
	ImplementCheck(recover(), err)
	return value
}

func Int16(value int16, err error) int16 {
	ImplementCheck(recover(), err)
	return value
}

func Int16Slice(value []int16, err error) []int16 {
	ImplementCheck(recover(), err)
	return value
}

func Int32(value int32, err error) int32 {
	ImplementCheck(recover(), err)
	return value
}

func Int32Slice(value []int32, err error) []int32 {
	ImplementCheck(recover(), err)
	return value
}

func Int64(value int64, err error) int64 {
	ImplementCheck(recover(), err)
	return value
}

func Int64Slice(value []int64, err error) []int64 {
	ImplementCheck(recover(), err)
	return value
}

// Uints

func Uint(value uint, err error) uint {
	ImplementCheck(recover(), err)
	return value
}

func UintSlice(value []uint, err error) []uint {
	ImplementCheck(recover(), err)
	return value
}

func Uint8(value uint8, err error) uint8 {
	ImplementCheck(recover(), err)
	return value
}

func Uint8Slice(value []uint8, err error) []uint8 {
	ImplementCheck(recover(), err)
	return value
}

func Uint16(value uint16, err error) uint16 {
	ImplementCheck(recover(), err)
	return value
}

func Uint16Slice(value []uint16, err error) []uint16 {
	ImplementCheck(recover(), err)
	return value
}

func Uint32(value uint32, err error) uint32 {
	ImplementCheck(recover(), err)
	return value
}

func Uint32Slice(value []uint32, err error) []uint32 {
	ImplementCheck(recover(), err)
	return value
}

func Uint64(value uint64, err error) uint64 {
	ImplementCheck(recover(), err)
	return value
}

func Uint64Slice(value []uint64, err error) []uint64 {
	ImplementCheck(recover(), err)
	return value
}

func Uintptr(value uintptr, err error) uintptr {
	ImplementCheck(recover(), err)
	return value
}

func UintptrSlice(value []uintptr, err error) []uintptr {
	ImplementCheck(recover(), err)
	return value
}

// Floats

func Float32(value float32, err error) float32 {
	ImplementCheck(recover(), err)
	return value
}

func Float32Slice(value []float32, err error) []float32 {
	ImplementCheck(recover(), err)
	return value
}

func Float64(value float64, err error) float64 {
	ImplementCheck(recover(), err)
	return value
}

func Float64Slice(value []float64, err error) []float64 {
	ImplementCheck(recover(), err)
	return value
}

// Complex

func Complex64(value complex64, err error) complex64 {
	ImplementCheck(recover(), err)
	return value
}

func Complex64Slice(value []complex64, err error) []complex64 {
	ImplementCheck(recover(), err)
	return value
}

func Complex128(value complex128, err error) complex128 {
	ImplementCheck(recover(), err)
	return value
}

func Complex128Slice(value []complex128, err error) []complex128 {
	ImplementCheck(recover(), err)
	return value
}

// Other primitive types

func Bool(value bool, err error) bool {
	ImplementCheck(recover(), err)
	return value
}

func BoolSlice(value []bool, err error) []bool {
	ImplementCheck(recover(), err)
	return value
}

func String(value string, err error) string {
	ImplementCheck(recover(), err)
	return value
}

func StringSlice(value []string, err error) []string {
	ImplementCheck(recover(), err)
	return value
}

func Byte(value byte, err error) byte {
	ImplementCheck(recover(), err)
	return value
}

func ByteSlice(value []byte, err error) []byte {
	ImplementCheck(recover(), err)
	return value
}

func Rune(value rune, err error) rune {
	ImplementCheck(recover(), err)
	return value
}

func RuneSlice(value []rune, err error) []rune {
	ImplementCheck(recover(), err)
	return value
}

// Other well-known types

func IoWriter(value io.Writer, err error) io.Writer {
	ImplementCheck(recover(), err)
	return value
}

func IoWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ImplementCheck(recover(), err)
	return value
}

func IoReader(value io.Reader, err error) io.Reader {
	ImplementCheck(recover(), err)
	return value
}

func IoReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ImplementCheck(recover(), err)
	return value
}

func BufioWriter(value *bufio.Writer, err error) *bufio.Writer {
	ImplementCheck(recover(), err)
	return value
}

func BufioReader(value *bufio.Reader, err error) *bufio.Reader {
	ImplementCheck(recover(), err)
	return value
}

func BufioReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ImplementCheck(recover(), err)
	return value
}
