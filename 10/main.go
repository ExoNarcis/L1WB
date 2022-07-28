// эталонный main.go
package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
// Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
// Почему при - значениях в меньшую сторону, а при + в большую ...
// Хочу поизвращатся с указателями на срезы ... GC вышел из чата
type Rang struct {
	rang      int
	tempslice []float64
}

// простая структура где будет текущий рендж температуры и список
func (r *Rang) Pushtemp(v float64) {
	r.tempslice = append(r.tempslice, v)
}

// простая вставка

func genaddrange(vl float64, rangein **[]Rang) {
	// функция добавляет новый Rang в слайс []Rang
	// сам по себе указатель менять нельзя если передать в функцию по этому я передал указатель на указатель...
	var v float64
	if vl > 0 { // проверяем по if меньше или больше 0 значение так как сокращаем в большую сторону для - и в меньшую для +
		v = (math.Floor(vl/10) * 10) //меньшая для +
	} else {
		v = (math.Ceil(vl/10) * 10) // большая для 0
	}
	ret := append(**rangein, Rang{rang: int(v), tempslice: []float64{vl}}) // берем новый слайс
	*rangein = &ret                                                        // подменяем указатель
}

func addinrange(vl float64, rangein *[]Rang) bool {
	// функция просто добавляет в tempslice структуры Rang если входит в rang
	for k, v := range *rangein { // цикл
		delt := vl - float64(v.rang)                             //берем дельту
		if (delt < 10 && delt > 0) || (delt > -10 && delt < 0) { // если разница меньше 10 при + зеначении значит там уже есть нужный Rang
			// или если разница больше чем -10 при условии - значения
			(*rangein)[k].Pushtemp(vl) // вставим
			return true                // если вставили значение дальше ничего не делаем вернет true
		}
	}
	return false // нет нужнего rang
}

func getrange(sl []float64) []Rang {
	// функция просто генерирует нужный слайс структур из слайса ренджей
	rangein := &[]Rang{}   // инитим пустой слайс и берем указать
	for _, v := range sl { // бегаем по температуре или еще чему
		if !addinrange(v, rangein) { // если не нашел в текущих (следовательно и не вставили)
			genaddrange(v, &rangein) // вставим новый range в слайс
		}
	}
	return *rangein // вернем по значению структуру слайса которая по факту есть ссылка на массив с len, cap :D
}

func main() {
	temp := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5} // температура
	rangein := getrange(temp)                                            // получим слайс
	fmt.Println(rangein)                                                 // вывод
}
