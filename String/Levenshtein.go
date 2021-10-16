package String

// Levenshtein 计算两个字符串之间的编辑距离.
func (*String)Levenshtein(a, b *string) int {
	la := len(*a)
	lb := len(*b)

	if *a == *b {
		return 0
	} else if la > 255 || lb > 255 {
		return -1
	}

	d := make([]int, la+1)
	var lastdiag, olddiag, temp int
	for i := 1; i <= la; i++ {
		d[i] = i
	}
	for i := 1; i <= lb; i++ {
		d[0] = i
		lastdiag = i - 1
		for j := 1; j <= la; j++ {
			olddiag = d[j]
			min := d[j] + 1
			if (d[j-1] + 1) < min {
				min = d[j-1] + 1
			}
			if (*a)[j-1] == (*b)[i-1] {
				temp = 0
			} else {
				temp = 1
			}
			if (lastdiag + temp) < min {
				min = lastdiag + temp
			}
			d[j] = min
			lastdiag = olddiag
		}
	}
	return d[la]
}

