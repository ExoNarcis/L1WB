package main

import (
	"fmt"
)

//Реализовать пересечение двух неупорядоченных множеств.

func valequal1(a []int, b []int) []int { // множество это слайс
	ret := []int{}        // слайз пересечения
	for _, v := range a { // цикл по 1 масиву
		for _, o := range b { // цикл по 2 массиву внутри 1
			if v == o { // проверям величины
				ret = append(ret, v) // если = то вставляем в срез, вложенный цикл завершаем
				break
			}
		}
	}
	return ret
}

func valequal2(a map[string]int, b map[string]int) []int {
	ret := []int{}        // слайз пересечения
	for _, v := range a { // цикл по 1 масиву
		for _, o := range b { // цикл по 2 массиву внутри 1
			if v == o { // проверям величины
				ret = append(ret, v) // если = то вставляем в срез, вложенный цикл завершаем
				break
			}
		}
	}
	return ret
}

func keyequal(a map[string]int, b map[string]int) []string {
	ret := []string{}     // слайз пересечения
	for k, _ := range a { // цикл по 1 мапе
		for i, _ := range b { // цикл по 2 мапе внутри 1
			if k == i { // проверям ключи
				ret = append(ret, k) // если = то вставляем в срез, вложенный цикл завершаем
				break
			}
		}
	}
	return ret
}

func mapequal(a map[string]int, b map[string]int) map[string]int {
	ret := make(map[string]int) // мапа пересечений
	for k, v := range a {       // цикл по 1 мапе
		for i, o := range b { // цикл по 2 мапе
			if k == i && v == o { // если и ключь и значение =
				ret[k] = o // вставляем в мапу пересечений
				break
			}
		}
	}
	return ret
}

func main() {
	// данные
	mn11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mn12 := []int{2, 5, 6, 9, 15, 18}

	mn21 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5, "f": 6, "g": 7, "h": 8, "i": 9}
	mn22 := map[string]int{"a": 1, "b": 5, "c": 3, "g": 9, "i": 15, "z": 18}
	//
	fmt.Println("value in slice eq ", valequal1(mn11, mn12))
	fmt.Println("value in map eq ", valequal2(mn21, mn22))
	fmt.Println("key in map eq ", keyequal(mn21, mn22))
	fmt.Println("full map eq ", mapequal(mn21, mn22))
}
