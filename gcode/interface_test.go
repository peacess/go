package gcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestSameEqual(t *testing.T) {
	a1 := 10
	a2 := 10
	p1 := &a1
	p2 := &a2
	assert.Equal(t, p1, p2)   //地址不同，但值相等，所以成立
	assert.NotSame(t, p1, p2) //比较指针是否相等，所以成立
}

func TestFat(t *testing.T) {
	var fat1 interface{} = nil // {interface{}} nil
	_ = fat1
	var i = 0
	var fat2 interface{} = i // {interface{}|int} 0
	_ = fat2
	var pi = &i
	var fat3 interface{} = pi //{interface{}|*int} 0
	_ = fat3
	pi = nil
	var fat4 interface{} = pi //{interface{}|*int} nil
	_ = fat4

	var fat5 interface{} = 10 //{interface{}|int} 0
	var fat6 = fat5           //{interface{}|int} 0
	_ = fat6
}

func TestInterfaceEq(t *testing.T) {
	var a1 interface{} = nil
	var point *int = nil
	var a2 interface{} = point
	assert.NotEqual(t, a1, a2)

	a1 = (*int)(nil)
	assert.Equal(t, a1, a2)
	if a1 == a2 {
		println("")
	}

	var i = 10
	point = &i //这一行并不影响a2的值
	assert.Equal(t, a1, a2)

	point = &i
	point2 := &i
	a2 = point
	a1 = point2
	assert.Equal(t, a1, a2)
	if a1 == a2 { //比较两个字段是否相等
		println("")
	}
	a2 = &point
	a1 = &point2
	assert.Equal(t, a1, a2) //point与point2的值相等，但是使用是 deepValueEqual，最终比较的是值相等
	if a1 == a2 {           //类型相等，但是地址不相同，所以不等
		println("")
	}

	{
		var i = 1
		var i2 = 1
		var pi = &i
		var pi2 = &i2
		var pi3 = &i
		fmt.Printf("pi == pi3: %v, pi == pi2: %v\n", pi == pi3, pi == pi2)
		//assert.Equal != "=="
	}
	{
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := ch1
		fmt.Printf("ch1 == ch2: %v, ch1 == ch3: %v\n", ch1 == ch2, ch1 == ch3)
	}
}

type Inter interface {
	Name() string
	SetName(name string)
}

type SubData struct {
	tName string
}

type Data struct {
	SubData
	C int
}

func (c SubData) Name() string {
	return c.tName
}

func (c Data) Name() string {
	return c.tName
}

func (c *Data) SetName(name string) {
	c.tName = name
}

func TestData(t *testing.T) {
	d := Data{SubData: SubData{tName: "123"}, C: 10}
	var inter Inter = &d
	var inter2 Inter = &d

	// layout of Itab known to compilers
	// allocated in non-garbage-collected memory
	// Needs to be in sync with
	// ../cmd/compile/internal/gc/reflect.go:/^func.dumptabs.
	type itab struct {
		inter *int
		_type *int
		hash  uint32 // copy of _type.hash. Used for type switches.
		_     [4]byte
		fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
	}

	type iface struct {
		tab  *itab
		data *Data
	}

	i1 := (*iface)(unsafe.Pointer(&inter))
	i2 := (*iface)(unsafe.Pointer(&inter2))

	d.C = 9

	//d.SetName("")
	d.Name()
	inter.Name()
	inter.SetName("")

	_ = i1
	_ = i2

}
