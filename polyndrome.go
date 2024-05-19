package main

import (
	"fmt"
)

func main() {
	fmt.Print("Введите число: ")
	var number int
	fmt.Scanln(&number)

	originalNumber := number
	if number < 0 {
		number = -number
	}

	reversedNumber := 0
	temp := number
	for temp != 0 {
		remainder := temp % 10
		reversedNumber = reversedNumber*10 + remainder
		temp /= 10
	}

	isPalindrome := number == reversedNumber

	if isPalindrome {
		fmt.Printf("Число %d является палиндромом.\n", originalNumber)
	} else {
		fmt.Printf("Число %d не является палиндромом.\n", originalNumber)
	}
}

//Задача 3*
//Проверка числа на палиндром:
//Запросите у пользователя число и определите, является ли оно
//палиндромом. Палиндромом считается число, которое читается
//одинаково как справа налево, так и слева направо.
//Например, 121, 12321.
