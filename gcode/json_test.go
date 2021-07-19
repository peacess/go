package gcode

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

type JsonData struct {
	name string
}

func (c *JsonData) MarshalJSON() (bytes []byte, err error) {
	return json.Marshal(c.name)
}

func (c *JsonData) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &c.name)
}

func TestJsonData(t *testing.T) {
	d := JsonData{name: "test"}
	bytes, err := json.Marshal(d)
	assert.Equal(t, nil, err)
	var d2 JsonData
	err = json.Unmarshal(bytes, &d2)
	assert.Equal(t, nil, err) //has error，为什么？
	assert.Equal(t, d, d2)    //not eq， 有两种方法，可以更正结果，一种是增加一个符号“&”,一种是删除一个符号“*”
}
