package gcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceClone(t *testing.T) {
	data := []int{1}
	//clone 一个slice的最高效方法， data为nil时也可以正确工作
	clone1 := append(data[:0:0], data...)
	clone2 := append([]int(nil), data...)
	//不要使用下面的方法，他们都不会复制
	clone3Not := append(data[:0:len(data)], data...)
	clone4Not := append(data[:0], data...)
	//先 make后再 copy， 且语句更多，且需要特别处理nil情况，所以不建议使用

	assert.NotSame(t, &data[0], &clone1[0])
	assert.Equal(t, data, clone1)
	assert.NotSame(t, &data[0], &clone2[0])
	assert.Equal(t, data, clone2)

	assert.Same(t, &data[0], &clone3Not[0])
	assert.Same(t, &data[0], &clone4Not[0])
}
func TestSliceMergeClone(t *testing.T) {
	data1 := []int{1}
	data2 := []int{4}

	//合并两个slice，建议方法
	merge := make([]int, 0, len(data1)+len(data2))
	merge = append(merge, data1...)
	merge = append(merge, data2...)
	assert.Equal(t, data1, merge[0:len(data1)])
	assert.Equal(t, data2, merge[len(data1):])

	//合并两个slice，高效方法，由于代码比较多所以不建议使用，如果把MergeClone做成库，那么应该使用下面的方法
	merge = nil
	data1Len, data2Len := len(data1), len(data2)
	switch data1Len + data2Len {
	case data2Len:
		merge = append(data2[:0:0], data2...)
	case data1Len:
		merge = append(data1[:0:0], data1...)
	default:
		merge = append(data1, data2...)
	}
	assert.Equal(t, data1, merge[0:len(data1)])
	assert.Equal(t, data2, merge[len(data1):])
}
