package main

import (
	"errors"
	"fmt"
	"strconv"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func changebin(i64 int64, i int, cb int) (int64, error) { // метод с кастом в строку
	if cb != 1 && cb != 0 {
		return 0, errors.New("cb != 0 or 1") // проверка на дурака
	}
	binstr := fmt.Sprintf("%b", i64) // кастим в строку
	if i > len(binstr) {
		return 0, errors.New("index out of range") // проверка на дурака 2
	}
	var last string       // остаток строки
	if i == len(binstr) { // мы меняем последний бит?
		last = ""                // остатка не будет
		binstr = binstr[0 : i-1] // срезаем последний бит
		i--                      //индекс тоже
	} else {
		last = binstr[i+1:] // не последний значит есть еще переменные
	}
	binstr = binstr[0:i] + strconv.Itoa(cb) + last // собираем строку
	v, err := strconv.ParseInt(binstr, 2, 64)      // парсим bin
	if err != nil {
		return 0, err // ошибка парсинга
	}
	return v, nil // вернем значение
}

func changebin2(i64 int64, i int, cb int) (int64, error) { // метод порязрядки
	if cb != 1 && cb != 0 {
		return 0, errors.New("cb != 0 or 1") // проверка на дурака
	}
	var ib64 int64
	ib64 = 1 << int64(len(fmt.Sprintf("%b", i64))-1-i) // сделаем n переменную где на место бита i будет 1 а все остально по 0
	if cb == 0 {
		return i64 &^ ib64, nil // если надо установить bit на 0 то применяем комбинацию and not - сброс бита (там где в ib64 стоит 1 будет сброс на 0)
		//a	b	z = a &^ b
		//0	0	  0
		//0	1	  0
		//1	0     1
		//1	1     0
	}
	return i64 | ib64, nil // иначе применяем побитовое or
}

func main() {
	var i64 int64 // переменная int 64
	i64 = 12555
	i := 0
	cb := 0
	v, _ := changebin(i64, i, cb) // вызываем функцию передаем переменную индекс и бит
	fmt.Println(v)                // вывод
	v, _ = changebin2(i64, i, cb) // вызываем функцию2 передаем переменную индекс и бит
	fmt.Println(v)                // вывод
}
