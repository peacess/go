package gcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEq(t *testing.T){
	a := [2]int{1,3}
	a2 := [2]int{2,3}
	if a == a2 {
		assert.Equal(t, a,a2)
	}

	b := []int{1,3}
	b2 := []int{2,3}
	assert.NotEqual(t, b,b2)// 不能使用 b == b2，不支持slice == slice的语法
}
