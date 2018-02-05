package main

import "fmt"
import "unsafe"

// func main() {
// 	const LENGTH int = 10
// 	const WIDTH int = 5
// 	var area int
// 	const a, b, c = 1 , false , "str" //多重赋值

// 	area = LENGTH * WIDTH
// 	fmt.Printf("面积为： %d" ,area)
// 	println()
// 	println(a,b,c)

// 	const (
// 	    Unknown = 0
// 	    Female = 1
// 	    Male = 2
// 	)
// 	println(Female)
// }
const (
    Unknown = 0
    Female = 1
    Male = 2
)
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)
// println(Female)

//2.fmt.Printf不能打印整形
func main() {
	fmt.Printf("%d",Unknown)
    println(a, b, c)
}

