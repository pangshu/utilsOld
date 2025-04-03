package Convert

import (
	"golang.org/x/text/width"
)

// DbcToSbc 半角转全角.
func (*Convert)DbcToSbc(s string) string {
	return width.Widen.String(s)
}
