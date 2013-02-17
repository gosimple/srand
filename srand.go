// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package srand implements simple pseudo-random generators.
package srand

import (
	"math/rand"
	"time"
)

//=============================================================================

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// IntFT returns a random non-negative number from min to max
// (0 <= min <= nr <= max).
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

// IntFTSample returns a slice of random non-negative numbers from min to max.
func IntFTSample(min, max, amount int) (sample []int) {
	for i := 0; i < amount; i++ {
		sample = append(sample, IntFT(min, max))
	}
	return sample
}

// IntBT returns a random number between min to max (0 <= min < nr < max).
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
