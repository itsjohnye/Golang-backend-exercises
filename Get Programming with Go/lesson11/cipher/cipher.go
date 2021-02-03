package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "your message goes here"
	keyword := "GOLANG"
	keyIndex := 0
	cipherText := ""

	text = strings.ToUpper(strings.Replace(text, " ", "", -1)) //Replace函数第1个参数表示要转换的内容scope，第2个参数表示将要被替代的内容，第3个参数表示将要变成什么，第4个参数表示替换几项，比如3表示替换掉前3项，-1表示所有都替换掉。这里是把所有空格都去掉的意思。
	keyword = strings.ToUpper(strings.Replace(keyword, " ", "", -1))

	for i := range text {
		c := text[i]
		if c >= 'A' && c <= 'Z' {
			// A=0, B=1, ... Z=25
			c -= 'A'
			k := keyword[keyIndex] - 'A'

			// cipher letter + key letter
			c = (c+k)%26 + 'A'

			// increment keyIndex
			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
