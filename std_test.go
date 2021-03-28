package errf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Std_With(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		Std.With(WrapperFmtErrorw("wrapped")).CheckInt(1, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "wrapped: error")
}

func Test_CheckInt(t *testing.T) {
	value := 1

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt(value, nil))
		Std.CheckInt(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckIntSlice(t *testing.T) {
	value := []int{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckIntSlice(value, nil))
		Std.CheckIntSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt8(t *testing.T) {
	value := int8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt8(value, nil))
		Std.CheckInt8(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt8Slice(t *testing.T) {
	value := []int8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt8Slice(value, nil))
		Std.CheckInt8Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt16(t *testing.T) {
	value := int16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt16(value, nil))
		Std.CheckInt16(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt16Slice(t *testing.T) {
	value := []int16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt16Slice(value, nil))
		Std.CheckInt16Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt32(t *testing.T) {
	value := int32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt32(value, nil))
		Std.CheckInt32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt32Slice(t *testing.T) {
	value := []int32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt32Slice(value, nil))
		Std.CheckInt32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt64(t *testing.T) {
	value := int64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt64(value, nil))
		Std.CheckInt64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt64Slice(t *testing.T) {
	value := []int64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckInt64Slice(value, nil))
		Std.CheckInt64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint(t *testing.T) {
	value := uint(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint(value, nil))
		Std.CheckUint(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintSlice(t *testing.T) {
	value := []uint{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUintSlice(value, nil))
		Std.CheckUintSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint8(t *testing.T) {
	value := uint8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint8(value, nil))
		Std.CheckUint8(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint8Slice(t *testing.T) {
	value := []uint8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint8Slice(value, nil))
		Std.CheckUint8Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint16(t *testing.T) {
	value := uint16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint16(value, nil))
		Std.CheckUint16(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint16Slice(t *testing.T) {
	value := []uint16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint16Slice(value, nil))
		Std.CheckUint16Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint32(t *testing.T) {
	value := uint32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint32(value, nil))
		Std.CheckUint32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint32Slice(t *testing.T) {
	value := []uint32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint32Slice(value, nil))
		Std.CheckUint32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint64(t *testing.T) {
	value := uint64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint64(value, nil))
		Std.CheckUint64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint64Slice(t *testing.T) {
	value := []uint64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUint64Slice(value, nil))
		Std.CheckUint64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintptr(t *testing.T) {
	value := uintptr(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUintptr(value, nil))
		Std.CheckUintptr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintptrSlice(t *testing.T) {
	value := []uintptr{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckUintptrSlice(value, nil))
		Std.CheckUintptrSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat32(t *testing.T) {
	value := float32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckFloat32(value, nil))
		Std.CheckFloat32(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat32Slice(t *testing.T) {
	value := []float32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckFloat32Slice(value, nil))
		Std.CheckFloat32Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat64(t *testing.T) {
	value := float64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckFloat64(value, nil))
		Std.CheckFloat64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat64Slice(t *testing.T) {
	value := []float64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckFloat64Slice(value, nil))
		Std.CheckFloat64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex64(t *testing.T) {
	value := complex64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckComplex64(value, nil))
		Std.CheckComplex64(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex64Slice(t *testing.T) {
	value := []complex64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckComplex64Slice(value, nil))
		Std.CheckComplex64Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex128(t *testing.T) {
	value := complex128(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckComplex128(value, nil))
		Std.CheckComplex128(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex128Slice(t *testing.T) {
	value := []complex128{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckComplex128Slice(value, nil))
		Std.CheckComplex128Slice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckBool(t *testing.T) {
	value := true

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckBool(value, nil))
		Std.CheckBool(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckBoolSlice(t *testing.T) {
	value := []bool{true}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckBoolSlice(value, nil))
		Std.CheckBoolSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckString(t *testing.T) {
	value := "123"

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckString(value, nil))
		Std.CheckString(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckStringSlice(t *testing.T) {
	value := []string{"123"}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckStringSlice(value, nil))
		Std.CheckStringSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckByte(t *testing.T) {
	value := byte(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckByte(value, nil))
		Std.CheckByte(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckByteSlice(t *testing.T) {
	value := []byte{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckByteSlice(value, nil))
		Std.CheckByteSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckRune(t *testing.T) {
	value := rune(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckRune(value, nil))
		Std.CheckRune(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckRuneSlice(t *testing.T) {
	value := []rune{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, value, Std.CheckRuneSlice(value, nil))
		Std.CheckRuneSlice(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckIntErr(t *testing.T) {
	value := 1

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckIntErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckIntErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckIntSliceErr(t *testing.T) {
	value := []int{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckIntSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckIntSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt8Err(t *testing.T) {
	value := int8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt8Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt8Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt8SliceErr(t *testing.T) {
	value := []int8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt8SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt8SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt16Err(t *testing.T) {
	value := int16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt16Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt16Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt16SliceErr(t *testing.T) {
	value := []int16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt16SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt16SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt32Err(t *testing.T) {
	value := int32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt32SliceErr(t *testing.T) {
	value := []int32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt64Err(t *testing.T) {
	value := int64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckInt64SliceErr(t *testing.T) {
	value := []int64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckInt64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckInt64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintErr(t *testing.T) {
	value := uint(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUintErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUintErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintSliceErr(t *testing.T) {
	value := []uint{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUintSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUintSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint8Err(t *testing.T) {
	value := uint8(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint8Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint8Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint8SliceErr(t *testing.T) {
	value := []uint8{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint8SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint8SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint16Err(t *testing.T) {
	value := uint16(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint16Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint16Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint16SliceErr(t *testing.T) {
	value := []uint16{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint16SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint16SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint32Err(t *testing.T) {
	value := uint32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint32SliceErr(t *testing.T) {
	value := []uint32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint64Err(t *testing.T) {
	value := uint64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUint64SliceErr(t *testing.T) {
	value := []uint64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUint64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUint64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintptrErr(t *testing.T) {
	value := uintptr(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUintptrErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUintptrErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckUintptrSliceErr(t *testing.T) {
	value := []uintptr{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckUintptrSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckUintptrSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat32Err(t *testing.T) {
	value := float32(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckFloat32Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckFloat32Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat32SliceErr(t *testing.T) {
	value := []float32{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckFloat32SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckFloat32SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat64Err(t *testing.T) {
	value := float64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckFloat64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckFloat64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckFloat64SliceErr(t *testing.T) {
	value := []float64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckFloat64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckFloat64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex64Err(t *testing.T) {
	value := complex64(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckComplex64Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckComplex64Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex64SliceErr(t *testing.T) {
	value := []complex64{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckComplex64SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckComplex64SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex128Err(t *testing.T) {
	value := complex128(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckComplex128Err(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckComplex128Err(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckComplex128SliceErr(t *testing.T) {
	value := []complex128{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckComplex128SliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckComplex128SliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckBoolErr(t *testing.T) {
	value := true

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckBoolErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckBoolErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckBoolSliceErr(t *testing.T) {
	value := []bool{true}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckBoolSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckBoolSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckStringErr(t *testing.T) {
	value := "123"

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckStringErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckStringErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckStringSliceErr(t *testing.T) {
	value := []string{"123"}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckStringSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckStringSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckByteErr(t *testing.T) {
	value := byte(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckByteErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckByteErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckByteSliceErr(t *testing.T) {
	value := []byte{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckByteSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckByteSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckRuneErr(t *testing.T) {
	value := rune(1)

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckRuneErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckRuneErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func Test_CheckRuneSliceErr(t *testing.T) {
	value := []rune{1}

	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		actual, checkErr := Std.CheckRuneSliceErr(value, nil)
		assert.Equal(t, value, actual)
		assert.Nil(t, checkErr)
		_, _ = Std.CheckRuneSliceErr(value, fmt.Errorf("error"))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
