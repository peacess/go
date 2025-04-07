package interface_

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

type InterfaceA interface {
	A()
}

type InterfaceB interface {
	B()
}

type Impl struct{}

func (c *Impl) A() {
	println("impl a")
}

func (c *Impl) B() {
	println("impl b")
}

func TestType(t *testing.T) {
	ass_ := assert.New(t)
	var a InterfaceA = &Impl{}

	switch a.(type) {
	case InterfaceB:
		println("case InterfaceB")
	case InterfaceA:
		println("case InterfaceA")
	case *Impl:
		println("case *Impl")
	}
	{
		i := (*iface)(unsafe.Pointer(&a))
		println(i)
	}
	{
		ta, ok := a.(InterfaceA)
		ass_.Equal(ok, true)
		i := (*iface)(unsafe.Pointer(&ta))
		println(i)
	}
	{
		tb, ok := a.(InterfaceB)
		ass_.Equal(ok, true)
		i := (*iface)(unsafe.Pointer(&tb))
		println(i)
	}
	{
		_, ok := a.(*Impl)
		ass_.Equal(ok, true)
	}

}

// see go
type iface struct {
	Tab  *itab
	Data unsafe.Pointer
}

type eface struct {
	Type *_type
	Data unsafe.Pointer
}
type _type struct {
	Size_       uintptr
	PtrBytes    uintptr // number of (prefix) bytes in the type that can contain pointers
	Hash        uint32  // hash of type; avoids computation in hash tables
	TFlag       uint8   // extra type information flags
	Align_      uint8   // alignment of variable with this type
	FieldAlign_ uint8   // alignment of struct field with this type
	Kind_       uint8   // enumeration for C
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	// GCData stores the GC type data for the garbage collector.
	// Normally, GCData points to a bitmask that describes the
	// ptr/nonptr fields of the type. The bitmask will have at
	// least PtrBytes/ptrSize bits.
	// If the TFlagGCMaskOnDemand bit is set, GCData is instead a
	// **byte and the pointer to the bitmask is one dereference away.
	// The runtime will build the bitmask if needed.
	// (See runtime/type.go:getGCMask.)
	// Note: multiple types may have the same value of GCData,
	// including when TFlagGCMaskOnDemand is set. The types will, of course,
	// have the same pointer layout (but not necessarily the same size).
	GCData    *byte
	Str       int32 // string form
	PtrToThis int32 // type for pointer to this type, may be zero
}
type itab struct {
	Inter *InterfaceType
	Type  *_type
	Hash  uint32     // copy of Type.Hash. Used for type switches.
	Fun   [1]uintptr // variable sized. fun[0]==0 means Type does not implement Inter.
}
type InterfaceType struct {
	_type
	PkgPath Name      // import path
	Methods []Imethod // sorted by hash
}
type Name struct {
	Bytes *byte
}
type Imethod struct {
	Name int32 // name of method
	Typ  int32 // .(*FuncType) underneath
}
