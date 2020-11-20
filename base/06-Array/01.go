package main

import "fmt"

func main01() {
	/*
		定义数组
		var 数组变量名 [元素数量]T
	*//*  */

	// 定义长度为2的int类型数组
	var num [2]int
	num[0] = 10
	num[1] = 20
	// fmt.Println(num)

	/*
	方法一
		初始化数组时可以使用初始化列表来设置数组元素的值。
	*/
	// 数组初始化
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]
}

/* 
	方法二
		按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下我们可以让编译器根据初始值的个数自行推断数组的长度，例如：
	*/
func main02() {
	var testArray [3]int
	var numArray = [...]int{1, 2} // 长度自动推断
	var cityArray = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray)                          //[0 0 0]
	fmt.Println(numArray)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray)   //type of numArray:[2]int
	fmt.Println(cityArray)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray) //type of cityArray:[3]string
}

/*
方法三
我们还可以使用指定索引值的方式来初始化数组，例如:
 */

func main04() {
	a := [...]int{1: 5, 3: 8}
	fmt.Println(a)
}

/*
数组的遍历
*/
func main05() {
	arr := [...]string{"西安", "北京", "南京"}

	//for循环遍历属组
	for index := 0; index < len(arr); index++ {
		fmt.Printf("索引%d =>%s \n", index, arr[index])
	}

	//for range遍历
	for i, value := range arr {
		fmt.Println(i, value)
	}
}

func reverse(arr [3]int) {
	arr[1] = 100
	fmt.Println("reverse =", arr)
}

/*
数组是值传递，作为参数传递时传递的是数组的副本，修改不会影响原来的数组
*/
func main06() {
	arr := [...]int{1, 2, 3}
	reverse(arr)
	fmt.Println("arr = ", arr)
}

func main() {
	// 二维数组第定义
	arr := [3][2]string{
		{"a1", "a2"},
		{"a3", "a3"},
		{"a5", "a6"},
	}

	// 二维数组的遍历
	for i, v := range arr {
		for j, item := range v {
			fmt.Println(arr[i], j, item)
		}
	}
}
