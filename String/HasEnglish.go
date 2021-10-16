package String

// HasEnglish 是否含有英文字符,HasLetter的别名.
func (toolString *String)HasEnglish(str string) bool {
	return toolString.HasLetter(str)
}

