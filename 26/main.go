package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

func search1(str string) bool { // принемаем строки
	runs := []rune(strings.ToLower(str)) // кастуем в руны по tolowercase строке (перевод в нижний регистр)
	for k, v := range runs {             // цикл по рунам
		for i := k + 1; i < len(runs); i++ { // ищем есть ли дальше в рунах такой элемент
			if v == runs[i] { // нашли
				return false // строка не уникальная
			}
		}
	}
	return true // если уникальная
}

func main() {
	fmt.Println(search1("abcd")) // вывод и примера
	fmt.Println(search1("abCdefAaf"))
	fmt.Println(search1("aabcd"))
}
