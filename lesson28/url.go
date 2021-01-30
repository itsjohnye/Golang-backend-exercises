package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	u, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err)
		fmt.Printf("%#v\n", err)

		//对*url.Error类型进行ok断言。
		if e, ok := err.(*url.Error); ok { //这里的ok是为了防止err底下没有*url.Error的情况出现，从而引发panic。如果err底下有*url.Error，则e被赋值等于*url.Error，ok为true。
			fmt.Println("Operation:", e.Op) //operation
			fmt.Println("URL:", e.URL)
			fmt.Println("Error:", e.Err)
		}
		os.Exit(1)
	}

	fmt.Println(u)
}
