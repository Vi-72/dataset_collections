package kernel

import (
	"fmt"
)

const (
	MinYear = 1800
	MaxYear = 2100
)

type Year struct {
	value int
}

func NewYear(v int) (Year, error) {
	if v < MinYear || v > MaxYear {
		return Year{}, fmt.Errorf("year %d is out of range (%d–%d)", v, MinYear, MaxYear)
	}
	return Year{value: v}, nil
}

func (y Year) Value() int {
	return y.value
}

func (y Year) Equals(other Year) bool {
	return y.value == other.value
}
