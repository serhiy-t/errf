package errflow

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownTypes_Int(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckInt(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IntSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []int{1}, CheckIntSlice([]int{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int8(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckInt8(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int8Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []int8{1}, CheckInt8Slice([]int8{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int16(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckInt16(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int16Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []int16{1}, CheckInt16Slice([]int16{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int32(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckInt32(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int32Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []int32{1}, CheckInt32Slice([]int32{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int64(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckInt64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int64Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []int64{1}, CheckInt64Slice([]int64{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUint(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_UintSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uint{1}, CheckUintSlice([]uint{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint8(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUint8(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint8Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uint8{1}, CheckUint8Slice([]uint8{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint16(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUint16(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint16Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uint16{1}, CheckUint16Slice([]uint16{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint32(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUint32(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint32Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uint32{1}, CheckUint32Slice([]uint32{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint64(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUint64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint64Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uint64{1}, CheckUint64Slice([]uint64{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uintptr(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckUintptr(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_UintptrSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []uintptr{1}, CheckUintptrSlice([]uintptr{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float32(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1.0, CheckFloat32(1.0, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float32Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []float32{1.0}, CheckFloat32Slice([]float32{1.0}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float64(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckFloat64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float64Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []float64{1.0}, CheckFloat64Slice([]float64{1.0}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex64(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1+1i, CheckComplex64(1+1i, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex64Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []complex64{1 + 1i}, CheckComplex64Slice([]complex64{1 + 1i}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex128(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1+1i, CheckComplex128(1+1i, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex128Slice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []complex128{1 + 1i}, CheckComplex128Slice([]complex128{1 + 1i}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Bool(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, true, CheckBool(true, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BoolSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []bool{true}, CheckBoolSlice([]bool{true}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_String(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, "value", CheckString("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_StringSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []string{"value"}, CheckStringSlice([]string{"value"}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Byte(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckByte(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_ByteSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []byte{1}, CheckByteSlice([]byte{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Rune(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, 1, CheckRune(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_RuneSlice(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, []rune{1}, CheckRuneSlice([]rune{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

type testStruct struct{}

func (ts testStruct) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (ts testStruct) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (ts testStruct) Close() (err error) {
	return nil
}

func TestKnownTypes_IoWriter(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, testStruct{}, CheckIoWriter(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoWriteCloser(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, testStruct{}, CheckIoWriteCloser(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoReader(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, testStruct{}, CheckIoReader(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoReadCloser(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		assert.Equal(t, testStruct{}, CheckIoReadCloser(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioWriter(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		instance := bufio.NewWriter(testStruct{})
		assert.Same(t, instance, CheckBufioWriter(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioReader(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		instance := bufio.NewReader(testStruct{})
		assert.Same(t, instance, CheckBufioReader(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioReadWriter(t *testing.T) {
	fn := func() (err error) {
		defer IfError().ThenAssignTo(&err)
		instance := bufio.NewReadWriter(bufio.NewReader(testStruct{}), bufio.NewWriter(testStruct{}))
		assert.Same(t, instance, CheckBufioReadWriter(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
