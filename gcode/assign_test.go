package gcode

import (
	"fmt"
	"reflect"
	"testing"
	"unicode/utf8"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

//赋值是编辑语言的基础，理解赋值就拿到打开这门语言的钥匙。
//函数调用时传递参数与函数返回值都是赋值（如果编译器作函数返回值优化时，返回值是没有赋值）
//在go语言中所有的变量都是“值”，赋值就是值的复制。
//指针的背后的类型为 int类型，对指针赋值相当于一个int的赋值
//array是一个值，在赋值时是整个数据全复制的
//slice背后的类型为 reflect.SliceHeader，所以对slice的赋值相当于结构体reflect.SliceHeader的赋值
//string背后的类型为reflect.StringHeader，所以对string的赋值相当于结构体reflect.StringHeader的赋值
//slice可以与nil比较， 而string与array都不能与nil比较
//下面是各种类型的大小，这个也可以来恒量它们赋值时的代价
//	        string	func	slice	map	channel	Interface{}	 *
//	sizeof	16   	8	    24	    8	8	    16	         8

// bool int int32 ...
func TestBoolIntegers(t *testing.T) {
	a := 10
	b := 60
	b = a //把a的内存复制到b的内存
	//对于bool, integers类型，赋值就是内存复制，完成后它们是相互独立的
	_ = b //这行代码只是让编译通过，因为go的编译器认为定义的变量没有使用是一个错误
}

// 指针类型
func TestPointer(t *testing.T) {
	a := 10
	pa := &a
	b := 60
	pb := &b

	*pa = 11 // 相当于 a = 11，这时pa这个指针地址没有改变，只是地址指向的值改变了
	pa = pb  // 相当于一个int类型的赋值。 把pb的地址赋值给 pa

	//指针类型相当于一个int, 它的赋值也与int类似
	//在指针前面加“＊”,表示取指针指向的“值”。这是对指向的“值”进行操作，并不会改指针自己的int值
}

type StructSample struct {
	Name  string
	Point int
}

func (c *StructSample) SetName(name string) {
	c.Name = name
}

func (u StructSample) SetPoint(point int) {
	u.Point = point
}
func TestStruct(t *testing.T) {
	assert := assert.New(t)
	{ //=
		u := StructSample{
			Name:  "test",
			Point: 6,
		}
		u2 := u
		u2.Point = 8
		assert.NotEqual(u.Point, u2.Point)
		//u 与 u2 是copy, 它们的内存地址不相同
	}

	{
		u := StructSample{
			Name:  "test",
			Point: 6,
		}
		u.SetName("new")
		assert.Equal("new", u.Name)

		old := u.Point
		u.SetPoint(8)
		assert.Equal(old, u.Point)

		p := &u
		old = p.Point
		p.SetPoint(8)
		assert.Equal(old, p.Point)

		p.SetName("new2")
		assert.Equal("new2", p.Name)

		//接收者为指针时， 不管调用者是否为指针，它们都是同一个实例
		//接收者不是指针时，不管理调用者是否为指针，它们都不是同一个实例（会产生一次struct的赋值，也就是一个新的副本）
	}

	{ //sizeof
		string_ := ""
		slice_ := make([]int32, 10)
		func_ := func() {}
		map_ := make(map[int]int32, 10)
		chan_ := make(chan [17]int32)
		var p_ uint8
		p := &p_

		var inter_ interface{} = p_

		assert.Equal(16, int(unsafe.Sizeof(string_)))
		assert.Equal(24, int(unsafe.Sizeof(slice_)))
		assert.Equal(8, int(unsafe.Sizeof(func_)))
		assert.Equal(8, int(unsafe.Sizeof(map_)))
		assert.Equal(8, int(unsafe.Sizeof(chan_)))
		assert.Equal(8, int(unsafe.Sizeof(p)))
		assert.Equal(16, int(unsafe.Sizeof(inter_)))
	}

	//汇总：
	//赋值是把所有的字段进行赋值（字段是什么类型就按照什么类型赋值），也就从一个内存复制到另外的内存
	//“指针”类型，赋值是把指针部分的数据，从一个内存复制到另外内存，而数据本向不变化
	//
	//值类型，boolean,数字类型，struct, array
	//	“指针”类型，string ,func,slice,map,channel,interface{},任何类型的指针（使用“*”定义的）
	//	        string	func	slice	map	channel	Interface{}	 *
	//	sizeof	16   	8	    24	    8	8	    16	         8
	//	注：这个是在64位系统输出的结果
}

func TestString(t *testing.T) {
	//string类型使用的是utf-8编码，不可变。下面看看utf-8编码对我们使用string有那些影响
	str := "中国"
	//str[1] 不等于 "国" // str[1]取出第二byte的值，经代码验证值为“184”类型为uint8，uint8就是byte类型

	// 如果要取第二个字符需要把它转换为rune（实际是一个int32类型）类型：
	rs := []rune(str) //rs[1]等于 "国"

	//在for rang中string会转换成rune不是byte
	for _, v := range str {
		fmt.Println(reflect.TypeOf(v))
		break
	}
	//这里打印出来的是int32类型，也就是rune类型

	//len(str)计算的是byte，不是字符个数，计算字符个数使用：
	assert.True(t, 2 != len(str))
	assert.Equal(t, 6, len(str))
	assert.Equal(t, 2, len([]rune(str)))
	assert.Equal(t, 2, utf8.RuneCountInString(str)) //建议使用这个方法
	//问题：如果要使用byte方式遍历 string，怎么处理
	//这里有两个方法，一是使用下标的方式来for
	for i := 0; i < len(str); i++ {
		_ = str[i] //这里是byte
	}
	//另一个方法是，把str转换为byte再再来循环
	bytes := []byte(str)
	for i, v := range bytes {
		_, _ = i, v
	}

	//汇总string的特点
	// go中的string是utf8编码
	// str[i]是byte类型
	// for range str时，使用的是rune，也就是字符，而不是byte
	// 计算字符个数建议使用utf8.RuneCountInString(str)
	// string不能与nil比较
	// string的默认值为 “”，而不是nil
	_ = rs

	//下面我们来看看，string的赋值，做了什么
	//string类型不是一个简单类型，它实际上是一个"胖指针"
	//type StringHeader struct {
	//	Data uintptr
	//	Len  int
	//}
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	_ = stringHeader
	//所以对于string的赋值相当于结构体StringHeader的赋值
}

func TestArraySlice(t *testing.T) {
	assert := assert.New(t)

	//数组分为两种，array 与 slice，array不变数组，slice是可变化数组（也叫动态数组）
	//他们还有一个区别， array是一个"值"(不可以与nil比较)，而slice是一个指针

	//初始化
	{
		a := [2]int32{}   //array, 长度为2，默认值为零
		a2 := [2]int32{0} //array, 长度为2，第一个指定为零，第二个没有指定，默认为零
		var a3 [2]int32   //array,长度为2，默认值为零
		//这三个变量的区别是什么？
		_, _, _ = a, a2, a3 //这一行只是为编译通过，没有实际功能
		assert.Equal([2]int32{0, 0}, a)
		assert.Equal([2]int32{0, 0}, a2)
		assert.Equal([2]int32{0, 0}, a3)
		//从测试来看它们的值是相等的，效果相同
	}
	{
		b := [2]int32{}      //array,
		b2 := [2]int32{1, 2} //array
		b3 := []int32{1, 2}  //slice，只要没有指定大小的，就是slice
		//这三个变量的区别是什么？
		_, _, _ = b, b2, b3 //这一行只是为编译通过，没有实际功能
		assert.Equal([2]int32{0, 0}, b)
		assert.Equal([2]int32{1, 2}, b2)
		assert.Equal([]int32{1, 2}, b3)

		assert.Equal(2, cap(b))
		assert.Equal(2, cap(b2))
		assert.Equal(2, cap(b3))
	}
	{
		c := make([]int32, 0)      //slice, 不为nil, len == cap == 0
		c2 := make([]int32, 0, 0)  //slice, 不为nil, len == cap == 0
		c3 := make([]int32, 0, 10) //slice, 不为nil, len = 0, cap = 10
		var c4 []int32             //slice, 为nil
		_, _, _ = c, c2, c3
		assert.True(nil == c4)
	}
	{ //创建零长度的slice
		d := []int32{}
		d2 := make([]int32, 0)

		assert.Equal(0, len(d))
		assert.Equal(0, len(d2))
	}

	{ //slice
		a := []int32{1, 2, 3}
		b := a
		assert.Equal(a, b)
		a[0] = 9
		assert.Equal(a, b)
		//slice在赋值时不会产生copy，也就是a b两个 slice的指向同一个内存，所以修改a,b的值也会变
		c := a[2:]
		c[0] = 8
		assert.Equal(a[2:], c)
		//从一个slice取出部分作为一个新slice时，他们仍然指向同一内存，所以修改 c后， a b c的值都发生变化了

		d := append(a, 1)
		d[0] = 1
		assert.NotEqual(a[0], d[0])
		//在使用append函数时，如果append后的元素大于 cap，那么会返回一个新的地址
		//此时 a 与 d指向不同内存，所以修改d后，a没有发生变化
		e := make([]int32, 1, 6) // 从a中取 [0,3)的数据生成一个cap为10的slice
		f := append(e, 2)
		f[0] = 7
		assert.Equal(e[0], f[0])
		//在使用append时，没有超出cap所以， append并没有分配新的内存，那么e f两个指向的内存相同，所以修改f后，e的值也根着变化
	}
	// slice 复制
	{
		a := []int32{1, 2}
		var b []int32
		b = append(b, a...) //这里不会分配新的内存
		assert.Equal(a, b)
		b[0] = 10
		assert.NotEqual(a[0], b[0])
		//a b指向不相同
	}
	//array slice =
	{
		a := [2]int32{1, 2}
		b := a
		assert.Equal(a, b)
		a[0] = 9
		assert.NotEqual(a, b)
		//array在赋值时会产生copy，也就是a b两个 array的内存不一样，所以修改a不会影响b

		c := []int32{1, 2}
		d := c
		assert.Equal(c, d)
		c[0] = 9
		assert.Equal(c[0], d[0])
		//slice在赋值后，它们都指向同一内存，所以修改一个，另外一个也会变化
		//slice与string在赋值上很像，那么slice的“指针”结构是如下
		header := reflect.SliceHeader{}
		assert.Equal(24, int(unsafe.Sizeof(header)))
		//那么slice的一次赋值会产生 24byte的内存复制，string是16
	}
}

func TestLoop(t *testing.T) {
	{
		users := []StructSample{{Name: "1", Point: 1}, {Name: "2", Point: 2}}
		for _, v := range users {
			v.Name = "test" //这里能修改到 users中的数据吗？ 不能因为 v 相当于一次 struct的赋值，它是一个copy
		}
		//上面这个循环的性能问题在那里？
		//下面是另外的实现
		for i := range len(users) {
			users[i].Name = "test2" //这里能修改到 users中的数据吗？,可以，users[i]这个操作不会产生copy，取下标也不是一个赋值
		}
		//注：在遍历时，一定要注意 “赋值”带来的副作用，如要struct比较大，是有性能问题的
	}
	{
		users2 := []*StructSample{{Name: "1", Point: 1}, {Name: "2", Point: 2}}
		for _, v := range users2 {
			v.Name = "test" //这里能修改到 users中的数据吗？ 可以，因为是指针，这里的“赋值”，相当于“ v = p”
		}
		//上面这个循环的性能问题在那里？, 这里产生了一个 int 类型的赋值，它对性能的影响非常小
		//下面是另外的实现
		for i := range len(users2) {
			users2[i].Name = "test2" //这里能修改到 users中的数据吗？ 可以
		}

		//users2中的两个循环的是否差不多，写代码时怎么选择?
		//对于下面这种数据类型，在写for时，应该选择那一种？
		strs := []string{}
		_ = strs
	}

	//总结
	//成员为“指针”类型来说， 两种循环都可以选择
	//成员为非指针类时，建议使用下标方式来使用循环

	//自己分析一下，如果集合是map需要怎么处理，比较好
}

func TestStringBytes(t *testing.T) {
	assert := assert.New(t)
	//不复制内存，把string转换为[]byte
	String2Bytes := func(s string) []byte {
		l := len(s)
		p := unsafe.StringData(s)
		return unsafe.Slice(p, l)
	}

	//不复制内存，把[]byte转换为string
	Bytes2String := func(bytes []byte) string {
		l := len(bytes)
		p := unsafe.SliceData(bytes)
		return unsafe.String(p, l)
	}

	str := "66"
	bytes := []byte(str)  //复制数据本身
	str2 := string(bytes) //复制数据本身
	_ = str2

	//string类型是不可修改的。但如果是通过unsafe方法产生出的变量，在语法上是可以修改的，如果你真的去修改它会有一个运行时的panic
	bytes2 := String2Bytes(str)
	//bytes2[0] = 48 //不要修改由string转换过来的值,有运行时的panic

	str3 := Bytes2String(bytes2)
	assert.Equal(str3, str)
}
