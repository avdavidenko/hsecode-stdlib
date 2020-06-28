package strings

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func backtrack(matrix [][]int, s1, s2 string, i, j int) string {
	if i == 0 || j == 0 {
		return ""
	}
	if s1[i-1] == s2[j-1] {
		return backtrack(matrix, s1, s2, i-1, j-1) + string(s1[i-1])
	}

	if matrix[i][j-1] > matrix[i-1][j] {
		return backtrack(matrix, s1, s2, i, j-1)
	}

	return backtrack(matrix, s1, s2, i-1, j)
}

func lcs(s1, s2 string) string {
	if len(s1) == 0 || len(s2) == 0 {
		return ""
	}

	l1 := len(s1)
	matrix := make([][]int, l1+1)
	l2 := len(s2)
	for i := range matrix {
		matrix[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			if s1[i-1] == s2[j-1] {
				matrix[i][j] = matrix[i-1][j-1] + 1
			} else {
				matrix[i][j] = max(matrix[i][j-1], matrix[i-1][j])
			}
		}
	}

	return backtrack(matrix, s1, s2, l1, l2)
}

func LCS(s1, s2 string) string {
	prefix := 0
	minLen := min(len(s1), len(s2))
	for prefix < minLen && s1[prefix] == s2[prefix] {
		prefix++
	}

	prefixStr := s1[:prefix]
	minLen -= prefix
	if prefix > 0 {
		s1 = s1[prefix:]
		s2 = s2[prefix:]
	}

	suffix := 0
	for suffix < minLen && s1[len(s1)-1-suffix] == s2[len(s2)-1-suffix] {
		suffix++
	}

	suffixStr := s1[len(s1)-suffix:]
	if suffix > 0 {
		s1 = s1[:len(s1)-suffix]
		s2 = s2[:len(s2)-suffix]
	}

	return prefixStr + lcs(s1, s2) + suffixStr
}
