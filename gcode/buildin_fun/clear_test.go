package buildin_fun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClear(t *testing.T) {
	assert_ := assert.New(t)
	{ //slice, just zero value. the len and capacity do not change
		as := [2]int{1, 2}
		old_len := len(as)
		clear(as[:])
		assert_.Equal(old_len, len(as))
		assert_.Equal(as, [2]int{0, 0})

		sl := make([]int, 2, 3)
		sl[0] = 1
		sl[1] = 2
		old_len = len(sl)
		old_cap := cap(sl)
		clear(sl)
		assert_.Equal(old_len, len(sl))
		assert_.Equal(old_cap, cap(sl))
		assert_.Equal(sl, []int{0, 0})
	}

	{ //map,
		data := map[int]int{1: 1, 2: 2}
		old_len := len(data)
		clear(data)
		assert_.Equal(old_len, 2)
		assert_.Equal(0, len(data))
		assert_.Equal(data, map[int]int{})
	}
}
