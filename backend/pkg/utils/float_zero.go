package utils

import "strconv"

type FloatKeepZero float64

func (f FloatKeepZero) MarshalJSON() ([]byte, error) {
	if float64(f) == float64(int(f)) {
		return []byte(strconv.FormatFloat(float64(f), 'f', 1, 64)), nil
	}
	return []byte(strconv.FormatFloat(float64(f), 'f', -1, 64)), nil
}
