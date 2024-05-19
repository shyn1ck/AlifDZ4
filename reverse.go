package main

import "fmt"

func main() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scanln(&number)
	fmt.Println("Ваше число: ", number)
	for i := number; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println(" //")
}

//Задача 2
//Вывод чисел в обратном порядке:
//Запросите у пользователя число и выведите все
//числа от этого числа до 1 в обратном порядке.
//// 5
//// 5 4 3 2 1
