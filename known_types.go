package errflow

import (
	"bufio"
	"io"
)

// Int calls errflow.Check and returns a typed value from a function call.
func Int(value int, err error) int {
	ImplementCheck(recover(), err)
	return value
}

// IntSlice calls errflow.Check and returns a typed value from a function call.
func IntSlice(value []int, err error) []int {
	ImplementCheck(recover(), err)
	return value
}

// Int8 calls errflow.Check and returns a typed value from a function call.
func Int8(value int8, err error) int8 {
	ImplementCheck(recover(), err)
	return value
}

// Int8Slice calls errflow.Check and returns a typed value from a function call.
func Int8Slice(value []int8, err error) []int8 {
	ImplementCheck(recover(), err)
	return value
}

// Int16 calls errflow.Check and returns a typed value from a function call.
func Int16(value int16, err error) int16 {
	ImplementCheck(recover(), err)
	return value
}

// Int16Slice calls errflow.Check and returns a typed value from a function call.
func Int16Slice(value []int16, err error) []int16 {
	ImplementCheck(recover(), err)
	return value
}

// Int32 calls errflow.Check and returns a typed value from a function call.
func Int32(value int32, err error) int32 {
	ImplementCheck(recover(), err)
	return value
}

// Int32Slice calls errflow.Check and returns a typed value from a function call.
func Int32Slice(value []int32, err error) []int32 {
	ImplementCheck(recover(), err)
	return value
}

// Int64 calls errflow.Check and returns a typed value from a function call.
func Int64(value int64, err error) int64 {
	ImplementCheck(recover(), err)
	return value
}

// Int64Slice calls errflow.Check and returns a typed value from a function call.
func Int64Slice(value []int64, err error) []int64 {
	ImplementCheck(recover(), err)
	return value
}

// Uint calls errflow.Check and returns a typed value from a function call.
func Uint(value uint, err error) uint {
	ImplementCheck(recover(), err)
	return value
}

// UintSlice calls errflow.Check and returns a typed value from a function call.
func UintSlice(value []uint, err error) []uint {
	ImplementCheck(recover(), err)
	return value
}

// Uint8 calls errflow.Check and returns a typed value from a function call.
func Uint8(value uint8, err error) uint8 {
	ImplementCheck(recover(), err)
	return value
}

// Uint8Slice calls errflow.Check and returns a typed value from a function call.
func Uint8Slice(value []uint8, err error) []uint8 {
	ImplementCheck(recover(), err)
	return value
}

// Uint16 calls errflow.Check and returns a typed value from a function call.
func Uint16(value uint16, err error) uint16 {
	ImplementCheck(recover(), err)
	return value
}

// Uint16Slice calls errflow.Check and returns a typed value from a function call.
func Uint16Slice(value []uint16, err error) []uint16 {
	ImplementCheck(recover(), err)
	return value
}

// Uint32 calls errflow.Check and returns a typed value from a function call.
func Uint32(value uint32, err error) uint32 {
	ImplementCheck(recover(), err)
	return value
}

// Uint32Slice calls errflow.Check and returns a typed value from a function call.
func Uint32Slice(value []uint32, err error) []uint32 {
	ImplementCheck(recover(), err)
	return value
}

// Uint64 calls errflow.Check and returns a typed value from a function call.
func Uint64(value uint64, err error) uint64 {
	ImplementCheck(recover(), err)
	return value
}

// Uint64Slice calls errflow.Check and returns a typed value from a function call.
func Uint64Slice(value []uint64, err error) []uint64 {
	ImplementCheck(recover(), err)
	return value
}

// Uintptr calls errflow.Check and returns a typed value from a function call.
func Uintptr(value uintptr, err error) uintptr {
	ImplementCheck(recover(), err)
	return value
}

// UintptrSlice calls errflow.Check and returns a typed value from a function call.
func UintptrSlice(value []uintptr, err error) []uintptr {
	ImplementCheck(recover(), err)
	return value
}

