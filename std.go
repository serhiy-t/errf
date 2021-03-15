package errf

type StdErrflow struct {
	errflow *Errflow
}

func (ef StdErrflow) With(options ...ErrflowOption) StdErrflow {
	return StdErrflow{errflow: ef.errflow.With(options...)}
}

// TryInt calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt(value int, err error) int {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryIntSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryIntSlice(value []int, err error) []int {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt8 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt8(value int8, err error) int8 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt8Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt8Slice(value []int8, err error) []int8 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt16 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt16(value int16, err error) int16 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt16Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt16Slice(value []int16, err error) []int16 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt32(value int32, err error) int32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt32Slice(value []int32, err error) []int32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt64(value int64, err error) int64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryInt64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt64Slice(value []int64, err error) []int64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint(value uint, err error) uint {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUintSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintSlice(value []uint, err error) []uint {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint8 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint8(value uint8, err error) uint8 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint8Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint8Slice(value []uint8, err error) []uint8 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint16 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint16(value uint16, err error) uint16 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint16Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint16Slice(value []uint16, err error) []uint16 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint32(value uint32, err error) uint32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint32Slice(value []uint32, err error) []uint32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint64(value uint64, err error) uint64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUint64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint64Slice(value []uint64, err error) []uint64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUintptr calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintptr(value uintptr, err error) uintptr {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryUintptrSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintptrSlice(value []uintptr, err error) []uintptr {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryFloat32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat32(value float32, err error) float32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryFloat32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat32Slice(value []float32, err error) []float32 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryFloat64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat64(value float64, err error) float64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryFloat64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat64Slice(value []float64, err error) []float64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryComplex64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex64(value complex64, err error) complex64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryComplex64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex64Slice(value []complex64, err error) []complex64 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryComplex128 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex128(value complex128, err error) complex128 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryComplex128Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex128Slice(value []complex128, err error) []complex128 {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryBool calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryBool(value bool, err error) bool {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryBoolSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryBoolSlice(value []bool, err error) []bool {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryString calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryString(value string, err error) string {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryStringSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryStringSlice(value []string, err error) []string {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryByte calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryByte(value byte, err error) byte {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryByteSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryByteSlice(value []byte, err error) []byte {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryRune calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryRune(value rune, err error) rune {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// TryRuneSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryRuneSlice(value []rune, err error) []rune {
	ef.errflow.ImplementTry(recover(), err)
	return value
}

// Errs

// TryInt calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryIntErr(value int, err error) (int, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryIntSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryIntSliceErr(value []int, err error) ([]int, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt8 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt8Err(value int8, err error) (int8, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt8Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt8SliceErr(value []int8, err error) ([]int8, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt16 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt16Err(value int16, err error) (int16, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt16Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt16SliceErr(value []int16, err error) ([]int16, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt32Err(value int32, err error) (int32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt32SliceErr(value []int32, err error) ([]int32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt64Err(value int64, err error) (int64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryInt64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryInt64SliceErr(value []int64, err error) ([]int64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintErr(value uint, err error) (uint, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUintSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintSliceErr(value []uint, err error) ([]uint, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint8 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint8Err(value uint8, err error) (uint8, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint8Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint8SliceErr(value []uint8, err error) ([]uint8, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint16 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint16Err(value uint16, err error) (uint16, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint16Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint16SliceErr(value []uint16, err error) ([]uint16, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint32Err(value uint32, err error) (uint32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint32SliceErr(value []uint32, err error) ([]uint32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint64Err(value uint64, err error) (uint64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUint64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUint64SliceErr(value []uint64, err error) ([]uint64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUintptr calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintptrErr(value uintptr, err error) (uintptr, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryUintptrSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryUintptrSliceErr(value []uintptr, err error) ([]uintptr, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryFloat32 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat32Err(value float32, err error) (float32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryFloat32Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat32SliceErr(value []float32, err error) ([]float32, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryFloat64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat64Err(value float64, err error) (float64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryFloat64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryFloat64SliceErr(value []float64, err error) ([]float64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryComplex64 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex64Err(value complex64, err error) (complex64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryComplex64Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex64SliceErr(value []complex64, err error) ([]complex64, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryComplex128 calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex128Err(value complex128, err error) (complex128, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryComplex128Slice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryComplex128SliceErr(value []complex128, err error) ([]complex128, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryBool calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryBoolErr(value bool, err error) (bool, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryBoolSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryBoolSliceErr(value []bool, err error) ([]bool, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryString calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryStringErr(value string, err error) (string, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryStringSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryStringSliceErr(value []string, err error) ([]string, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryByte calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryByteErr(value byte, err error) (byte, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryByteSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryByteSliceErr(value []byte, err error) ([]byte, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryRune calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryRuneErr(value rune, err error) (rune, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}

// TryRuneSlice calls errflow.Try and returns a typed value from a function call.
func (ef StdErrflow) TryRuneSliceErr(value []rune, err error) ([]rune, error) {
	ef.errflow.ImplementTry(recover(), err)
	return value, err
}
