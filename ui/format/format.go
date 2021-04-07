// Package format provides displaying formatter functions
package format

import (
	"fmt"
	"strings"
)

// Currency formats monetary values and adds a dollar sign
func Currency(num float32) string {
	s := fmt.Sprintf("%.2f", num)
	c := "$"
	if strings.HasPrefix(s, "-") {
		s = strings.Replace(s, "-", "-"+c, 1)
	} else {
		s = c + s
	}
	return s
}
