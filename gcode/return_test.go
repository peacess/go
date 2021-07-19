package gcode

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReturn(t *testing.T) {
	var f = func() (d int, err error) {
		return
	}

	{
		d := 1
		err := errors.New("new err")
		d, err = f()
		assert.Equal(t, 0, d)
		assert.Equal(t, nil, err)
	}
	{
		d := 1
		//err := errors.New("new err")
		d, err := f()
		assert.Equal(t, 0, d)
		assert.Equal(t, nil, err)
	}

	{
		//d := 1
		err := errors.New("new err")
		d, err := f()
		assert.Equal(t, 0, d)
		assert.Equal(t, nil, err)
	}

}
