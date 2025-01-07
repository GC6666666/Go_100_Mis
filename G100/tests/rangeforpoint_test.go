package tests

import (
	"fmt"
	"testing"
)

// student 结构体表示一个学生
type student struct {
	studentId int
	name      string
	mathScore int
}

// String 方法实现 fmt.Stringer 接口，用于格式化输出 student 结构体
func (s student) String() string {
	return fmt.Sprintf("{studentId: %d, name: %s, mathScore: %d}", s.studentId, s.name, s.mathScore)
}

// class 结构体表示一个班级，包含多个学生
type class struct {
	classname string
	students  map[string]*student
}

// String 方法实现 fmt.Stringer 接口，用于格式化输出 class 结构体
func (c class) String() string {
	result := fmt.Sprintf("class: %s \n{students: ", c.classname)
	for key, student := range c.students {
		result += fmt.Sprintf("%s: %s, ", key, student)
	}
	result += "}"
	return result
}

// Put 方法将一个学生添加到班级中
func (c *class) Put(iid string, stu *student) {
	c.students[iid] = stu
}

// TestRangeForPoint 测试函数，验证 class 和 student 结构体的功能
func TestRangeForPoint(t *testing.T) {
	// 创建一个学生实例
	students := make([]*student, 3)
	for i := range students {
		students[i] = &student{
			studentId: i,
			name:      "GongChao",
			mathScore: 90 + i,
		}
	}
	fmt.Println("Initial students:")
	for _, v := range students {
		fmt.Println(v)
	}

	// 创建一个班级实例，并将学生添加到班级中
	c := class{
		classname: "Class 1",
		students:  make(map[string]*student),
	}
	for i, stu := range students {
		c.Put(fmt.Sprintf("student%d", i), stu)
	}

	// 打印班级和学生信息
	fmt.Println("After adding students to the class:")
	fmt.Println(c)

	// 修改班级中的学生信息
	c.students["student0"].studentId = 22
	c.students["student0"].name = "UpdatedName"
	fmt.Println("After modifying student information:")
	fmt.Println(c)
	students[0].studentId = 100
	fmt.Println(c)
}
