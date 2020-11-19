package luhn

import (
	"strings"
)

func Valid(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	sum := 0
	if len(s) < 2 {
		return false
	}
	double := len(s)%2 == 0
	for _, v := range s {
		digit := int(v - '0')
		if digit < 0 || digit > 9 {
			return false
		}
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	return sum%10 == 0
}
