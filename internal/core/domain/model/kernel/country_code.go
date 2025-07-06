package kernel

import (
	"fmt"
	"regexp"
	"strings"
)

var codePattern = regexp.MustCompile(`^[A-Z]{3}$`)

type CountryCode struct {
	value string
}

func NewCountryCode(code string) (CountryCode, error) {
	code = strings.ToUpper(strings.TrimSpace(code))
	if !codePattern.MatchString(code) {
		return CountryCode{}, fmt.Errorf("invalid country code: %s", code)
	}
	return CountryCode{value: code}, nil
}

func (c CountryCode) Value() string {
	return c.value
}

func (c CountryCode) Equals(other CountryCode) bool {
	return c.value == other.value
}
