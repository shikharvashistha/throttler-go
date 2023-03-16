package utils

import (
	"strconv"
	"time"
)

func ParseTimestamp(str string) (*time.Time, error) {
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return nil, err
	}

	last_time := time.Unix(i, 0)
	return &last_time, nil
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
