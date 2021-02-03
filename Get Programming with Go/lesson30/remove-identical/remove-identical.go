package main

import (
	"fmt"
)

func sourceString(downstream chan string) {
	for _, v := range []string{"a", "a", "b", "c", "c", "a", "d", "d", "e", "d", "e"} {
		downstream <- v
	}
	close(downstream)
}

func filterString(upstream, downstream chan string) {
	//先建一个preList存储上游来的字符。
	preList := []string{}
	for uncheckItem := range upstream {
		preList = append(preList, uncheckItem)
	}

	//处理重复的值。此时建立了一个map做键值对比用。
	checkedList := []string{}
	keys := make(map[string]bool)
	//如果没有对应的键值（bool为false），就把bool值改为true，并加入checkedList（登记上）；如果有，就跳过。
	for _, entry := range preList { //entry：登记
		if _, value := keys[entry]; !value { //这里!value = false，即 if false
			keys[entry] = true
			checkedList = append(checkedList, entry)
		}
	}

	//再把处理好的checkedList一个个送入通道。
	for _, v := range checkedList {
		downstream <- v
	}
	close(downstream)
}

func printString(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}
func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceString(c0)
	go filterString(c0, c1)
	printString(c1)

}
