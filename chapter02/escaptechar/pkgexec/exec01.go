package main

//表示该hello.go文件所在的包是main,在go中，每个文件都必须归属于一个包。
import "fmt" //fmt包中提供格式化，输出，输入的函数.

func main() {
	//要求：要求：请使用一句输出语句，达到输入如下图形的效果
	fmt.Println("姓名\t年龄\t籍贯\t地址\njohn\t20\t河北\t北京")

	//1 	go build 编译go文件, mac会生成unix可执行文件,windos生成.exe文件
	//2     步骤1 生成的可执行文件可以直接运行
	//3  go run exec01.go 通过gorun命令可以直接运行hello.go程序[类似执行一个脚本文件的形式]

	//go 执行流程
	//1  .go 文件 -> go build -> unix可自行文件  -> 执行 ->结果
	//2  .go 文件 -> go run -> 结果

}
