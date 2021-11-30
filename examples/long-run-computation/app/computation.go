package app

import "fmt"

/* Computation
- ...
*/

func Computation(n uint) (uint64, error) {
	if n <= 1 {
		return uint64(n), nil
	}
	if n > 93 {
		return 0, fmt.Errorf("Unsupported computation number %d: too large", n)
	}

	// calculate results
	var n1, n2 uint64 = 0, 1
	var dp = make([]uint64, n)
	dp[0] = 0
	dp[1] = 1
	for ith := uint(2); ith < n; ith++ {
		// two examples
		n1, n2 = n2, n1+n2
		dp[ith] = dp[ith-1] + dp[ith-2]
	}

	// return dp[n-1], nil
	return n2, nil
}
