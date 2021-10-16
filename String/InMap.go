package String

// 数组中是否包含
func (*String)InMap(str string, strMap []string) bool {
	for _, v := range strMap {
		if v == str {
			return true
		}
	}

	return false
}
