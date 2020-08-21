package main

import "fmt"
import "container/list"

// ex 4.1
func sum(i []int) int {
	if i == nil {
		return 0
	}
	if len(i) == 1 {
		return i[0]
	}
	rest := i[1:]
	return i[0] + sum(rest)
}

// ex 4.2
func lenOfList(l *list.List) int {
	counter := 0
	for e := l.Front(); e != nil; e = e.Next() {
		counter++
	}
	return counter
}

// ex 4.3
func maxInList(l *list.List) int {
	if l == nil {
		return 0 //err
	}
	max := 0
	//also need check for int list
	for _, v := range sliceInts(l) {
		if v > max {
			max = v
		}
	}
	return max
}

func sliceInts(l *list.List) (items []int) {
	for e := l.Front(); e != nil; e = e.Next() {
		items = append(items, e.Value.(int))
	}
	return
}

// ex 4.4

func binarySearchRecursevly(i []int, x, l, r int) int {
	if l > r {
		return -1
	}

	mid := int((l + r) / 2)
	if i[mid] > x {
		return binarySearchRecursevly(i, x, l, mid)
	} else if i[mid] < x {
		return binarySearchRecursevly(i, x, mid+1, r)
	} else {
		return mid
	}
}

func main() {
	l := new(list.List)
	l.PushBack(2)
	l.PushBack(23)
	l.PushBack(14)
	fmt.Println(sum([]int{2, 3, 5}))
	fmt.Println(lenOfList(l))
	fmt.Println(maxInList(l))
	fmt.Println(binarySearchRecursevly([]int{0, 1, 2, 3, 4, 5, 6}, 5, 0, 5))
}
