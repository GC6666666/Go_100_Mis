package tests

import (
	"fmt"
	"math"
	"testing"
)

// 定义一个 Shape 接口
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 定义 Rectangle 结构体
type Rectangle struct {
	Width, Height float64
}

// Rectangle 实现 Area 方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle 实现 Perimeter 方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 定义 Circle 结构体
type Circle struct {
	Radius float64
}

// Circle 实现 Area 方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle 实现 Perimeter 方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 定义一个函数，接受 Shape 接口类型的参数
func PrintShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func TestPrintShapeInfo(t *testing.T) {
	r := Rectangle{Width: 5, Height: 3}
	c := Circle{Radius: 2.5}

	fmt.Println("Rectangle:")
	PrintShapeInfo(r)

	fmt.Println("Circle:")
	PrintShapeInfo(c)
}
