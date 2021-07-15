// # Pythagorean Triplet

// A Pythagorean triplet is a set of three natural numbers, {a, b, c}, for
// which,

// ```text
// a**2 + b**2 = c**2
// ```

// and such that,

// ```text
// a < b < c
// ```

// For example,

// ```text
// 3**2 + 4**2 = 9 + 16 = 25 = 5**2.
// ```

// Given an input integer N, find all Pythagorean triplets for which `a + b + c = N`.

// For example, with N = 1000, there is exactly one Pythagorean triplet for which `a + b + c = 1000`: `{200, 375, 425}`.

package pythagorean

type Triplet [3]int

func pyth(a, b, c int) bool {
	return a*a+b*b == c*c
}

func Range(min, max int) (ts []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if pyth(a, b, c) {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

func Sum(sum int) (ts []Triplet) {
	max := sum / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := sum - a - b; pyth(a, b, c) {
				ts = append(ts, Triplet{a, b, c})
			}
		}
	}
	return
}
