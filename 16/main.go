package main

import (
	"fmt"
	"math/rand"
)

// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Данная схема использует два индекса (один в начале массива, другой в конце), которые приближаются друг к другу,
// пока не найдётся пара элементов, где один больше опорного и расположен перед ним, а второй меньше и расположен после.
// Эти элементы меняются местами. Обмен происходит до тех пор, пока индексы не пересекутся. Алгоритм возвращает последний индекс.

func getparts[T Numeric](ar []T) int {
	i, j := 0, len(ar)-1
	base := ar[int(j/2)] // получаю базовое значение, беру из середины +- 1
	for {                // цикл подобие do while
		for ar[i] < base { // ищем первое с 0 значение которое больше базового с начала
			i++ // толкаем индекс пока не найдется число или мы не упремся в базовое
		}
		for ar[j] > base { // ищем второе значение которое меньше базового с конца
			j-- // толкаем индекс пока не найдется число или не наткнемся на базовое
		}
		if i >= j { // индесы пересеклись
			return j // вернем последний индекс что бы потом посмотреть все что до этого индеса
		}
		ar[i], ar[j] = ar[j], ar[i] // свапаем значение
		i, j = i+1, j-1             // толкаем 2 индекса друг к другу
	}
}

func QSort[T Numeric](ar []T) { // по интерфейсу все что числа или от них производное
	if len(ar) > 1 { // массив с длиной в 1 элемент уже идиально сортирован
		part := getparts(ar) // получим индекс на котором мы остановились либо (произвели обмен по базовому числу)
		QSort(ar[:part+1])   // допроверка до этого т.е мы будем проверять верно ли все и менять если надо от индекса слева (повторяем с левой части)
		QSort(ar[part+1:])   // допроверка до этого т.е мы будем проверять верно ли все и менять если надо от индекса справа (повторяем с левой справа)
	}
}

func main() {
	ar := [15]int{} // массив
	for i, _ := range ar {
		ar[i] = rand.Intn(100) // заполним псевдо-рандомом (я же зерно не установил)
	}
	fmt.Println(ar) // принтнем
	QSort(ar[:])    // вызов сортировки с неявным преобразованием с срез
	fmt.Println(ar) // повторный принт

}

// https://pkg.go.dev/golang.org/x/exp/constraints
type Signed interface { // все что численное со знаком и от них производное
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface { // все что безнаковое численное и от них производное
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Float interface { // все что числа с плав. зяп. и от них производное
	~float32 | ~float64
}

type Integer interface { // все что целые числа или от них производное
	Signed | Unsigned
}

type Numeric interface { // все что числа или от них производное
	Integer | Float
}
