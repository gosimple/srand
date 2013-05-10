// Copyright 2013 by Dobrosław Żybort. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package srand simplify choosing random values, rand helper.
package srand

import (
	"math/rand"
	"time"
)

//=============================================================================

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Int returns a non-negative random number from min to max
// (0 <= min <= nr <= max).
// Will return 0 if max < 0
func Int(min, max int) int {
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

// IntMany returns a slice of non-negative random numbers from min to max.
func IntMany(min, max, quantity int) (out []int) {
	for i := 0; i < quantity; i++ {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		out = append(out, Int(min, max))
	}
	return out
}

// IntSample returns a slice of non-negative random, unique numbers
// from min to max. If the quantity is bigger than max - min sample size
// will be the size of max - min.
func IntSample(min, max, quantity int) (sample []int) {
	if max < 0 {
		return []int{0}
	}
	if min < 0 {
		min = 0
	}
	maxQuantity := max - min
	if maxQuantity > quantity {
		maxQuantity = quantity
	}
	i := 0
	for {
		rand.Seed(time.Now().UTC().UnixNano() + int64(i))
		randNum := Int(min, max)
		if !containsInt(sample, randNum) {
			sample = append(sample, randNum)
			i++
		} else {
			continue
		}
		if i == maxQuantity {
			break
		}
	}
	return sample
}

func containsInt(list []int, elem int) bool {
	for _, t := range list {
		if t == elem {
			return true
		}
	}
	return false
}

//=============================================================================

var (
	// rfc1924
	BASE85 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~"

	// rfc4648 alphabets
	BASE16    = BASE85[:16]
	BASE32    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ234567"
	BASE32HEX = BASE85[:32]
	BASE64    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	BASE64URL = BASE64[:62] + "-_"

	// Only contain alphanumeric characters
	BASE62 = BASE85[:62]
)

// String returns a random string with length from minLength to maxLength.
// Use BASE62 alphabet by default.
// Will return empty string if maxLength < 0
func String(minLength, maxLength int) string {
	var length int

	if minLength < 0 {
		minLength = 0
	}
	if maxLength < 0 {
		return ""
	}

	length = Int(minLength, maxLength)

	buf := make([]byte, length)
	for i := 0; i < length; i++ {
		buf[i] = BASE62[rand.Intn(len(BASE62)-1)]
	}
	return string(buf)
}
