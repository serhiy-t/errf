package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Std_With(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Std.With(WrapperFmtErrorw("wrapped")).TryInt(1, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_TryInt(t *testing.T) {
	value := 1

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt(value, nil))
		Std.TryInt(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryIntSlice(t *testing.T) {
	value := []int{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryIntSlice(value, nil))
		Std.TryIntSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt8(t *testing.T) {
	value := int8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt8(value, nil))
		Std.TryInt8(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt8Slice(t *testing.T) {
	value := []int8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt8Slice(value, nil))
		Std.TryInt8Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt16(t *testing.T) {
	value := int16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt16(value, nil))
		Std.TryInt16(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt16Slice(t *testing.T) {
	value := []int16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt16Slice(value, nil))
		Std.TryInt16Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt32(t *testing.T) {
	value := int32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt32(value, nil))
		Std.TryInt32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt32Slice(t *testing.T) {
	value := []int32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt32Slice(value, nil))
		Std.TryInt32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt64(t *testing.T) {
	value := int64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt64(value, nil))
		Std.TryInt64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt64Slice(t *testing.T) {
	value := []int64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryInt64Slice(value, nil))
		Std.TryInt64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint(t *testing.T) {
	value := uint(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint(value, nil))
		Std.TryUint(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintSlice(t *testing.T) {
	value := []uint{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUintSlice(value, nil))
		Std.TryUintSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint8(t *testing.T) {
	value := uint8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint8(value, nil))
		Std.TryUint8(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint8Slice(t *testing.T) {
	value := []uint8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint8Slice(value, nil))
		Std.TryUint8Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint16(t *testing.T) {
	value := uint16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint16(value, nil))
		Std.TryUint16(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint16Slice(t *testing.T) {
	value := []uint16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint16Slice(value, nil))
		Std.TryUint16Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint32(t *testing.T) {
	value := uint32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint32(value, nil))
		Std.TryUint32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint32Slice(t *testing.T) {
	value := []uint32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint32Slice(value, nil))
		Std.TryUint32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint64(t *testing.T) {
	value := uint64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint64(value, nil))
		Std.TryUint64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint64Slice(t *testing.T) {
	value := []uint64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUint64Slice(value, nil))
		Std.TryUint64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintptr(t *testing.T) {
	value := uintptr(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUintptr(value, nil))
		Std.TryUintptr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintptrSlice(t *testing.T) {
	value := []uintptr{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryUintptrSlice(value, nil))
		Std.TryUintptrSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat32(t *testing.T) {
	value := float32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryFloat32(value, nil))
		Std.TryFloat32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat32Slice(t *testing.T) {
	value := []float32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryFloat32Slice(value, nil))
		Std.TryFloat32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat64(t *testing.T) {
	value := float64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryFloat64(value, nil))
		Std.TryFloat64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat64Slice(t *testing.T) {
	value := []float64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryFloat64Slice(value, nil))
		Std.TryFloat64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex64(t *testing.T) {
	value := complex64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryComplex64(value, nil))
		Std.TryComplex64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex64Slice(t *testing.T) {
	value := []complex64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryComplex64Slice(value, nil))
		Std.TryComplex64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex128(t *testing.T) {
	value := complex128(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryComplex128(value, nil))
		Std.TryComplex128(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex128Slice(t *testing.T) {
	value := []complex128{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryComplex128Slice(value, nil))
		Std.TryComplex128Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryBool(t *testing.T) {
	value := true

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryBool(value, nil))
		Std.TryBool(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryBoolSlice(t *testing.T) {
	value := []bool{true}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryBoolSlice(value, nil))
		Std.TryBoolSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryString(t *testing.T) {
	value := "123"

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryString(value, nil))
		Std.TryString(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryStringSlice(t *testing.T) {
	value := []string{"123"}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryStringSlice(value, nil))
		Std.TryStringSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryByte(t *testing.T) {
	value := byte(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryByte(value, nil))
		Std.TryByte(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryByteSlice(t *testing.T) {
	value := []byte{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryByteSlice(value, nil))
		Std.TryByteSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryRune(t *testing.T) {
	value := rune(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryRune(value, nil))
		Std.TryRune(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryRuneSlice(t *testing.T) {
	value := []rune{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.TryRuneSlice(value, nil))
		Std.TryRuneSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryIntErr(t *testing.T) {
	value := 1

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryIntErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryIntErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryIntSliceErr(t *testing.T) {
	value := []int{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryIntSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryIntSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt8Err(t *testing.T) {
	value := int8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt8Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt8Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt8SliceErr(t *testing.T) {
	value := []int8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt8SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt8SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt16Err(t *testing.T) {
	value := int16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt16Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt16Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt16SliceErr(t *testing.T) {
	value := []int16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt16SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt16SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt32Err(t *testing.T) {
	value := int32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt32SliceErr(t *testing.T) {
	value := []int32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt64Err(t *testing.T) {
	value := int64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryInt64SliceErr(t *testing.T) {
	value := []int64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryInt64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryInt64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintErr(t *testing.T) {
	value := uint(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUintErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUintErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintSliceErr(t *testing.T) {
	value := []uint{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUintSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUintSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint8Err(t *testing.T) {
	value := uint8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint8Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint8Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint8SliceErr(t *testing.T) {
	value := []uint8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint8SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint8SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint16Err(t *testing.T) {
	value := uint16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint16Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint16Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint16SliceErr(t *testing.T) {
	value := []uint16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint16SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint16SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint32Err(t *testing.T) {
	value := uint32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint32SliceErr(t *testing.T) {
	value := []uint32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint64Err(t *testing.T) {
	value := uint64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUint64SliceErr(t *testing.T) {
	value := []uint64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUint64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUint64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintptrErr(t *testing.T) {
	value := uintptr(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUintptrErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUintptrErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryUintptrSliceErr(t *testing.T) {
	value := []uintptr{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryUintptrSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryUintptrSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat32Err(t *testing.T) {
	value := float32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryFloat32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryFloat32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat32SliceErr(t *testing.T) {
	value := []float32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryFloat32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryFloat32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat64Err(t *testing.T) {
	value := float64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryFloat64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryFloat64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryFloat64SliceErr(t *testing.T) {
	value := []float64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryFloat64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryFloat64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex64Err(t *testing.T) {
	value := complex64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryComplex64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryComplex64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex64SliceErr(t *testing.T) {
	value := []complex64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryComplex64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryComplex64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex128Err(t *testing.T) {
	value := complex128(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryComplex128Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryComplex128Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryComplex128SliceErr(t *testing.T) {
	value := []complex128{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryComplex128SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryComplex128SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryBoolErr(t *testing.T) {
	value := true

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryBoolErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryBoolErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryBoolSliceErr(t *testing.T) {
	value := []bool{true}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryBoolSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryBoolSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryStringErr(t *testing.T) {
	value := "123"

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryStringErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryStringErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryStringSliceErr(t *testing.T) {
	value := []string{"123"}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryStringSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryStringSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryByteErr(t *testing.T) {
	value := byte(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryByteErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryByteErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryByteSliceErr(t *testing.T) {
	value := []byte{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryByteSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryByteSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryRuneErr(t *testing.T) {
	value := rune(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryRuneErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryRuneErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_TryRuneSliceErr(t *testing.T) {
	value := []rune{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, tryErr := Std.TryRuneSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, tryErr)
		Std.TryRuneSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
