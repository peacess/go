package forv122

import (
	"fmt"
	"testing"
)

func TestForV(t *testing.T) {
	for i := 0; i < 6; i++ {
		fmt.Printf("%p\n", &i)
		fmt.Println(i)
		i = i + 1
	}
}
