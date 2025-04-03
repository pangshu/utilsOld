package Url

import (
	"net/http"
)

// IsUrlExists 检查URL是否存在.
func (toolUrl *Url)IsUrlExists(str string) bool {
	if !toolUrl.IsUrl(str) {
		return false
	}

	client := &http.Client{}
	resp, err := client.Head(str)
	if err != nil {
		return false
	} else if resp.StatusCode == 404 {
		return false
	}

	return true
}
