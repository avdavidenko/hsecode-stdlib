package levenshtein

type Levenshtein struct {
	distance   int
	transcript string
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func New(src, dst string) *Levenshtein {
	prefix := 0
	minLen := min(len(src), len(dst))
	for prefix < minLen && src[prefix] == dst[prefix] {
		prefix++
	}

	minLen -= prefix
	prefixStr := strings.Repeat("M", prefix)
	if prefix > 0 {
		src = src[prefix:]
		dst = dst[prefix:]
	}

	suffix := 0
	for suffix < minLen && src[len(src)-1-suffix] == dst[len(dst)-1-suffix] {
		suffix++
	}

	suffixStr := strings.Repeat("M", suffix)
	if suffix > 0 {
		src = src[:len(src)-suffix]
		dst = dst[:len(dst)-suffix]
	}

	distance, transcript := makeTranscript(src, dst)
	return &Levenshtein{distance: distance, transcript: prefixStr + transcript + suffixStr}
}

func (ls *Levenshtein) Distance() int {
	return ls.distance
}

func (ls *Levenshtein) Transcript() string {
	return ls.transcript
}

func max(x, y int) int {
	if x < y {
		return y
	}

	return x
}

func makeTranscript(src, dst string) (int, string) {
	srcLen := len(src)
	dstLen := len(dst)

	fmt.Println(src, dst)
	if srcLen == 0 && dstLen == 0 {
		return 0, ""
	}

	matrix := make([][]int, srcLen+1)

	for i := 0; i <= srcLen; i++ {
		matrix[i] = make([]int, dstLen+1)
		matrix[i][0] = i
	}

	for j := 1; j <= dstLen; j++ {
		matrix[0][j] = j
	}

	for i := 1; i <= srcLen; i++ {
		for j := 1; j <= dstLen; j++ {
			subCost := 0
			if src[i-1] != dst[j-1] {
				subCost = 1
			}

			matrix[i][j] = min(min(matrix[i-1][j], matrix[i][j-1])+1, matrix[i-1][j-1]+subCost)
		}
	}

	fmt.Println(matrix)
	return matrix[srcLen][dstLen], backtraceTranscript(srcLen, dstLen, matrix)
}

func backtraceTranscript(i, j int, matrix [][]int) string {
	if i > 0 && matrix[i-1][j]+1 == matrix[i][j] {
		return backtraceTranscript(i-1, j, matrix) + "D"
	}

	if j > 0 && matrix[i][j-1]+1 == matrix[i][j] {
		return backtraceTranscript(i, j-1, matrix) + "I"
	}

	if i > 0 && j > 0 && matrix[i-1][j-1]+1 == matrix[i][j] {
		return backtraceTranscript(i-1, j-1, matrix) + "R"
	}

	if i > 0 && j > 0 && matrix[i-1][j-1] == matrix[i][j] {
		return backtraceTranscript(i-1, j-1, matrix) + "M"
	}

	return ""
}
