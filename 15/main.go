package main

import (
	"fmt"
	"strings"
)

// 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

// 1) я хз гарантирует ли длину строки createHugeString такой какой ее передали
// 2) работает хорошо пока не появляются всякие кодировОчки - руны и т.д. хочу я что бы createHugeString на выходе сделал строку с символом '€' n повторений тогда justString = v[:100] удачно сломает строку

var justString string

func createHugeString(count int) string {
	sb := strings.Builder{}
	for i := 0; i < count; i++ {
		sb.WriteString(" ")
		//sb.WriteString("1")
		//sb.WriteString("€")
	}
	return sb.String()
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 { // проверю длину
		// justString = v[:100] - сломает руну если там utf-8
		justString = string([]rune(v)[:100])
	}
}

func main() {
	someFunc()
}
