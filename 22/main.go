package main

import (
	"fmt"
	"math/big"
)

// 22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a, b := big.NewInt(12<<50), big.NewInt(13<<49) // иницилазируем 2 переменные
	fmt.Println("a: ", a, "\tb: ", b)              // печать
	a.Mul(a, b)                                    // умножение
	fmt.Println("a: ", a, "\tb: ", b)              // печать
	a.Div(a, b)                                    // деление
	fmt.Println("a: ", a, "\tb: ", b)              // печать
	a.Add(a, b)                                    // сложение
	fmt.Println("a: ", a, "\tb: ", b)              // печать
	a.Sub(a, b)                                    // вычитание
	fmt.Println("a: ", a, "\tb: ", b)              // печать
}
