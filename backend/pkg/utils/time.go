package utils

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const (
	DateTimeFormatddMMyyHHmmss         = "02/01/2006 15:04:05"
	DateFormatddMMyy                   = "02/01/2006"
	TimesheetDateFormat                = "2006-01-02"
	TimesheetYearFormat                = "2006"
	TimesheetMonthFormat               = "01"
	OpsTimesheetDateFormat             = "02-01-2006"
	OpsTimesheetTimeFormat             = "15:04"
	DateTimeFormatyyMMddHHmmss         = "2006-01-02 15:04:05"
	DateTimeFormatGachCheoyyMMddHHmmss = "2006/01/02 15:04:05"
	DateTimeFormatddMMyyHHmm           = "02/01/2006 15:04"
)

// TimestampMilliseconds custom timestamp
type TimestampMilliseconds time.Time

// MarshalJSON rewrite MarshalJSON to return unix timestamp instead of string
func (t TimestampMilliseconds) MarshalJSON() ([]byte, error) {
	ts := time.Time(t).Unix() * 1000
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

// UnmarshalJSON rewrite UnmarshalJSON to handle unix timestamp from request
func (t *TimestampMilliseconds) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*t = TimestampMilliseconds(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}

// Value get value from timestamp
func (t TimestampMilliseconds) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan rewire scan for support sql driver
func (t *TimestampMilliseconds) Scan(value interface{}) error {
	*t = TimestampMilliseconds(value.(time.Time))
	return nil
}

func (t *TimestampMilliseconds) Format(format string) string {
	value, err := t.Value()
	if err != nil {
		return ""
	}
	return value.(time.Time).Format(format)
}

// Timestamp custom timestamp
type Timestamp time.Time

// MarshalJSON rewrite MarshalJSON to return unix timestamp instead of string
func (t Timestamp) MarshalJSON() ([]byte, error) {
	ts := time.Time(t).Unix()
	stamp := fmt.Sprint(ts)

	return []byte(stamp), nil
}

// Return unix timestamp instead of string
func (t Timestamp) GetTimeUnix() int64 {
	return time.Time(t).Unix()
}

// Return unix timestamp instead of string
func (t Timestamp) GetTime() time.Time {
	return time.Time(t)
}

// UnmarshalJSON rewrite UnmarshalJSON to handle unix timestamp from request
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))
	if err != nil {
		return err
	}

	*t = Timestamp(time.Unix(int64(ts), 0))

	return nil
}

// Value get value from timestamp
func (t Timestamp) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// Scan rewire scan for support sql driver
func (t *Timestamp) Scan(src interface{}) error {
	if val, ok := src.(time.Time); ok {
		*t = Timestamp(val)
	} else {
		return errors.New("time Scanner passed a non-time object")
	}

	return nil
}

func (t *Timestamp) IsZero() bool {
	value, _ := t.Value()
	return value.(time.Time).IsZero()
}

func (t *Timestamp) Format(format string, hideZero bool) string {
	value, _ := t.Value()
	if hideZero && value.(time.Time).IsZero() {
		return ""
	}
	return value.(time.Time).Format(format)
}
