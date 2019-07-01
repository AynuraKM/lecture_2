package main

import (
	"fmt"
)

func main() {
	st := stack{}

	for i := 0; i < 10; i++ {
		st.push(i)
	}
	fmt.Println(st)
	fmt.Println(st.Len())
	fmt.Println(st.Pop())
}

type stack struct {
	values []int
	l      int
	x      int
}

func (st *stack) push(value int) {
	st.values = append(st.values, value)
}

func (st *stack) Len() int {
	return len(st.values)
}

func (st *stack) Pop() interface{} {
	if len(st.values) == 0 {
		return nil
	}
	l := len(st.values)
	x := st.values[l-1]
	l --
	fmt.Println(l)
	fmt.Println(st.values)
	return x
}

func (st *stack) peek() interface{} {
	if len(st.values) == 0 {
		return nil
	}
	l := len(st.values)
	res := st.values[l-1]
	return res
}
