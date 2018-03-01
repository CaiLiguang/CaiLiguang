package main
import  (
	"fmt"
	"math"
	)

func main() {
	//定义局部变量
	var a int = 100
	var b int = 200
	// var ret int

	// //调用函数并返回最大值
	// ret = max(a,b)
	// fmt.Printf("最大值是：%d\n" , ret)
	fmt.Printf("交换前a的值为：%d\n" , a)
	fmt.Printf("交换前b的值为：%d\n" , b)

	//通过调用函数来交换值
	swap(&a,&b)
	fmt.Printf("交换后a的值：%d\n" , a)
	fmt.Printf("交换后b的值：%d\n" , b)
    fmt.Println(a, b)

    //求平方根
    //getSquareRoot := func(x float64) float64{
	get := func(x float64) float64{
    	return math.Sqrt(x)
    }

    //使用平方根喊护士
    fmt.Println(get(16))
 
}

//函数返回两个数的最大值
func max(num1 , num2 int) int {
	//定义局部变量
	var result int
	if( num1 > num2)  {
		result = num1
	} else {
		result = num2
	}
	return result
}

//定义相互交换值的函数
func swap(x *int, y *int) int {
	var temp int
	temp = *x //保存x的值
	*x = *y //將y值赋给x
	*y = temp //將temp值赋给y
	return temp
}