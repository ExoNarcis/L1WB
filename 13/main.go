// эталонный main.go
package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // тут скорее всего кто то поделил на 0...
		}

	}()
	a := 5
	b := 8
	fmt.Println("a=", a, "b=", b)
	a, b = b, a // паралельное присваивание (паралельное не в смысле что одновременно)
	fmt.Println("a=", a, "b=", b)
	//XOR
	// a = 0101, b = 1000
	a ^= b // a = a^b (0101 XOR 1000 = 1101) = 13
	b ^= a // b = b^a (1000 XOR 1101 = 0101) = 5
	a ^= b // a = a^b (1101 XOR 0101 = 1000) = 8
	// a	b	a Xor b
	// 0	0		0
	// 0	1		1
	// 1	0		1
	// 1	1		0
	fmt.Println("a=", a, "b=", b) // XOR
	// +-
	a += b                        // a = 5 + 8 = 13
	b = a - b                     // b = 13 - 8 = 5
	a -= b                        // a = 13 - 5 = 8
	fmt.Println("a=", a, "b=", b) // +/-
	// на умножение надо наложить ОДЗ на != 0
	a *= b                        // a = 5 * 8 = 40
	b = a / b                     // b = 40 / 8 = 5
	a /= b                        // a = 40 / 5 = 8
	fmt.Println("a=", a, "b=", b) // * and /
}
