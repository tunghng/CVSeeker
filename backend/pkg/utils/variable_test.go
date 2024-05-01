package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCurrentTimePointer(t *testing.T) {
	timeNow := CurrentTimePointer()
	assert.NotNil(t, timeNow)
}

func TestStringToPointer(t *testing.T) {
	testString := "abc"
	result := StringToPointer(testString)
	assert.NotNil(t, result)

	testString = ""
	result = StringToPointer(testString)
	assert.Nil(t, result)
}

func TestPointerFloat64(t *testing.T) {
	result := PointerFloat64(float64(1))
	assert.NotNil(t, result)
	assert.Equal(t, float64(1), *result)
}

func TestPointerFloat32(t *testing.T) {
	result := PointerFloat32(float32(1))
	assert.NotNil(t, result)
	assert.Equal(t, float32(1), *result)
}

func TestPointerInt(t *testing.T) {
	result := PointerInt(int(1))
	assert.NotNil(t, result)
	assert.Equal(t, int(1), *result)
}

func TestPointerInt8(t *testing.T) {
	result := PointerInt8(int8(1))
	assert.NotNil(t, result)
	assert.Equal(t, int8(1), *result)
}

func TestPointerInt16(t *testing.T) {
	result := PointerInt16(int16(1))
	assert.NotNil(t, result)
	assert.Equal(t, int16(1), *result)
}

func TestPointerInt32(t *testing.T) {
	result := PointerInt32(int32(1))
	assert.NotNil(t, result)
	assert.Equal(t, int32(1), *result)
}

func TestPointerInt64(t *testing.T) {
	result := PointerInt64(int64(1))
	assert.NotNil(t, result)
	assert.Equal(t, int64(1), *result)
}

func TestPointerUInt(t *testing.T) {
	result := PointerUInt(uint(1))
	assert.NotNil(t, result)
	assert.Equal(t, uint(1), *result)
}

func TestPointerUInt8(t *testing.T) {
	result := PointerUInt8(uint8(1))
	assert.NotNil(t, result)
	assert.Equal(t, uint8(1), *result)
}

func TestPointerUInt16(t *testing.T) {
	result := PointerUInt16(uint16(1))
	assert.NotNil(t, result)
	assert.Equal(t, uint16(1), *result)
}

func TestPointerUInt32(t *testing.T) {
	result := PointerUInt32(uint32(1))
	assert.NotNil(t, result)
	assert.Equal(t, uint32(1), *result)
}

func TestPointerUInt64(t *testing.T) {
	result := PointerUInt64(uint64(1))
	assert.NotNil(t, result)
	assert.Equal(t, uint64(1), *result)
}

func TestPointerBoolean(t *testing.T) {
	result := PointerBoolean(true)
	assert.NotNil(t, result)
	assert.True(t, *result)
}

func TestString(t *testing.T) {
	testStr := "abc"
	result := String(&testStr)
	assert.NotNil(t, result)
	assert.Len(t, result, 3)

	result = String(nil)
	assert.Len(t, result, 0)
}

func TestInt(t *testing.T) {
	def := 1
	value := 123
	result := Int(&value, def)
	assert.Equal(t, value, result)

	result = Int(nil, def)
	assert.Equal(t, def, result)
}

func TestInt8(t *testing.T) {
	def := int8(1)
	value := int8(123)
	result := Int8(&value, def)
	assert.Equal(t, value, result)

	result = Int8(nil, def)
	assert.Equal(t, def, result)
}

func TestInt16(t *testing.T) {
	def := int16(1)
	value := int16(123)
	result := Int16(&value, def)
	assert.Equal(t, value, result)

	result = Int16(nil, def)
	assert.Equal(t, def, result)
}

func TestInt32(t *testing.T) {
	def := int32(1)
	value := int32(123)
	result := Int32(&value, def)
	assert.Equal(t, value, result)

	result = Int32(nil, def)
	assert.Equal(t, def, result)
}

func TestInt64(t *testing.T) {
	def := int64(1)
	value := int64(123)
	result := Int64(&value, def)
	assert.Equal(t, value, result)

	result = Int64(nil, def)
	assert.Equal(t, def, result)
}

func TestUInt(t *testing.T) {
	def := uint(1)
	value := uint(123)
	result := UInt(&value, def)
	assert.Equal(t, value, result)

	result = UInt(nil, def)
	assert.Equal(t, def, result)
}

func TestUInt8(t *testing.T) {
	def := uint8(1)
	value := uint8(123)
	result := UInt8(&value, def)
	assert.Equal(t, value, result)

	result = UInt8(nil, def)
	assert.Equal(t, def, result)
}

func TestUInt16(t *testing.T) {
	def := uint16(1)
	value := uint16(123)
	result := UInt16(&value, def)
	assert.Equal(t, value, result)

	result = UInt16(nil, def)
	assert.Equal(t, def, result)
}

func TestUInt32(t *testing.T) {
	def := uint32(1)
	value := uint32(123)
	result := UInt32(&value, def)
	assert.Equal(t, value, result)

	result = UInt32(nil, def)
	assert.Equal(t, def, result)
}

func TestUInt64(t *testing.T) {
	def := uint64(1)
	value := uint64(123)
	result := UInt64(&value, def)
	assert.Equal(t, value, result)

	result = UInt64(nil, def)
	assert.Equal(t, def, result)
}

func TestFloat32(t *testing.T) {
	def := float32(1)
	value := float32(123)
	result := Float32(&value, def)
	assert.Equal(t, value, result)

	result = Float32(nil, def)
	assert.Equal(t, def, result)
}

func TestFloat64(t *testing.T) {
	def := float64(1)
	value := float64(123)
	result := Float64(&value, def)
	assert.Equal(t, value, result)

	result = Float64(nil, def)
	assert.Equal(t, def, result)
}

func TestBoolean(t *testing.T) {
	def := true
	value := false
	result := Boolean(&value, def)
	assert.Equal(t, value, result)

	result = Boolean(nil, def)
	assert.Equal(t, def, result)
}

// TestTimestamp_MarshalJSON test timestamp marshal json
func TestTimestamp_MarshalJSON(t *testing.T) {
	timeTest := Timestamp(time.Now())
	b, err := timeTest.MarshalJSON()
	assert.Nil(t, err)
	assert.NotNil(t, b)
}

// TestTimestamp_UnmarshalJSON test timestamp unmarshal json
func TestTimestamp_UnmarshalJSON(t *testing.T) {
	timeTest := Timestamp(time.Now())
	b, err := timeTest.MarshalJSON()
	assert.Nil(t, err)
	assert.NotNil(t, b)

	err = timeTest.UnmarshalJSON(b)
	assert.Nil(t, err)
}

// TestTimestamp_Value test timestamp value
func TestTimestamp_Value(t *testing.T) {
	timeTest := Timestamp(time.Now())
	driver, err := timeTest.Value()
	assert.NotNil(t, driver)
	assert.Nil(t, err)
}

func TestTimestamp_Scan(t *testing.T) {
	timeTest := Timestamp(time.Now())
	err := timeTest.Scan(time.Now())
	assert.Nil(t, err)
}
