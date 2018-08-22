package lintcode

func coinProblem(n int, m int) int { //lintcode 1546
	// Write your code here
	var res, n100, n50, n20, n10, n5, n2, n1 int
	res = n - m
	n100 = res / 100
	res -= 100 * n100
	n50 = res / 50
	res -= 50 * n50
	n20 = res / 20
	res -= 20 * n20
	n10 = res / 10
	res -= 10 * n10
	n5 = res / 5
	res -= 5 * n5
	n2 = res / 2
	res -= 2 * n2
	n1 = res / 1
	res -= 1 * n1
	return n100 + n50 + n20 + n10 + n5 + n2 + n1
}
