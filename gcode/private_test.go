package gcode

import (
	"github.com/modern-go/reflect2"
	"log"
	"testing"
	"unsafe"
)

func TestPrivate(t *testing.T) {
	type PrivateField struct {
		Name string
		str  string
		num  int64
	}
	//通过reflect修改私有字段的值
	func() {
		pf := PrivateField{
			Name: "1",
			str:  "2",
			num:  3,
		}
		tt2 := (reflect2.TypeOf(pf)).(*reflect2.UnsafeStructType)
		v := tt2.FieldByName("str")
		str := "sdf"
		v.UnsafeSet(unsafe.Pointer(&pf), unsafe.Pointer(&str))
		v = tt2.FieldByName("num")
		num := 9
		v.UnsafeSet(unsafe.Pointer(&pf), unsafe.Pointer(&num))

		v = tt2.FieldByName("Name")
		v.Set(&pf, &str)
		log.Print(pf.str, pf.num)
	}()

}
