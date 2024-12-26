package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")

	hello("Марина") //1
	hello("Лунтик") //1
	fmt.Println("---")

	fmt.Println(printEven(2, 7))   //2
	fmt.Println(printEven(0, 36))  //2
	fmt.Println(printEven(55, 17)) //2
	fmt.Println("---")

	fmt.Println(apply(3, 2, "+")) //3
	fmt.Println(apply(3, 2, "-")) //3
	fmt.Println(apply(3, 2, "*")) //3
	fmt.Println(apply(3, 2, "/")) //3
	fmt.Println(apply(3, 0, "/")) //3
}

func hello(name string) {
	fmt.Println("Hello, " + name + "!")
}

func printEven(num1 int, num2 int) (string, error) {
	if num1 > num2 {
		return "0", (fmt.Errorf("ошибка"))
	}
	var str = ""
	for i := num1; i <= num2; i++ {
		var del = i % 2
		if del == 0 {
			str = str + " " + strconv.Itoa(i)
		}
	}
	return str, nil

}

func apply(num1 float64, num2 float64, oper string) (float64, error) {
	if oper == "+" {
		var res = num1 + num2
		return res, nil
	} else if oper == "-" {
		var res = num1 - num2
		return res, nil
	} else if oper == "*" {
		var res = num1 * num2
		return res, nil
	} else if oper == "/" {
		if num2 != 0 {
			var res float64 = (num1 / num2)
			return res, nil
		}
		if num2 == 0 {
			return 0, (fmt.Errorf("ошибка. Деление на ноль невозможно"))

		}
	}
	return 0, (fmt.Errorf("оператор не найден"))
}
