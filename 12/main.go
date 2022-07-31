package main

import (
	"fmt"
	"strconv"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
// если множество это map и cat,cat ... это ключи то не выйдет, ключь уникален хотяяя...
// обычный мап не пойдет напишем свой с блекджеком и ...
func main() {
	tm1 := &TrueMap{} // первый метод с структурой невидимых ключей (скрываем реальный ключ так как при добавлении мы меняем ключь если эл такой есть)
	tm1.Init()        // инитаем
	tm1.Add("cat", 1) // добавим из списка
	tm1.Add("cat", 2)
	tm1.Add("dog", 1)
	tm1.Add("cat", 3)
	tm1.Add("tree", 1)
	fmt.Println(tm1)            // вывод всего
	fmt.Println(tm1.Get("cat")) // вывод 1 эл
	// метод 2
	ms := make(map[string][]int) // группировка
	ms["cat"] = []int{1, 2, 3}
	ms["dog"] = []int{1}
	ms["tree"] = []int{1}
	fmt.Println(ms)

}

type TrueMap struct { // свой овномап
	inmap map[string]interface{}
}

func (t *TrueMap) Init() { // инитим
	t.inmap = make(map[string]interface{})
}

func (t *TrueMap) Add(key string, val interface{}) { // добавим ключи
	vkey := key // переменная с которой мы будем выдавать виртуальные ключики
	for i := 0; true; i++ {
		if i != 0 { // если уже есть элемент с обычным ключем
			vkey = key + "`" + strconv.Itoa(i) + "`" // добавим итератор плюс ``
		}
		if t.inmap[vkey] == nil { // если вдруг такого значения нет значит добавим
			t.inmap[vkey] = val
			return
		}
	}
}

func (t *TrueMap) Get(key string) []interface{} { // получим по ключу
	ret := []interface{}{} // логика как и у add только мы тут собираем все возможные ключи до того как итератор зайдет к несуществующему эл
	vkey := key
	for i := 0; true; i++ {
		if i != 0 {
			vkey = key + "`" + strconv.Itoa(i) + "`"
		}
		if t.inmap[vkey] != nil {
			ret = append(ret, t.inmap[key])
		} else {
			return ret
		}
	}
	return nil
}

func (t *TrueMap) Del(key string) { // удаляем по ключику
	vkey := key // то же самое как и в get только удаляем все сразу
	for i := 0; true; i++ {
		if i != 0 {
			vkey = key + "`" + strconv.Itoa(i) + "`"
		}
		if t.inmap[vkey] != nil {
			delete(t.inmap, vkey)
		} else {
			return
		}
	}

}

func (t *TrueMap) Print() {
	fmt.Println(t.inmap) // просто печать реальных ключей и значений
}

func (t *TrueMap) String() string { // для вывода в fmt неявно вызывается String если такой есть
	str := "["
	for k, v := range t.inmap { // смысл окуратно снести наши `1` и т.д. из вывода
		runs := []rune(k)
		isnum := false
		offset := len(runs)
		for i := len(runs) - 1; i >= 0; i-- {
			if runs[i] == 96 {
				if !isnum {
					isnum = true
				} else {
					offset = i
				}
			}
		}
		str += "{" + string(runs[:offset]) + ":" + fmt.Sprintf("%v", v) + "},"
	}
	return str[:len(str)-1] + "]"
}
