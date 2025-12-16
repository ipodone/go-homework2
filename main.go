package main

import (
	"fmt"

	five_mu "github.com/ipodone/go-homework2/five-mu"
	"github.com/ipodone/go-homework2/four_channel"
	"github.com/ipodone/go-homework2/one_ptr"
	"github.com/ipodone/go-homework2/three_object"
	"github.com/ipodone/go-homework2/two_goroutine"
)

func main() {
	fmt.Println("one_ptr GetOne 开始===")
	a := 1
	fmt.Println("函数外部，修改前：", a)
	one_ptr.GetOne(&a)
	fmt.Println("函数外部，修改后的值（使用指针参数，会修改原实例的值；使用值参数，不会）：", a)
	fmt.Println("one_ptr GetOne 结束===")
	fmt.Println()

	fmt.Println("one_ptr GetTwo 开始===")
	b := []int{1, 2, 3}
	fmt.Println("函数外部，修改前：", b)
	one_ptr.GetTwo(&b)
	fmt.Println("函数外部，修改后的值（切片底层是数组，修改切片即修改数组的值，无论使用值传递或指针传递，都会修改原实例值）：", b)
	fmt.Println("one_ptr GetTwo 结束===")
	fmt.Println()

	fmt.Println("two_goroutine GetOne 开始===")
	two_goroutine.GetOne()
	fmt.Println("two_goroutine GetOne 结束===")
	fmt.Println()

	fmt.Println("two_goroutine GetTwo 开始===")
	two_goroutine.GetTwo()
	fmt.Println("two_goroutine GetTwo 结束===")
	fmt.Println()

	fmt.Println("three_goroutine GetOne 开始===")
	three_object.GetOne()
	fmt.Println("three_goroutine GetOne 结束===")
	fmt.Println()

	fmt.Println("three_goroutine GetTwo 开始===")
	three_object.GetTwo()
	fmt.Println("three_goroutine GetTwo 结束===")
	fmt.Println()

	fmt.Println("four_channel GetOne 开始===")
	four_channel.GetOne()
	fmt.Println("four_channel GetOne 结束===")
	fmt.Println()

	fmt.Println("four_channel GetTwo 开始===")
	four_channel.GetTwo()
	fmt.Println("four_channel GetTwo 结束===")
	fmt.Println()

	fmt.Println("five_mu GetOne 开始===")
	five_mu.GetOne()
	five_mu.GetTwo()
	fmt.Println("five_mu GetOne 结束===")
	fmt.Println()

	fmt.Println("five_mu GetTwo 开始===")
	five_mu.GetThree()
	five_mu.GetFour()
	five_mu.GetFive()
	fmt.Println("five_mu GetTwo 结束===")
	fmt.Println()
}
