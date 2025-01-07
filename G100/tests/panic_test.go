package tests

//
//import (
//	"fmt"
//	"testing"
//)
//
//func TestPanic(t *testing.T) {
//	//defer func() {
//	//	if r := recover(); r != nil {
//	//		fmt.Println("recover", r)
//	//	}
//	//}()
//
//	f()
//	if r := recover(); r != nil {
//		fmt.Println("recover", r)
//	}
//}
//
//func f() {
//	fmt.Println("a")
//	panic("foo")
//	fmt.Println("b")
//}
//
//func Foo() error {
//	err := bar()
//	if err != nil {
//		return fmt.Errorf("bar failed: %w", err)
//	}
//	// ...
//}
//
//type BarError struct {
//	Err error
//}
//
//func (b BarError) Error() string {
//	return "bar failed:" + b.Err.Error()
//}
