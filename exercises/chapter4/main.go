package main

import "fmt"
import "container/list"

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

func lenOfList(l *list.List) int {
	counter := 0
	for e := l.Front(); e != nil; e = e.Next() {
		counter++
	}
	return counter
}

func main() {
	l := new(list.List)
	l.PushBack("hi")
	l.PushBack("world")
	l.PushBack(2)
	fmt.Println(lenOfList(l))
	fmt.Println(sum([]int{2, 3, 5}))
}
