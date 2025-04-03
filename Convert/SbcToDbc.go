package Convert

import (
	"golang.org/x/text/width"
)

// SbcToDbc 全角转半角.
func (*Convert)SbcToDbc(s string) string {
	return width.Narrow.String(s)
}
