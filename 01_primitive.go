package main

import (
	"fmt"
)

var (
	i    int //
	i8   int8
	i1   int16
	i32  int32
	i64  int64
	ui32 uint32

	f  float32 // f float64
	s  string
	r  rune
	ii interface{}
)

func main() {
	//format code using this shortcut Ctrl+Alt+l
	i32 = 1<<31 - 1
	ui32 = 1<<32 - 1

	var array = []int{5, 2, 2, 3}
	fmt.Println("this is my integer", i32)
	fmt.Print("this is my integer: %v, %v, %v", array, 123, 123)

	fmt.Println("this is my unsigned integer", ui32)

	var runes []rune
	for i := 0; i < 100; i++ {
		runes = append(runes, 's')
		//Ctrl+space  вызовет меню
	}
	s := "Ainura"
	fmt.Println([]rune(s))

	var arr1 = []int{1, -31, 5, 0}

	fmt.Println("initial capacity is:", cap(arr1))
	arr1 = append(arr1, 5)
	fmt.Println("initial capacity is:", cap(arr1))

	//Ctrl + shift + v показывает все, что есть в буфере
	i := 0
	for i < 10000 {
		i++
		cap0 := cap(arr1)
		arr1 = append(arr1, 5)
		currCap := cap(arr1)
		if currCap != cap0 {
			fmt.Println("current capacity is: %1v\n", cap(arr1))
			fmt.Printf("koeff capacity is: %4.2f\n", float32(currCap)/float32(cap0))
		}

	}
}
