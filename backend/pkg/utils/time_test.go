package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestTimestamp_MarshalJSON test timestamp marshal json
func TestTimestampMilliSeconds_MarshalJSON(t *testing.T) {
	timeTest := TimestampMilliseconds(time.Now())
	b, err := timeTest.MarshalJSON()
	assert.Nil(t, err)
	assert.NotNil(t, b)
}

// TestTimestamp_UnmarshalJSON test timestamp unmarshal json
func TestTimestampMilliSeconds_UnmarshalJSON(t *testing.T) {
	_now := time.Now()
	timeTest := TimestampMilliseconds(_now)
	timeUnmarshal := TimestampMilliseconds(time.Time{})

	b, err := timeTest.MarshalJSON()
	assert.Nil(t, err)
	assert.NotNil(t, b)

	err = timeUnmarshal.UnmarshalJSON(b)
	assert.Nil(t, err)
}

// TestTimestamp_Value test timestamp value
func TestTimestampMilliSeconds_Value(t *testing.T) {
	timeTest := TimestampMilliseconds(time.Now())
	driver, err := timeTest.Value()
	assert.NotNil(t, driver)
	assert.Nil(t, err)
}

func TestTimestampMilliSeconds_Scan(t *testing.T) {
	timeTest := TimestampMilliseconds(time.Now())
	err := timeTest.Scan(time.Now())
	assert.Nil(t, err)
}
