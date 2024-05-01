package utils

import (
	"strings"
	"time"
)

// StringEmpty empty string
const StringEmpty = ""

// CurrentTimePointer return time.Time pointer
func CurrentTimePointer() *time.Time {
	currentTime := time.Now()
	currentDate := time.Date(
		currentTime.Year(),
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
		currentTime.Nanosecond(),
		time.UTC)
	return &currentDate
}

// StringToPointer get pointer from value
func StringToPointer(s string) *string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil
	}
	return &s
}

// PointerFloat64 get pointer from value
func PointerFloat64(v float64) *float64 {
	return &v
}

// PointerFloat32 get pointer from value
func PointerFloat32(v float32) *float32 {
	return &v
}

// PointerInt get pointer from value
func PointerInt(v int) *int {
	return &v
}

// PointerInt8 get pointer from value
func PointerInt8(v int8) *int8 {
	return &v
}

// PointerInt16 get pointer from value
func PointerInt16(v int16) *int16 {
	return &v
}

// PointerInt32 get pointer from value
func PointerInt32(v int32) *int32 {
	return &v
}

// PointerInt64 get pointer from value
func PointerInt64(v int64) *int64 {
	return &v
}

// PointerUInt get pointer from value
func PointerUInt(v uint) *uint {
	return &v
}

// PointerUInt8 get pointer from value
func PointerUInt8(v uint8) *uint8 {
	return &v
}

// PointerUInt16 get pointer from value
func PointerUInt16(v uint16) *uint16 {
	return &v
}

// PointerUInt32 get pointer from value
func PointerUInt32(v uint32) *uint32 {
	return &v
}

// PointerUInt64 get pointer from value
func PointerUInt64(v uint64) *uint64 {
	return &v
}

// PointerBoolean get pointer from value
func PointerBoolean(b bool) *bool {
	return &b
}

// String convert pointer to value
func String(s *string) string {
	if s == nil {
		return StringEmpty
	}
	return *s
}

// Int convert pointer to value
func Int(i *int, def int) int {
	if i == nil {
		return def
	}
	return *i
}

// Int8 convert pointer to value
func Int8(i *int8, def int8) int8 {
	if i == nil {
		return def
	}
	return *i
}

// Int16 convert pointer to value
func Int16(i *int16, def int16) int16 {
	if i == nil {
		return def
	}
	return *i
}

// Int32 convert pointer to value
func Int32(i *int32, def int32) int32 {
	if i == nil {
		return def
	}
	return *i
}

// Int64 convert pointer to value
func Int64(i *int64, def int64) int64 {
	if i == nil {
		return def
	}
	return *i
}

// UInt convert pointer to value
func UInt(i *uint, def uint) uint {
	if i == nil {
		return def
	}
	return *i
}

// UInt8 convert pointer to value
func UInt8(i *uint8, def uint8) uint8 {
	if i == nil {
		return def
	}
	return *i
}

// UInt16 convert pointer to value
func UInt16(i *uint16, def uint16) uint16 {
	if i == nil {
		return def
	}
	return *i
}

// UInt32 convert pointer to value
func UInt32(i *uint32, def uint32) uint32 {
	if i == nil {
		return def
	}
	return *i
}

// UInt64 convert pointer to value
func UInt64(i *uint64, def uint64) uint64 {
	if i == nil {
		return def
	}
	return *i
}

// Float32 convert pointer to value
func Float32(f *float32, def float32) float32 {
	if f == nil {
		return def
	}
	return *f
}

// Float64 convert pointer to value
func Float64(f *float64, def float64) float64 {
	if f == nil {
		return def
	}
	return *f
}

// Boolean convert pointer to value
func Boolean(b *bool, def bool) bool {
	if b == nil {
		return def
	}
	return *b
}
