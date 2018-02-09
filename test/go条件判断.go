package main
import "fmt"

func main() {
	// //if 判断 定义局部变量
	// var a = 200
	// var b = 300

	// //if判断bool表达试
	// if( a < b){
	// 	//为true的情况下
	// 	fmt.Printf("a小于b\n")
	// }


	//switch
	// var grade string = ""
	// var marks int = 40
	// switch marks {
	// 	case 90 : 
	// 		grade = "A"
	// 	case 80 : 
	// 		grade = "B"
	// 	case 50,60,70 :
	// 		 grade = "C"
	// 	default : 
	// 		grade = "D"
	// }
	// switch {
	// 	case grade == "A" :
	// 		fmt.Printf("优秀！\n")
	// 	case grade == "B" :
	// 		fmt.Printf("良好! \n")
	// 	case grade == "C" :
	// 		fmt.Printf("及格! \n")
	// 	case grade == "D" :
	// 		fmt.Printf("不及格! \n")
	// 	default :
	// 		fmt.Printf("差! \n")
	// }	
	// fmt.Printf("你的等级是 %s \n" , grade)


    var c1, c2, c3 chan int
    var i1, i2 int
    select {
        case i1 = <-c1:
         fmt.Printf("received ", i1, " from c1\n")
        case c2 <- i2:
         fmt.Printf("sent ", i2, " to c2\n")
        case i3, ok := (<-c3):  // same as: i3, ok := <-c3
         if ok {
            fmt.Printf("received ", i3, " from c3\n")
        } else {
            fmt.Printf("c3 is closed\n")
        } 
        default:
         fmt.Printf("no communication\n")
    } 

}