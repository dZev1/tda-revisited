package solution

func PowLog(n, k int) int {
	if k == 0 {
		return 1
	}
	if k % 2 == 0 {
		return PowLog(n * n, k / 2)
	}
	return n * PowLog(n * n, (k - 1) / 2)
}

// T(n) = a * T(n/b) + f(n)

// log_b(a) = k

// f(n) == O(n^c) c < k => T(n) == Θ(n^k)

// f(n) == Θ(n^k) => T(n) == Θ(n^k log n)
// f(n) == Θ(n^k * log^i(n)) => T(n) == Θ(n^k * log^(i + 1)(n))

// ((f(n) == Ω(n^c) c > k) ∧ (a * f(n/b) <= k f(n))) => T(n) == Θ(f(n))


// T(k) = 1*T(k/2) + Θ(1)
// t = log_2(1) == 0 <=> 2^0 == 1
// f(k) == Θ(1) == Θ(k^t) == Θ(k^0)
// => T(n) == Θ(1 * log k) == Θ(log k)