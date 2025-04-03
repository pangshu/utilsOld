package String

import (
	crand "crypto/rand"
	"fmt"
)

// UuidV4 获取UUID(Version4).
func (*String)UuidV4() (string, error) {
	buf := make([]byte, 16)
	_, err := crand.Read(buf)

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%12x",
		buf[0:4],
		buf[4:6],
		buf[6:8],
		buf[8:10],
		buf[10:16]), err
}
