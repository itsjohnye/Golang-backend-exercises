package main

import "fmt"

func main() {

	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c -= 3 //解密时向左移动
			if c < 'a' {
				c += 26 //回绕，因为上面解密时是减，所以要加回来。以英文26个字母计。
			}
		} else if c >= 'A' && c <= 'Z' { //for capital letters
			c -= 3
			if c < 'A' {
				c += 26
			}
		}

		fmt.Printf("%c", c) //格式化变量‘%c’打印出的是字符而不是数字本身。
	}

}
