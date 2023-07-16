package reg

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLookahead(t *testing.T) {
	re := regexp.MustCompile(`\w+(?=\d)`)

	str := "hello123 world456"

	// 查找匹配的结果
	matches := re.FindAllString(str, -1)

	// 输出匹配的结果
	for _, match := range matches {
		fmt.Println(match)
	}
}
