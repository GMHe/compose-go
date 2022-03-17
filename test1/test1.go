package test1

import (
	"crypto/md5"
	"fmt"
)

func TestDd() {


		var input string

		print("请输入一个字符串：")

		fmt.Scanf("%s", &input)

		fmt.Printf("输入字符串为：%s\n",input)

		data := []byte(input)

		dds := md5.Sum(data)
		md5 := md5.Sum(data)


		fmt.Printf("md5后的值为:%x%x\n", dds,md5)


}
