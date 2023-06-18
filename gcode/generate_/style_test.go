package generate_

// [Talk about Go Generics](https://github.com/gopherchina/conference/blob/master/2023/2.1.6%20%E8%B0%88%E8%B0%88%20Go%20%E6%B3%9B%E5%9E%8B.pdf)
// [Talk-about-generics](https://github.com/smallnest/talk-about-go-generics)

// ------------------------
// ~ 可以使用基本类型，不可以使用interface，也不可以使用有名的struct
type Error struct {
	Code  int
	Error string
}

// type I0 interface {
// 	~Error // struct
// 	~error // interface
// }

// ~ 中可以使用这种无名的struct
type Error2 = struct {
	Code  int
	Error string
}
type I00 interface {
	~Error2 // OK
}

// -------------------------
// type T1[P any] P
// type T2[P any] struct{ *P }

// func I1[K any, V interface{ K }]()       {}
// func I2[K any, V interface{ int | K }]() {}

// --------------------------
func I15[K interface{ comparable }]() {

}
func I16[K interface {
	error
	comparable
}]() {

}

// func I17[K interface{ comparable | os.File }]() {}

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------

// --------------------------
