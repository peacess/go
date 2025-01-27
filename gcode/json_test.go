package gcode

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type JsonData struct {
	Name string `json:"name"`
}

// 只有指针实现， 当JsonData不是指针时，并没有实现MarshalJSON
// 也就是：
// json.Marshal(d)时，如果d不是指针时，就不会调用当前实现的方法，而是json库自己的默认实现
func (c *JsonData) MarshalJSON() (bytes []byte, err error) {
	return json.Marshal(c.Name)
}

func (c *JsonData) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &c.Name)
}

func TestJsonData(t *testing.T) {
	{
		d := JsonData{Name: "test"}
		// 调用到自己实现的“MarshalJSON”
		// the bytes is "test"
		bytes, _ := json.Marshal(&d)
		// 不会调用到自己实现的“MarshalJSON”
		// the bytes is "{"name":"test"}"
		bytes2, _ := json.Marshal(d)
		assert.NotEqual(t, len(bytes), len(bytes2))
	}

	{
		d := JsonData{Name: "test"}

		bytes, err := json.Marshal(d)
		assert.Equal(t, nil, err)

		var d2 JsonData
		err = json.Unmarshal(bytes, &d2)
		assert.NotEqual(t, nil, err) //json: cannot unmarshal object into Go value of type string
		// the parameter d2 is not a pointer
		err = json.Unmarshal(bytes, d2)
		assert.NotEqual(t, nil, err) //json: Unmarshal(non-pointer gcode.JsonData)
	}
	{
		d := JsonData{Name: "test"}
		bytes, err := json.Marshal(&d)
		assert.Equal(t, nil, err)

		var d2 JsonData
		err = json.Unmarshal(bytes, &d2)
		assert.Equal(t, nil, err)
		assert.Equal(t, d, d2)

		// the parameter d2 is not a pointer
		err = json.Unmarshal(bytes, d2)
		assert.NotEqual(t, nil, err) //json: Unmarshal(non-pointer gcode.JsonData)
	}

}

type JsonData2 struct {
	Name string `json:"name"`
}

// non-pointer and pointer can call this method(both make copy)
func (c JsonData2) MarshalJSON() (bytes []byte, err error) {
	return json.Marshal(c.Name)
}

// non-pointer and pointer can call this method(both make copy)
// so UnmarshalJSON can't return value
func (c JsonData2) UnmarshalJSON(bytes []byte) error {
	return json.Unmarshal(bytes, &c.Name)
}

func TestJsonData2(t *testing.T) {
	{
		d := JsonData2{Name: "test"}
		// 调用到自己实现的“MarshalJSON”
		// the bytes is "test"
		bytes, _ := json.Marshal(&d)
		// 调用到自己实现的“MarshalJSON”
		// the bytes is "test"
		bytes2, _ := json.Marshal(d)
		assert.Equal(t, string(bytes), string(bytes2))
	}

	{
		d := JsonData2{Name: "test"}

		bytes, err := json.Marshal(d)
		assert.Equal(t, nil, err)

		var d2 JsonData2
		err = json.Unmarshal(bytes, &d2)
		assert.Equal(t, nil, err)
		assert.NotEqual(t, d.Name, d2.Name)
		err = json.Unmarshal(bytes, d2)
		assert.NotEqual(t, nil, err) ////json: Unmarshal(non-pointer gcode.JsonData)
	}

}
