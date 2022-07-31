// эталонный main.go
package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

func revertwidthsplit(str *string) { // получаем строку
	srtarr := strings.Split(*str, " ")      // сплитаем в массив строк по " "
	*str = ""                               // обнуляем исходную строку
	for i := len(srtarr) - 1; i >= 0; i-- { // двигаемся обратно по массиву
		*str = *str + " " + srtarr[i] // складываем строки
	}
}

func main() {
	sw := "€snow dog sun — sun dog snow€" // строка с utf-8
	fmt.Println(sw)                       // печать до
	revertwidthsplit(&sw)                 // переворачивам по словам
	fmt.Println(sw)                       // печать после
}
