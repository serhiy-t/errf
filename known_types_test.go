package errflow

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKnownTypes_Int(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Int(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IntSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []int{1}, IntSlice([]int{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int8(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Int8(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int8Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []int8{1}, Int8Slice([]int8{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int16(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Int16(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int16Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []int16{1}, Int16Slice([]int16{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int32(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Int32(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int32Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []int32{1}, Int32Slice([]int32{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int64(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Int64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Int64Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []int64{1}, Int64Slice([]int64{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uint(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_UintSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uint{1}, UintSlice([]uint{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint8(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uint8(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint8Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uint8{1}, Uint8Slice([]uint8{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint16(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uint16(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint16Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uint16{1}, Uint16Slice([]uint16{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint32(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uint32(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint32Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uint32{1}, Uint32Slice([]uint32{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint64(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uint64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uint64Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uint64{1}, Uint64Slice([]uint64{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Uintptr(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Uintptr(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_UintptrSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []uintptr{1}, UintptrSlice([]uintptr{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float32(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1.0, Float32(1.0, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float32Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []float32{1.0}, Float32Slice([]float32{1.0}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float64(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Float64(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Float64Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []float64{1.0}, Float64Slice([]float64{1.0}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex64(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1+1i, Complex64(1+1i, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex64Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []complex64{1 + 1i}, Complex64Slice([]complex64{1 + 1i}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex128(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1+1i, Complex128(1+1i, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Complex128Slice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []complex128{1 + 1i}, Complex128Slice([]complex128{1 + 1i}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Bool(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, true, Bool(true, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BoolSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []bool{true}, BoolSlice([]bool{true}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_String(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, "value", String("value", fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_StringSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []string{"value"}, StringSlice([]string{"value"}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Byte(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Byte(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_ByteSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []byte{1}, ByteSlice([]byte{1}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_Rune(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, 1, Rune(1, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_RuneSlice(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, []rune{1}, RuneSlice([]rune{1}, fmt.Errorf("error")))
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
		defer Catch().WriteTo(&err)
		assert.Equal(t, testStruct{}, IoWriter(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoWriteCloser(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, testStruct{}, IoWriteCloser(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoReader(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, testStruct{}, IoReader(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_IoReadCloser(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		assert.Equal(t, testStruct{}, IoReadCloser(testStruct{}, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioWriter(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		instance := bufio.NewWriter(testStruct{})
		assert.Same(t, instance, BufioWriter(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioReader(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		instance := bufio.NewReader(testStruct{})
		assert.Same(t, instance, BufioReader(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}

func TestKnownTypes_BufioReadWriter(t *testing.T) {
	fn := func() (err error) {
		defer Catch().WriteTo(&err)
		instance := bufio.NewReadWriter(bufio.NewReader(testStruct{}), bufio.NewWriter(testStruct{}))
		assert.Same(t, instance, BufioReadWriter(instance, fmt.Errorf("error")))
		return nil
	}

	assert.EqualError(t, fn(), "error")
}
