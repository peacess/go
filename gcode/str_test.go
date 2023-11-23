package gcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestType(t *testing.T) {
	{
		c1 := '1' // c1是，int32类型
		c2 := '中' // c2是什么类型, int32
		var c4 rune = '中'
		_ = c4

		fmt.Println("c1 type: ", reflect.TypeOf(c1)) // int32
		fmt.Println("c2 type: ", reflect.TypeOf(c2)) // int32

		fmt.Println(c1, c2)           // 49 20013
		fmt.Printf("%c,%c\n", c1, c2) // 1,中

		// 总结：
		// 1,go中没有rune类型，rune只是int32的别名，所以单个字符的类型是int32（rune是它的别名）
		// 2, 直接输出单个字符或rune类型，会以int32输出。如果要输出字符使用 %c
		// 3, go中没有字符类型，它由字符的整数编码代替。有string类型，它是uint8的数组（官方原文：string is the set of all strings of 8-bit bytes）
		fmt.Printf("%T\n", c4) // int32 类型，因为rune是别名，所以不会输出rune
	}

	{ // find the type of variable x
		c3 := '中'
		fmt.Println("c3 type: ", reflect.TypeOf(c3))         // int32
		fmt.Println("c3 type: ", reflect.ValueOf(c3).Kind()) // int32
		fmt.Printf("c3 type:  %T\n", c3)                     // int32
	}

	{
		s := "中间abc"
		c := s[0]
		fmt.Println(c, " type: ", reflect.TypeOf(c)) // 228, uint8
		fmt.Printf("%c\n\n", c)                      // not '中'

		// c is int
		fmt.Println("for c := range s")
		for c := range s { //c is the index of rune
			fmt.Println(c, ":", reflect.TypeOf(c))
		}
		for index, c := range s { //c is the index of rune
			fmt.Printf("%d: %c, ", index, c)
		}
		fmt.Println("end")

		//s[i] is uint8
		fmt.Println("for i := 0; i < len(s); i++")
		for i := 0; i < len(s); i++ {
			fmt.Println(s[i], ":", reflect.TypeOf(s[i]))
		}
	}
	{
		// how to get rune of in string

	}
}

func TestDoString(t *testing.T) {
	s := "中间的abc"
	c := s[0]
	fmt.Println(c) // 228

}
