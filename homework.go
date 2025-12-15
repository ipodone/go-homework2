package homework01

import "fmt"

// 1. 只出现一次的数字
// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
func SingleNumber(nums []int) int {
	fmt.Println(nums)

	// 定义一个map，map里存的是map
	numMap := make(map[int]map[int]int)

	// 得到数组中每个元素k，在数组中出现的次数v，封装成一个map
	for i1, v1 := range nums {
		tempNumMap := make(map[int]int)
		for _, v2 := range nums {
			if v2 == v1 {
				tempNumMap[v1]++
			}
		}
		// 然后将这些map，都放到到一个总的map中
		numMap[i1] = tempNumMap
	}

	fmt.Println(numMap)

	for _, v1 := range numMap {
		for k2, v2 := range v1 {
			if v2 == 1 { // 找到数组出现次数为1的第一个元素就停止
				fmt.Printf("数组出现次数为1的第一个元素为：%d\n\n", k2)
				return k2
			}
		}
	}

	return 0
}

// 2. 回文数
// 判断一个整数是否是回文数
func IsPalindrome(x int) bool {
	// 回文数：是指正着读和反着读都一样的数字（那把其中一个反过来，两个做下比较就可以了）
	if x < 0 {
		fmt.Println(x, false)
		return false
	}

	fmt.Println(x, ReverseByNum(x) == x)
	return ReverseByNum(x) == x

	// 直接把数字翻转后比较（注意溢出时 Reverse 返回 0）
	// fmt.Println(x, Reverse(x) == x)
	// return x == Reverse(x) // AI生成代码：考虑溢出
}

func ReverseByNum(x int) int { // 此处x变化，不影响原x
	reversed := 0
	for x != 0 { // 取余翻转
		digit := x % 10                // 获取最后一位数字
		reversed = reversed*10 + digit // 构建反转数
		x /= 10                        // 去掉最后一位
	}
	return reversed
}

// Reverse: 将整数翻转，溢出时返回 0
// func Reverse(x int) int { // AI生成代码：考虑溢出
// 	xi := int64(x)
// 	sign := int64(1)
// 	if xi < 0 {
// 		sign = -1
// 		xi = -xi
// 	}

// 	var rev int64 = 0
// 	for xi > 0 {
// 		rev = rev*10 + xi%10
// 		xi /= 10
// 		// 检查是否超过 32 位有符号整数范围
// 		if rev > int64(1<<31-1) { // 1<<31（即1*2的31次方）位运算：左移 每一位，统一向左移动一位 // 乘以2的n次方
// 			return 0
// 		}
// 	}
// 	return int(rev * sign)
// }

// 3. 有效的括号
// 给定一个只包括 '(', ')', '{', '}', '[', ']' 的字符串，判断字符串是否有效
func IsValid(s string) bool {
	// 切片支持栈：后进先出（LIFO）

	// 栈的核心操作
	// stack := []int{} // 初始化栈（空切片）

	// 入栈（Push）- O(1)
	// stack = append(stack, 10) // [10] 即切片追加
	// stack = append(stack, 20) // [10, 20]

	// 查看栈顶（Peek）- O(1)
	// top := stack[len(stack)-1] // 20

	// 出栈（Pop）- ？O(1)
	// stack = stack[:len(stack)-1] // [10] 即切片移除/删除：将前 切片长度-1（但不包含切片长度-1） 放到一个新的切片里，则最后/最顶层的那个就删除掉了

	// 	思路：
	// 1、遍历字符串的每个字符。
	// 2、遇到左括号 (，{，[ 时，将其压入栈中。
	// 3、遇到右括号 )，}，] 时：
	// 3.1、如果栈为空，说明没有对应的左括号，直接返回 false。
	// 3.2、弹出栈顶元素，检查是否与当前右括号匹配，不匹配则返回 false。
	// 4、遍历结束后，栈应为空（所有左括号都被匹配），否则返回 false。
	fmt.Println(s)
	stack := make([]rune, 0)
	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' { // 查看栈顶
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' { // 查看栈顶
				return false
			}
			stack = stack[:len(stack)-1] // 出栈
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' { // 查看栈顶
				return false
			}
			stack = stack[:len(stack)-1] // 出栈

		}
	}
	return len(stack) == 0
}

// 4. 最长公共前缀
// 查找字符串数组中的最长公共前缀
func LongestCommonPrefix(strs []string) string {
	fmt.Println(strs, len(strs[0]))

	if len(strs) == 0 {
		fmt.Println("")
		return ""
	}

	if len(strs) == 1 {
		fmt.Println(strs[0])
		return strs[0]
	}

	for i := 0; i < len(strs[0]); i++ { // {"flower", "flow", "flight"}、{"dog", "racecar", "car"}
		for _, v := range strs {
			if i >= len(v) || strs[0][i] != v[i] {
				fmt.Println(strs[0][:i])
				return strs[0][:i]
			}
		}
	}

	return ""
}

// 5. 加一
// 给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
func PlusOne(digits []int) []int {
	fmt.Println(digits)

	// 直接给数组最后一位数字+1
	// digits[len(digits)-1] = digits[len(digits)-1] + 1
	// return digits

	// 把数组先合成一个数字，然后+1，再拆成数组 // 描述思路，让AI生成代码
	tempNum := 0
	for _, v := range digits { // {1, 2, 3}
		tempNum = tempNum*10 + v
	}

	tempNum++ // 加一
	fmt.Println("加一后：", tempNum)

	// 将数字拆成数组
	if tempNum == 0 {
		return []int{1}
	}

	result := []int{}
	for tempNum > 0 {
		result = append([]int{tempNum % 10}, result...)
		fmt.Println("result", result)
		tempNum /= 10
		fmt.Println("tempNum", tempNum)
	}

	return result
}

// 6. 删除有序数组中的重复项
// 给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
func RemoveDuplicates(nums []int) int {
	// 双指针法：O(1)空间
	// 快指针 i 用来遍历整个数组
	// 慢指针 k 指向应该放置下一个不重复元素的位置
	fmt.Println(nums)

	if len(nums) == 0 {
		return 0
	}

	k := 0 // 慢指针，指向第一个不重复元素的位置
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[k] { // 找到不同的元素
			k++
			nums[k] = nums[i] // 将不重复的元素移到k的位置
			fmt.Println(nums)
		}
	}

	fmt.Println(nums)
	return k + 1 // 返回新数组的长度
}

// 7. 合并区间
// 以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
func Merge(intervals [][]int) [][]int {
	fmt.Println(intervals)

	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i][0] > intervals[i+1][0] {
			temp := intervals[i]
			intervals[i] = intervals[i+1]
			intervals[i+1] = temp
		}
	}

	reslut := make([][]int, 0)
	for i := 0; i < len(intervals)-1; i++ {
		if (intervals[i][1] >= intervals[i+1][0]) && (intervals[i][1] <= intervals[i+1][1]) {
			intervals[i][1] = intervals[i+1][1]
			reslut = append(reslut, intervals[i])
		} else {
			reslut = append(reslut, intervals[i+1])
		}
	}

	fmt.Println(reslut)
	return reslut
}

// 8. 两数之和
// 给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func TwoSum(nums []int, target int) []int {
	fmt.Println(nums, target)

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}

	return nil
}
