package Html


// BrToNl 将br标签转换为换行符.
func (*Html)BrToNl(str string) string {
	// <br> , <br /> , <br/>
	// <BR> , <BR /> , <BR/>
	nlchar := []byte("\n")

	l := len(str)
	buf := make([]byte, 0, l) //prealloca

	for i := 0; i < l; i++ {
		switch str[i] {
		case 60: //<
			if l >= i+3 {
				/*
					b = 98
					B = 66
					r = 82
					R = 114
					SPACE = 32
					/ = 47
					> = 62
				*/

				if l >= i+3 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 62) { // <br> || <BR>
					buf = append(buf, nlchar...)
					i += 3
					continue
				}

				if l >= i+4 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 47 && str[i+4] == 62) { // <br/> || <BR/>
					buf = append(buf, nlchar...)
					i += 4
					continue
				}

				if l >= i+5 && ((str[i+1] == 98 || str[i+1] == 66) && (str[i+2] == 82 || str[i+2] == 114) && str[i+3] == 32 && str[i+4] == 47 && str[i+5] == 62) { // <br /> || <BR />
					buf = append(buf, nlchar...)
					i += 5
					continue
				}
			}
		default:
			buf = append(buf, str[i])
		}
	}

	return string(buf)
}
