package reg

import (
	"fmt"
	"regexp"
	"testing"
)

func TestLookahead(t *testing.T) {
	t.Run(`go reg 不支持 \w+(?=\d) `, func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				// this is ok
			} else {
				t.Errorf("Expected panic, but no panic occurred")
			}
		}()

		re := regexp.MustCompile(`\w+(?=\d)`)
		str := "hello123 world456"
		// 查找匹配的结果
		matches := re.FindAllString(str, -1)
		// 输出匹配的结果
		for _, match := range matches {
			fmt.Println(match)
		}
	})
	//

}