// Float32 calls errflow.Check and returns a typed value from a function call.
func Float32(value float32, err error) float32 {
	ImplementCheck(recover(), err)
	return value
}

// Float32Slice calls errflow.Check and returns a typed value from a function call.
func Float32Slice(value []float32, err error) []float32 {
	ImplementCheck(recover(), err)
	return value
}

// Float64 calls errflow.Check and returns a typed value from a function call.
func Float64(value float64, err error) float64 {
	ImplementCheck(recover(), err)
	return value
}

// Float64Slice calls errflow.Check and returns a typed value from a function call.
func Float64Slice(value []float64, err error) []float64 {
	ImplementCheck(recover(), err)
	return value
}

// Complex64 calls errflow.Check and returns a typed value from a function call.
func Complex64(value complex64, err error) complex64 {
	ImplementCheck(recover(), err)
	return value
}

// Complex64Slice calls errflow.Check and returns a typed value from a function call.
func Complex64Slice(value []complex64, err error) []complex64 {
	ImplementCheck(recover(), err)
	return value
}

// Complex128 calls errflow.Check and returns a typed value from a function call.
func Complex128(value complex128, err error) complex128 {
	ImplementCheck(recover(), err)
	return value
}

// Complex128Slice calls errflow.Check and returns a typed value from a function call.
func Complex128Slice(value []complex128, err error) []complex128 {
	ImplementCheck(recover(), err)
	return value
}

// Bool calls errflow.Check and returns a typed value from a function call.
func Bool(value bool, err error) bool {
	ImplementCheck(recover(), err)
	return value
}

// BoolSlice calls errflow.Check and returns a typed value from a function call.
func BoolSlice(value []bool, err error) []bool {
	ImplementCheck(recover(), err)
	return value
}

// String calls errflow.Check and returns a typed value from a function call.
func String(value string, err error) string {
	ImplementCheck(recover(), err)
	return value
}

// StringSlice calls errflow.Check and returns a typed value from a function call.
func StringSlice(value []string, err error) []string {
	ImplementCheck(recover(), err)
	return value
}

// Byte calls errflow.Check and returns a typed value from a function call.
func Byte(value byte, err error) byte {
	ImplementCheck(recover(), err)
	return value
}

// ByteSlice calls errflow.Check and returns a typed value from a function call.
func ByteSlice(value []byte, err error) []byte {
	ImplementCheck(recover(), err)
	return value
}

// Rune calls errflow.Check and returns a typed value from a function call.
func Rune(value rune, err error) rune {
	ImplementCheck(recover(), err)
	return value
}

// RuneSlice calls errflow.Check and returns a typed value from a function call.
func RuneSlice(value []rune, err error) []rune {
	ImplementCheck(recover(), err)
	return value
}

// IoWriter calls errflow.Check and returns a typed value from a function call.
func IoWriter(value io.Writer, err error) io.Writer {
	ImplementCheck(recover(), err)
	return value
}

// IoWriteCloser calls errflow.Check and returns a typed value from a function call.
func IoWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ImplementCheck(recover(), err)
	return value
}

// IoReader calls errflow.Check and returns a typed value from a function call.
func IoReader(value io.Reader, err error) io.Reader {
	ImplementCheck(recover(), err)
	return value
}

// IoReadCloser calls errflow.Check and returns a typed value from a function call.
func IoReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ImplementCheck(recover(), err)
	return value
}

// BufioWriter calls errflow.Check and returns a typed value from a function call.
func BufioWriter(value *bufio.Writer, err error) *bufio.Writer {
	ImplementCheck(recover(), err)
	return value
}

// BufioReader calls errflow.Check and returns a typed value from a function call.
func BufioReader(value *bufio.Reader, err error) *bufio.Reader {
	ImplementCheck(recover(), err)
	return value
}

// BufioReadWriter calls errflow.Check and returns a typed value from a function call.
func BufioReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ImplementCheck(recover(), err)
	return value
}
