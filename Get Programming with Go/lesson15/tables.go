package main

import "fmt"

type celsius float64
type fahrenheit float64

//celsiusToFahrenheit
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

//fahrenheitToCelsius
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

//设置常量画成表格
const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n" //%s is the uninterpreted bytes of the string or slice
	numberFormat = "%.2f"
)

type getRowFn func(row int) (string, string) //类型getRowFn是一个function，以row作为参数，返回两个string。

//draw celsiusToFahrenheit
func ctof(row int) (string, string) {
	c := celsius(row*5 - 40) //第一个row是0，第二个row是1...以5度为一个阶梯
	f := c.fahrenheit()
	cell1 := fmt.Sprintf(numberFormat, c) //with numberFormat
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2
}

//draw fahrenheitToCelsius
func ftoc(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

//draw the table
func drawTable(header1, header2 string, rows int, getRow getRowFn) { //此时getRow的类型是getRowFn，即它也以row作为参数，返回两个string。
	fmt.Println(line)
	fmt.Printf(rowFormat, header1, header2)
	fmt.Println(line)

	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2) //with rowFormat
	}

	fmt.Println(line)
}

func main() {
	drawTable("ºC", "ºF", 29, ctof) //把ctof这个函数作为drawTable的一个变量。29=(40+100)/5+1。
	fmt.Println()
	drawTable("ºF", "ºC", 29, ftoc)
}
