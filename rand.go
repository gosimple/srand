// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Choosing randomness simplified
package rand

import (
	"math/rand"
	"time"
)

//=============================================================================

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// IntFT returns a random number from min to max (0 <= min <= nr <= max).
// Will return 0 if max < 0
func IntFT(min, max int) int {
	var choice int

	if max < 0 {
		return 0
	}
	if min < 0 {
		min = 0
	}

	if min < max {
		choice = min + rand.Intn(max-min+1)
	} else {
		choice = min
	}
	return choice
}

// IntBT returns a random number between min to max (min < nr < max).
// Will return 0 if max < 0
func IntBT(min, max int) int {
	min += 1
	var choice int

	if max < 0 {
		return 0
	}
	if min < 0 {
		min = 0
	}

	if min < max {
		choice = min + rand.Intn(max-min)
	} else {
		choice = min
	}
	return choice
}
