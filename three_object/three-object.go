package three_object

import (
	"fmt"
	"math"
)

// 形状接口
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// 矩形结构体
type Rectangle struct {
	length float64 // 长
	width  float64 // 宽
}

func (r *Rectangle) Area() float64 {
	fmt.Println("矩形面积：", r.length*r.width)
	return r.length * r.width
}

func (r *Rectangle) Perimeter() float64 {
	fmt.Println("矩形周长：", (r.length+r.width)*2)
	return (r.length + r.width) * 2
}

// 圆形结构体
type Circle struct {
	Radius float64 // 半径
}

func (c *Circle) Area() float64 {
	fmt.Println("圆形面积：", math.Pi*(c.Radius*c.Radius))
	return math.Pi * (c.Radius * c.Radius)
}

func (c *Circle) Perimeter() float64 {
	fmt.Println("圆形周长：", 2*math.Pi*c.Radius)
	return 2 * math.Pi * c.Radius
}

// 题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
// 考察点 ：接口的定义与实现、面向对象编程风格。
func GetOne() {
	// 声明并初始化一个矩形
	// rectangle := Rectangle{2, 3}
	var rectangle Shape = &Rectangle{2, 3}
	rectangle.Area()
	rectangle.Perimeter()

	// 声明并初始化一个矩形
	// circle := &Circle{3}
	var circle Shape = &Circle{3}
	circle.Area()
	circle.Perimeter()
}

// 抽象人结构体
type AbstractPerson struct {
	Name string // 名字
}

// 人结构体
type Person struct {
	AbstractPerson
	Age int // 年龄
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Println("员工姓名：", e.Name)
	fmt.Println("员工年龄：", e.Age)
	fmt.Println("员工工号：", e.EmployeeID)
}

// 题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
// 考察点 ：组合的使用、方法接收者。
func GetTwo() {
	employee := Employee{
		Person{
			AbstractPerson{"Mrs Zhou"},
			23,
		},
		1,
	}
	employee.PrintInfo()
}
