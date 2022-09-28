package unsafe_

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

func TestStringBytes(t *testing.T) {
	ass_ := assert.New(t)

	str := "test 中文　world"
	bytes := []byte(str)
	//上面把string转换为bytes类型，　bytes会复制str的内存。因为bytes这个变量是slice类型，如果不做复制那么，一段修改bytes的值后，str的值就不正常
	//那么我们如何转换它们而不进行内存的复制呢，　可以使用unsafe相关的包来做。

	// []byte to string
	{
		bs := *(*string)(unsafe.Pointer(&bytes))
		ass_.Equal(str, bs)
		// 这个实现参考[strings.Builder](https://github.com/golang/go/blob/master/src/strings/builder.go)
		{
			builder := strings.Builder{}
			builder.String() //这里使用使用了bytes to string转换，且不复制内存
		}
		//现在我们来分析，这个转换为什么可以成功呢？
		//下面的代码是　reflect.StringHeader reflect.SliceHeader 的定义
		type StringHeader struct { //这就是真正的　string类型
			Data uintptr
			Len  int
		}
		type SliceHeader struct { //这是真正的　slice类型
			Data uintptr
			Len  int
			Cap  int
		}
		//从定义我们可以到，slice要比string多一个字段　“Cap”
		// 代码　“(*string)(unsafe.Pointer(&bytes))” 就是直接把bytes类型（也就是slice）的指针转换为　string类型。
		// 这个转换是安全的，因为当string的指针直接指向slice类型的内存时，string类型只会操作前两个字段，且它们与string是一样的。
		// bs := *(...) 这部分代码的作用是把　slice的前字段的值直接“复制”给一个新的string(bs)。注意bs的真实类型　StringHeader是新产生的
	}
	{ // string to []byte
		//我们参考前的方法，得到如下实现
		sb := *(*[]byte)(unsafe.Pointer(&str))
		ass_.Equal(bytes, sb)
		//上面的测试可以通过，但是这种转换是存在问题的
		//因为　SliceHeader的字段是三个，而StringHeader的是两个，那么把StringHeader的内存直接复制给SliceHeader时，第三个字段会是什么值呢？
		//它是StringHeader后面的一个int值，它的值是不确定的。所以我们需要给第三个字段给上正确的值
		(*reflect.SliceHeader)(unsafe.Pointer(&sb)).Cap = len(str)
		//这样可以得到正确的值，但是使用两条语句，不简洁，且[]byte在没有给Ｃap值之前是一种“不确定”状态，　所以我可以使用下面的方法
		sb = unsafe.Slice((*byte)(unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&str)).Data)), len(str))
		//上面的代码看起来有点长，可以使用下面的代码。这两种方式都可以使用，我建议使用上面一种，虽然长一点，但容易理解
		sb = unsafe.Slice(*(**byte)(unsafe.Pointer(&str)), len(str))

		//在上面的代码中我们有一个提问“第三个字段会是什么值呢？”，下面写代码来难证一下
		{
			var value1 int = 66
			var str2 string = "new "
			var value2 int = 99
			sb := *(*[]byte)(unsafe.Pointer(&str2))
			var value3 int = 102
			fmt.Printf("{%x}\n", (*reflect.SliceHeader)(unsafe.Pointer(&str2)).Cap)
			fmt.Printf("{%p}:{%p}:{%p}\n", &value1, &value2, &value3)
			//fmt.Printf("{%p}:{%p}:{%p}\n{%p}", &value1, &str2, &value2, &sb)
			_ = str2 //只是为了编译通过
			_ = value3
			_ = value2
			_ = value1
			_ = sb
			//下面是输出结果
			//{6d1f53}
			//{0xc000024c28}:{0xc000024c30}:{0xc000024c38}
		}
	}

}
