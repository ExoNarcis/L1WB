package main

import (
	"errors"
	"fmt"
)

// Удалить i-ый элемент из слайса.
func remove2(sl *[]int, index int) error { // удаление по короче :D
	if index > len(*sl)-1 { // проверим индексы
		return errors.New("index out of range")
	}
	*sl = append(((*sl)[:index]), (*sl)[index+1:]...) // складываем слайс до эл и после него и переприсваиваем
	return nil
}

func remove(sl *[]int, index int) error { // функция удаления указатель на слайс так как меняем длину
	if index > len(*sl)-1 { // проверим индексы
		return errors.New("index out of range")
	}
	for i := index; i+1 < len(*sl); i++ { // от индекса удаления надо сместить
		(*sl)[i] = (*sl)[i+1] // смешение которое затрет нужный эл
	}
	(*sl) = (*sl)[:len(*sl)-1] // после смешения надо удалить послед эл
	return nil
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}  // слайс 1
	sl2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // слайс 2
	remove(&sl, 3)                             // удаление 1 методом
	remove2(&sl2, 3)                           // удаление 2 методом
	fmt.Println(sl)                            // печать
	fmt.Println(sl2)                           // печать
}
