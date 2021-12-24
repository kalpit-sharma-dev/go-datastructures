package recursion

/*
Problem Description

Implement pow(x, n) % d.
In other words, given x, n and d,

find (xn % d)

Note that remainders on division cannot be negative. In other words, make sure the answer you return is non negative.

Input : x = 2, n = 3, d = 3
Output : 2

2^3 % 3 = 8 % 3 = 2.
*/

func pow(A int, B int, C int) int {
	if A == 0 {
		return 0
	}
	if B == 0 {
		return 1
	}

	if A < 0 {
		return pow(A+C, B, C)
	}
	ans := 0
	temp := pow(A, B/2, C)
	if B%2 == 0 {
		ans = ((temp % C) * (temp % C)) % C
	} else {
		ans = ((((A % C) * (temp % C)) % C) * (temp % C)) % C
	}

	return ans
}
