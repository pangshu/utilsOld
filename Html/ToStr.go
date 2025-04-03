package Html

import (
	"regexp"
	"strings"
)

func (*Html) ToStr(html string) string {
	re := regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllStringFunc(html, strings.ToLower)

	//remove STYLE
	re = regexp.MustCompile(`\<style[\S\s]+?\</style\>`)
	html = re.ReplaceAllString(html, "")

	//remove SCRIPT
	re = regexp.MustCompile(`\<script[\S\s]+?\</script\>`)
	html = re.ReplaceAllString(html, "")

	re = regexp.MustCompile(`\<[\S\s]+?\>`)
	html = re.ReplaceAllString(html, "\n")

	re = regexp.MustCompile(`\s{2,}`)
	html = re.ReplaceAllString(html, "\n")

	return strings.TrimSpace(html)
}
