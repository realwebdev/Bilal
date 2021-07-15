/*# Secret Handshake

> There are 10 types of people in the world: Those who understand
> binary, and those who don't.

You and your fellow cohort of those in the "know" when it comes to
binary decide to come up with a secret "handshake".

```text
1 = wink
10 = double blink
100 = close your eyes
1000 = jump


10000 = Reverse the order of the operations in the secret handshake.
```

Given a decimal number, convert it to the appropriate sequence of events for a secret handshake.

Here's a couple of examples:

Given the input 3, the function would return the array
["wink", "double blink"] because 3 is 11 in binary.

Given the input 19, the function would return the array
["double blink", "wink"] because 19 is 10011 in binary.
Notice that the addition of 16 (10000 in binary)
has caused the array to be reversed.
*/
package secret

var signals = []string{"wink", "double blink", "close your eyes", "jump"}

func Handshake(code uint) (h []string) {
	switch {
	case code < 1:
	case code&16 == 0:
		for _, s := range signals {
			if code&1 != 0 {
				h = append(h, s)
			}
			code >>= 1
		}

	default:
		for i := 3; i >= 0; i-- {
			if code&8 != 0 {
				h = append(h, signals[i])
			}
			code <<= 1
		}
	}
	return
}
