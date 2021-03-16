package errf

// Std contains collection of Check* functions for built-in types.
var Std = StdErrflow{}

// StdErrflow implements Check* functions for built-in types.
//
// Clients should not instantiate StdErrflow, use 'errf.Std' instead.
type StdErrflow struct {
	errflow *Errflow
}

// With implements Errflow.With(...) for built-in types.
func (ef StdErrflow) With(options ...ErrflowOption) StdErrflow {
	return StdErrflow{errflow: ef.errflow.With(options...)}
}

// CheckInt calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt(value int, err error) int {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckIntSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckIntSlice(value []int, err error) []int {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt8 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt8(value int8, err error) int8 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt8Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt8Slice(value []int8, err error) []int8 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt16 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt16(value int16, err error) int16 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt16Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt16Slice(value []int16, err error) []int16 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt32 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt32(value int32, err error) int32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt32Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt32Slice(value []int32, err error) []int32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt64 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt64(value int64, err error) int64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckInt64Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckInt64Slice(value []int64, err error) []int64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint(value uint, err error) uint {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUintSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUintSlice(value []uint, err error) []uint {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint8 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint8(value uint8, err error) uint8 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint8Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint8Slice(value []uint8, err error) []uint8 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint16 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint16(value uint16, err error) uint16 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint16Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint16Slice(value []uint16, err error) []uint16 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint32 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint32(value uint32, err error) uint32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint32Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint32Slice(value []uint32, err error) []uint32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint64 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint64(value uint64, err error) uint64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUint64Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUint64Slice(value []uint64, err error) []uint64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUintptr calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUintptr(value uintptr, err error) uintptr {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckUintptrSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckUintptrSlice(value []uintptr, err error) []uintptr {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckFloat32 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckFloat32(value float32, err error) float32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckFloat32Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckFloat32Slice(value []float32, err error) []float32 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckFloat64 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckFloat64(value float64, err error) float64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckFloat64Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckFloat64Slice(value []float64, err error) []float64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckComplex64 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckComplex64(value complex64, err error) complex64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckComplex64Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckComplex64Slice(value []complex64, err error) []complex64 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckComplex128 calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckComplex128(value complex128, err error) complex128 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckComplex128Slice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckComplex128Slice(value []complex128, err error) []complex128 {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckBool calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckBool(value bool, err error) bool {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckBoolSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckBoolSlice(value []bool, err error) []bool {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckString calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckString(value string, err error) string {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckStringSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckStringSlice(value []string, err error) []string {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckByte calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckByte(value byte, err error) byte {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckByteSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckByteSlice(value []byte, err error) []byte {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckRune calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckRune(value rune, err error) rune {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckRuneSlice calls errf.Check and returns a typed value from a function call.
func (ef StdErrflow) CheckRuneSlice(value []rune, err error) []rune {
	ef.errflow.ImplementCheck(recover(), err)
	return value
}

// CheckIntErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckIntErr(value int, err error) (int, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckIntSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckIntSliceErr(value []int, err error) ([]int, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt8Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt8Err(value int8, err error) (int8, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt8SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt8SliceErr(value []int8, err error) ([]int8, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt16Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt16Err(value int16, err error) (int16, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt16SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt16SliceErr(value []int16, err error) ([]int16, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt32Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt32Err(value int32, err error) (int32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt32SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt32SliceErr(value []int32, err error) ([]int32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt64Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt64Err(value int64, err error) (int64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckInt64SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckInt64SliceErr(value []int64, err error) ([]int64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUintErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUintErr(value uint, err error) (uint, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUintSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUintSliceErr(value []uint, err error) ([]uint, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint8Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint8Err(value uint8, err error) (uint8, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint8SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint8SliceErr(value []uint8, err error) ([]uint8, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint16Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint16Err(value uint16, err error) (uint16, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint16SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint16SliceErr(value []uint16, err error) ([]uint16, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint32Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint32Err(value uint32, err error) (uint32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint32SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint32SliceErr(value []uint32, err error) ([]uint32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint64Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint64Err(value uint64, err error) (uint64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUint64SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUint64SliceErr(value []uint64, err error) ([]uint64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUintptrErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUintptrErr(value uintptr, err error) (uintptr, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckUintptrSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckUintptrSliceErr(value []uintptr, err error) ([]uintptr, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckFloat32Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckFloat32Err(value float32, err error) (float32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckFloat32SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckFloat32SliceErr(value []float32, err error) ([]float32, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckFloat64Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckFloat64Err(value float64, err error) (float64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckFloat64SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckFloat64SliceErr(value []float64, err error) ([]float64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckComplex64Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckComplex64Err(value complex64, err error) (complex64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckComplex64SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckComplex64SliceErr(value []complex64, err error) ([]complex64, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckComplex128Err calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckComplex128Err(value complex128, err error) (complex128, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckComplex128SliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckComplex128SliceErr(value []complex128, err error) ([]complex128, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckBoolErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckBoolErr(value bool, err error) (bool, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckBoolSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckBoolSliceErr(value []bool, err error) ([]bool, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckStringErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckStringErr(value string, err error) (string, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckStringSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckStringSliceErr(value []string, err error) ([]string, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckByteErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckByteErr(value byte, err error) (byte, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckByteSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckByteSliceErr(value []byte, err error) ([]byte, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckRuneErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckRuneErr(value rune, err error) (rune, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}

// CheckRuneSliceErr calls errf.Check and returns a typed value and error from a function call.
func (ef StdErrflow) CheckRuneSliceErr(value []rune, err error) ([]rune, error) {
	ef.errflow.ImplementCheck(recover(), err)
	return value, err
}
