package main

import (
	"fmt"
)

// 演示golang中指针类型  https://blog.csdn.net/qq_42883074/article/details/121402072
func main() {

	//基本数据类型在内存布局
	var i int = 10
	// i 的地址是什么,&i
	fmt.Println("i的地址=", &i)

	//下面的 var ptr *int = &i
	//1. ptr 是一个指针变量
	//2. ptr 的类型 *int
	//3. ptr 本身的值&i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)            // 打印的是变量i的内存地址
	fmt.Printf("打印变量ptr的内存地址 =%v\n", &ptr) //打印变量ptr的内存地址
	fmt.Printf("ptr 指向的值=%v\n", *ptr)      // 如果想要输出存放的内存地址的值，要在变量前面加星号(*)   *ptr 取变量i地址中的值
	fmt.Printf("*&i=%v\n", *&i)            //  *&i  *取值 &打印地址   等同于 10
	fmt.Printf("ptr 指向的值=%v\n", *&ptr)     // *&ptr 取&ptr中存放的值  即i的内存地址  等同于&i
	fmt.Printf("ptr 指向的值=%v\n", **&ptr)    // 取变量i的值

}
