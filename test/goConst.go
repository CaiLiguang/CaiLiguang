package main

import "fmt"
// import "unsafe"

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

// unsafe.Sizeof 函数
//     t1 byte    1
//     t2 int32   4
//     t3 int64   8
//     t4 string  16
//     t5 bool    1
// const (
//     Unknown = 0
//     Female = 1
//     Male = 2
// )
// const (
//     a = "123"
//     b = len(a)
//     c = unsafe.Sizeof(true)
// )
// // println(Female)

// //2.fmt.Printf不能打印整形
// func main() {
// 	fmt.Printf("%d",Unknown)
//     println(a, b, c)
// }


const (
    i=1<<iota //0
	j=3<<iota //11 向左移1位 110 转换10进制 6    
	k=3<<2     	  //2 向左移2位 1100 转换10进制 12
	l=3<<3 		  //3  向左移3位 11000 转换10进制 24

    Active = 1 << iota //4 向左移4位 10000 转换10进制 16
    Send	//5 向左移5位 100000 转换10进制 32
    Receive	//6 向左移6位 1000000 转换10进制 64
)  
func main() {    
	fmt.Println("i=",i)   
	fmt.Println("j=",j)   
	fmt.Println("k=",k)   
	fmt.Println("l=",l) 
	fmt.Println(Active, Send, Receive)

}

// const (
    
// )
// func main() {
// }

