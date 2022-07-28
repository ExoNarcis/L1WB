package main

import (
	"fmt"
	"math/rand"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

func main() {
	chx := make(chan int)  // канал с x
	chx2 := make(chan int) // канал  c x^2

	go func(chx chan<- int) {
		for { // горутина генерирует x в канал chx (1 канал)
			r := rand.Intn(100)
			fmt.Println("rand gen x =", r)
			chx <- r
		}
	}(chx)

	go func(chx <-chan int, chx2 chan<- int) {
		for v := range chx { // горутина читает из канала chx возводит в квадрат и записывает в chx2
			fmt.Println(v, "^2 = ", v*v)
			chx2 <- v * v
		}

	}(chx, chx2)

	for v := range chx2 { // печатаем в "stdout"
		fmt.Println("v2 in out chann:", v)
	}
}
