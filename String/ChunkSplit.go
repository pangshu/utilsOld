package String


// ChunkSplit 将字符串分割成小块.str为要分割的字符,chunklen为分割的尺寸,end为行尾序列符号.
func (*String)ChunkSplit(str string, chunklen uint, end string) string {
	if end == "" {
		end = "\r\n"
	}
	runes, erunes := []rune(str), []rune(end)
	length := uint(len(runes))
	if length <= 1 || length < chunklen {
		return str + end
	}
	ns := make([]rune, 0, len(runes)+len(erunes))
	var i uint
	for i = 0; i < length; i += chunklen {
		if i+chunklen > length {
			ns = append(ns, runes[i:]...)
		} else {
			ns = append(ns, runes[i:i+chunklen]...)
		}
		ns = append(ns, erunes...)
	}
	return string(ns)
}

