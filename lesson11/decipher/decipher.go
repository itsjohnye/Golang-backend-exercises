package main

import "fmt"

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keywords := "GOLANG"
	message := ""
	keyIndex := 0

	/*
		// no range
		for i := 0; i < len(cipherText); i++ {
			// A=0, B=1, ... Z=25
			c := cipherText[i] - 'A'
			k := keywords[keyIndex] - 'A'

			// cipher letter - key letter
			c = (c-k+26)%26 + 'A'
			message += string(c)

			// increment keyIndex
			keyIndex++
			keyIndex %= len(keywords)
		}
	*/

	//range
	for i := range cipherText {
		// A=0, B=1, ... Z=25
		c := cipherText[i] - 'A' //此时减'A'的目的是让他落回以"A=0,B=1..."的26个字母计数内。实际'A'~'Z'对应65~91。
		k := keywords[keyIndex] - 'A'

		// cipher letter - key letter
		c = (c-k+26)%26 + 'A' //减k表示向左移动k位，加26表示值不应为负，%取模操作表示值要落在0～26区间，再加上'A'表示还原到原本的值
		message += string(c)

		// increment keyIndex
		keyIndex++
		keyIndex %= len(keywords) // keyIndex = keyIndex % len(keywords)， 取模运算符%计算两数相除的余数。如27 % 26 = 1
	}

	fmt.Println(message)
}
