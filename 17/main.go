package main

import (
	"fmt"
)

// Реализовать бинарный поиск встроенными методами языка.

func binforseachS[T Numeric](ar []T, key T) int { // Класика + сдвиги слайса
	ofset := 0
	for { // пока индексы не соприкоснулись
		midleindex := int(len(ar) / 2) // индекс середины
		mid := ar[midleindex]          // эл с середины
		if mid > key && len(ar) > 1 {  // если эл с середины больше чем ключь то значит эл слева и длина >1
			ar = ar[:midleindex] // так как эл слева то офсет не нужен срезаем срез
		} else if key > mid && len(ar) > 1 { // если ключь больше чем эл с середины то значит эл справа
			ofset += midleindex + 1 // офсет нужен так как мы идем правее
			ar = ar[midleindex+1:]  // срезаем слайс
		} else if key == mid {
			return ofset + midleindex // сравнили и получили
		} else {
			break // не нашли
		}
	}
	return -1 // если нет значения
}

func binforseach[T Numeric](ar []T, key T) int { // Класический ключь тут индекс
	i, j := 0, len(ar)-1 // берем стартовый индекс как 0 и конечный как последний индекс в массиве
	for i <= j {         // пока индексы не соприкоснулись
		midleindex := int((i + j) / 2) // индекс середины
		mid := ar[midleindex]          // эл с середины
		if mid > key {                 // если эл с середины больше чем ключь то значит эл слева
			j = midleindex - 1 // индекс (конца) скидываем
		} else if key > mid { // если ключь больше чем эл с середины то значит эл справа
			i = midleindex + 1 // двигаем начальный индекс
		} else if key == mid {
			return midleindex // сравнили и получили
		}
	}
	return -1 // если нет значения
}

func binRecSearch[T Numeric](ar []T, key T) int { // Рекурсивный
	midleindex := int(len(ar) / 2) // берем индекс эл в центре
	mid := ar[midleindex]          // берем эл из центра
	if mid > key && len(ar) > 1 {  // если значение больше чем ключь то значит эл слева (и эл у нас больше 1)
		if v := binRecSearch(ar[:midleindex], key); v != -1 { // предаем слайс до нашего эл и проверяем если -1 значит у нас беды
			return v // так как эл слева то офсет не нужен
		}
		return -1 // v = -1 беда (не найдено там эл)
	} else if key > mid && len(ar) > 1 { // эл справа
		if v := binRecSearch(ar[midleindex+1:], key); v != -1 { // тоже самое только значение midle +1 что бы не захватить наш midlle
			return v + midleindex + 1 // офсет + 1 так как мы взяли больше нашего + добавили текущий так как внутри там у слайса будут индексы с 0
		}
		return -1 // не нашли
	} else if mid == key { // попали ...
		return midleindex // вернем индекс
	}
	return -1 // если длина слайса 1 эл и при этом мы не попали в строчки выше (else if mid == key)
}

func main() {
	ar := [15]int{0, 11, 18, 25, 28, 40, 47, 56, 59, 62, 81, 81, 87, 89, 94} // Массив сортированный~!
	fmt.Println(binRecSearch(ar[:], 81))
	fmt.Println(binforseach(ar[:], 81))
	fmt.Println(binforseachS(ar[:], 81))
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
