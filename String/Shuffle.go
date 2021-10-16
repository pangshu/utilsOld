package String

import (
	"math/rand"
	"time"
)

// Shuffle 随机打乱字符串.
func (*String)Shuffle(str string) string {
	if str == "" {
		return str
	}
	rand.Seed(time.Now().UnixNano())
	runes := []rune(str)
	index := 0

	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)
		if i != index {
			runes[i], runes[index] = runes[index], runes[i]
		}
	}

	return string(runes)
}
