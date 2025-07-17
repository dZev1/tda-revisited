package solution

func multiply(X, Y [4][4]int) [4][4]int {
		var result [4][4]int
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				for k := 0; k < 4; k++ {
					result[i][j] += X[i][k] * Y[k][j]
				}
			}
		}
		return result
}
func sum(X,Y [4][4]int) [4][4]int {
	var result [4][4]int
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = X[i][j] + Y[i][j]
		}
	}
	return result
}
func potencia(A [4][4]int, n int) [4][4]int {
	var res [4][4]int
	for i := 0; i < 4; i++ {
		res[i][i] = 1
	}

	// Exponenciación rápida
	for n > 0 {
		if n%2 == 1 {
			res = multiply(res, A)
		}
		A = multiply(A, A)
		n /= 2
	}
	return res
}

// n = 2^k , k >= 1
// A = 4x4 matrix

func SumaDePotencia(A [4][4]int, n int) [4][4]int {
	if n == 1 {
		return A
	}
	m := n / 2
	Asq := potencia(A, m)
	sumM := SumaDePotencia(A, m)
	return sum(sumM, multiply(Asq, sumM))
}
