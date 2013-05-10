srand - simple rand
===================

Package srand simplify choosing random values.

[Documentation online](http://godoc.org/bitbucket.org/gosimple/srand)

	package main

	import (
		"bitbucket.org/gosimple/srand"
		"fmt"
	)

	func main() {
		fmt.Println(srand.Int(10, 20))          // 12
		fmt.Println(srand.IntMany(10, 20, 3))   // [17 17 13]
		fmt.Println(srand.IntSample(10, 20, 3)) // [20 14 12]

		fmt.Println(srand.String(10, 15))                // aaYCKpwvRhsZz
		fmt.Println(srand.StringA(10, 15, "abc"))        // aaaabcbbbab
		fmt.Println(srand.StringA(10, 15, srand.BASE85)) // T$a~(#HTQ>7q
	}


### Requests or bugs?
<https://bitbucket.org/gosimple/srand/issues>

## Installation

	go get -u bitbucket.org/gosimple/srand

## License

The source files are distributed under the
[Mozilla Public License, version 2.0](http://mozilla.org/MPL/2.0/),
unless otherwise noted.
Please read the [FAQ](http://www.mozilla.org/MPL/2.0/FAQ.html)
if you have further questions regarding the license.
