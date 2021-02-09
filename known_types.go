package errflow

import (
	"bufio"
	"io"
)

// CheckInt calls errflow.Check and returns a typed value from a function call.
func CheckInt(value int, err error) int {
	ImplementCheck(recover(), err)
	return value
}

// CheckIntSlice calls errflow.Check and returns a typed value from a function call.
func CheckIntSlice(value []int, err error) []int {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt8 calls errflow.Check and returns a typed value from a function call.
func CheckInt8(value int8, err error) int8 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt8Slice calls errflow.Check and returns a typed value from a function call.
func CheckInt8Slice(value []int8, err error) []int8 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt16 calls errflow.Check and returns a typed value from a function call.
func CheckInt16(value int16, err error) int16 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt16Slice calls errflow.Check and returns a typed value from a function call.
func CheckInt16Slice(value []int16, err error) []int16 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt32 calls errflow.Check and returns a typed value from a function call.
func CheckInt32(value int32, err error) int32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt32Slice calls errflow.Check and returns a typed value from a function call.
func CheckInt32Slice(value []int32, err error) []int32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt64 calls errflow.Check and returns a typed value from a function call.
func CheckInt64(value int64, err error) int64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckInt64Slice calls errflow.Check and returns a typed value from a function call.
func CheckInt64Slice(value []int64, err error) []int64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint calls errflow.Check and returns a typed value from a function call.
func CheckUint(value uint, err error) uint {
	ImplementCheck(recover(), err)
	return value
}

// CheckUintSlice calls errflow.Check and returns a typed value from a function call.
func CheckUintSlice(value []uint, err error) []uint {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint8 calls errflow.Check and returns a typed value from a function call.
func CheckUint8(value uint8, err error) uint8 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint8Slice calls errflow.Check and returns a typed value from a function call.
func CheckUint8Slice(value []uint8, err error) []uint8 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint16 calls errflow.Check and returns a typed value from a function call.
func CheckUint16(value uint16, err error) uint16 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint16Slice calls errflow.Check and returns a typed value from a function call.
func CheckUint16Slice(value []uint16, err error) []uint16 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint32 calls errflow.Check and returns a typed value from a function call.
func CheckUint32(value uint32, err error) uint32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint32Slice calls errflow.Check and returns a typed value from a function call.
func CheckUint32Slice(value []uint32, err error) []uint32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint64 calls errflow.Check and returns a typed value from a function call.
func CheckUint64(value uint64, err error) uint64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUint64Slice calls errflow.Check and returns a typed value from a function call.
func CheckUint64Slice(value []uint64, err error) []uint64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckUintptr calls errflow.Check and returns a typed value from a function call.
func CheckUintptr(value uintptr, err error) uintptr {
	ImplementCheck(recover(), err)
	return value
}

// CheckUintptrSlice calls errflow.Check and returns a typed value from a function call.
func CheckUintptrSlice(value []uintptr, err error) []uintptr {
	ImplementCheck(recover(), err)
	return value
}

// CheckFloat32 calls errflow.Check and returns a typed value from a function call.
func CheckFloat32(value float32, err error) float32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckFloat32Slice calls errflow.Check and returns a typed value from a function call.
func CheckFloat32Slice(value []float32, err error) []float32 {
	ImplementCheck(recover(), err)
	return value
}

// CheckFloat64 calls errflow.Check and returns a typed value from a function call.
func CheckFloat64(value float64, err error) float64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckFloat64Slice calls errflow.Check and returns a typed value from a function call.
func CheckFloat64Slice(value []float64, err error) []float64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckComplex64 calls errflow.Check and returns a typed value from a function call.
func CheckComplex64(value complex64, err error) complex64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckComplex64Slice calls errflow.Check and returns a typed value from a function call.
func CheckComplex64Slice(value []complex64, err error) []complex64 {
	ImplementCheck(recover(), err)
	return value
}

// CheckComplex128 calls errflow.Check and returns a typed value from a function call.
func CheckComplex128(value complex128, err error) complex128 {
	ImplementCheck(recover(), err)
	return value
}

// CheckComplex128Slice calls errflow.Check and returns a typed value from a function call.
func CheckComplex128Slice(value []complex128, err error) []complex128 {
	ImplementCheck(recover(), err)
	return value
}

// CheckBool calls errflow.Check and returns a typed value from a function call.
func CheckBool(value bool, err error) bool {
	ImplementCheck(recover(), err)
	return value
}

// CheckBoolSlice calls errflow.Check and returns a typed value from a function call.
func CheckBoolSlice(value []bool, err error) []bool {
	ImplementCheck(recover(), err)
	return value
}

// CheckString calls errflow.Check and returns a typed value from a function call.
func CheckString(value string, err error) string {
	ImplementCheck(recover(), err)
	return value
}

// CheckStringSlice calls errflow.Check and returns a typed value from a function call.
func CheckStringSlice(value []string, err error) []string {
	ImplementCheck(recover(), err)
	return value
}

// CheckByte calls errflow.Check and returns a typed value from a function call.
func CheckByte(value byte, err error) byte {
	ImplementCheck(recover(), err)
	return value
}

// CheckByteSlice calls errflow.Check and returns a typed value from a function call.
func CheckByteSlice(value []byte, err error) []byte {
	ImplementCheck(recover(), err)
	return value
}

// CheckRune calls errflow.Check and returns a typed value from a function call.
func CheckRune(value rune, err error) rune {
	ImplementCheck(recover(), err)
	return value
}

// CheckRuneSlice calls errflow.Check and returns a typed value from a function call.
func CheckRuneSlice(value []rune, err error) []rune {
	ImplementCheck(recover(), err)
	return value
}

// CheckIoWriter calls errflow.Check and returns a typed value from a function call.
func CheckIoWriter(value io.Writer, err error) io.Writer {
	ImplementCheck(recover(), err)
	return value
}

// CheckIoWriteCloser calls errflow.Check and returns a typed value from a function call.
func CheckIoWriteCloser(value io.WriteCloser, err error) io.WriteCloser {
	ImplementCheck(recover(), err)
	return value
}

// CheckIoReader calls errflow.Check and returns a typed value from a function call.
func CheckIoReader(value io.Reader, err error) io.Reader {
	ImplementCheck(recover(), err)
	return value
}

// CheckIoReadCloser calls errflow.Check and returns a typed value from a function call.
func CheckIoReadCloser(value io.ReadCloser, err error) io.ReadCloser {
	ImplementCheck(recover(), err)
	return value
}

// CheckBufioWriter calls errflow.Check and returns a typed value from a function call.
func CheckBufioWriter(value *bufio.Writer, err error) *bufio.Writer {
	ImplementCheck(recover(), err)
	return value
}

// CheckBufioReader calls errflow.Check and returns a typed value from a function call.
func CheckBufioReader(value *bufio.Reader, err error) *bufio.Reader {
	ImplementCheck(recover(), err)
	return value
}

// CheckBufioReadWriter calls errflow.Check and returns a typed value from a function call.
func CheckBufioReadWriter(value *bufio.ReadWriter, err error) *bufio.ReadWriter {
	ImplementCheck(recover(), err)
	return value
}
