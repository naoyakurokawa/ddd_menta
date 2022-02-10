package util

import "strconv"

func CastUint(str string) (uint16, error) {
	uint, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0, err
	}
	return uint16(uint), nil
}
