package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func newList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{nums[0], nil}
	tail := head
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		tail.Next = node
		tail = tail.Next
	}
	return head
}

func displayList(head *ListNode) string {
	var res strings.Builder
	if head == nil {
		return res.String()
	}
	cur := head
	for cur != nil {
		res.WriteString(strconv.Itoa(cur.Val))
		res.WriteString(" ")
		cur = cur.Next
	}
	return res.String()
}

func mergeSort(nums []int) {
	mergeSort_r(nums, 0, len(nums) - 1)
}

func mergeSort_r(nums []int, p, r int) {
	if p >= r {
		return
	}
	q := p + (r - p)/2
	mergeSort_r(nums, p, q)
	mergeSort_r(nums, q+1, r)
	merge(nums, p, q, r)
}

func merge(nums []int, p, q, r int) {
	i, j := p, q+1
	tmp := make([]int, r-p+1)
	k := 0
	for i <= q && j<= r {
		if nums[i] <= nums[j] {
			tmp[k] = nums[i]
			i++
		} else {
			tmp[k] = nums[j]
			j++
		}
		k++
	}

	for i <= q {
		tmp[k] = nums[i]
		i++
		k++
	}

	for j <= r {
		tmp[k] = nums[j]
		j++
		k++
	}
	for k = 0; k <= r - p; k++ {
		nums[p+k] = tmp[k]
	}
}

type User struct {
	ID int
	Num int
}

func test(array [2]int) {
	array[0] = 100
	array[1] = 100
}

func tests(array []int) {
	array[0] = 100
	array[1] = 100
}

func main() {
	list := newList([]int{45, 2, 89, 32, 198, 90})
	fmt.Println(displayList(list))
	nums := []int{45, 2, 89, 32, 198, 90}
	mergeSort(nums)
	fmt.Println(nums)

	s := make([]*User, 0, 10)
	s = append(s, &User{1, 99})
	s = append(s, &User{2, 99})

	for _, user := range s {
		fmt.Println("%p", &user)
		user.Num = 100
	}

	fmt.Println(s)

	for _, user := range s {
		fmt.Println(user.ID, user.Num)
	}

	array := [2]int{99, 99}
	test(array)
	fmt.Println(array)

	ss := make([]int, 2)
	ss[0] = 99
	ss[1] = 99
	tests(ss)
	fmt.Println(ss)

	fmt.Printf("%T\n", ss)
	fmt.Printf("%T\n", array)
	fmt.Println(reflect.TypeOf(ss))
}
