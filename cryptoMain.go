//main方法只能放在package main中，go run 是执行命令，
//必须要一个main用来调用，install可以直接编译成包文件，
//也可以编译出exe（如果有main函数的话）
//如果go run 的文件不是以main命名的就会报错
//go run: cannot run non-main package

package main 

import (
	"fmt"
	"./lib" //引用lib里面的自定义的包或是第三方包
)

func main() {
	//引用的包里面的函数名 必须是大写字母开头
	//不然会提示：cannot refer to unexported name 
	s := crypto.Encode("方燕良，哈哈哈哈，总是难忘")
	fmt.Println(s)
	fmt.Println(crypto.Decode(s))
}