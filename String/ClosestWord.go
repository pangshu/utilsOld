package String

// ClosestWord 获取与原字符串相似度最高的字符串,以及它们的编辑距离.
// word为原字符串,searchs为待查找的字符串数组.
func (toolString *String)ClosestWord(word string, searchs []string) (string, int) {
	distance := 10000
	res := ""
	for _, search := range searchs {
		newVal := toolString.Levenshtein(&word, &search)
		if newVal == 0 {
			distance = 0
			res = search
			break
		}

		if newVal < distance {
			distance = newVal
			res = search
		}
	}

	return res, distance
}

