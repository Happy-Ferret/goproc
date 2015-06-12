// Copyright 2015 Giacomo Stelluti Scala. All rights reserved. See doc/License.md in the project root for license information.

package internal

import "strconv"

func atoiOr(s string, alt int) int {
	value, err := strconv.Atoi(s)
	if err == nil {
		return value
	}
	return alt
}
