package tesla

import (
	"fmt"
)

// reverseDigits reverse digits of n, and prints to std out.
func reverseDigits(N int) {
	enablePrint := N % 10

	for N > 0 {
		if enablePrint == 0 && N%10 != 0 {
			enablePrint = 1
		}

		if enablePrint == 1 {
			fmt.Print(N % 10)
		}
		N /= 10
	}
}
