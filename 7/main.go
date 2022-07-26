package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.
// для конкурентной записи sync.Map - хотя под капотом там все тот же mutex
// так что можем и классический метод с mutex запилить
func main() {
	fmt.Println("1 - решение")
	wg := sync.WaitGroup{} // WaitGroup для ожидания
	sMap := sync.Map{}     // наш sync.Map
	wg.Add(2)              // 2 горутины 2 итерации ожидания
	go func(sMap *sync.Map) {
		sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // закидываем цифры
		for _, v := range sl {
			sMap.Store(v, v) // сторимся
		}
		wg.Done() // -1 к итерации ожидания
	}(&sMap)
	go func(sMap *sync.Map) {
		sl := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} // закидываем цифры 2 горутиной
		for _, v := range sl {
			sMap.Store(v, v)
		}
		wg.Done()
	}(&sMap)
	wg.Wait() // ждем пока горутины закинуть в map
	sMap.Range(func(k, v interface{}) bool {
		fmt.Print(fmt.Sprintf("%v\t", v)) // вывод
		return true
	})
	main2()
}
func main2() {
	fmt.Println("\n2 - решение")
	wg := sync.WaitGroup{} // WaitGroup для ожидания
	wg.Add(2)
	sMap := make(map[int]int) // наш теперь обычный map
	smux := sync.Mutex{}      // нам не нужно изощряться на блокировки только по записили и т.д.
	addmap := func(k int, v int, mp map[int]int, mux *sync.Mutex) {
		mux.Lock()   // лочим mux
		mp[k] = v    // сейвим
		mux.Unlock() // разлочим
	}
	go func(mp map[int]int, smux *sync.Mutex) {
		sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // закидываем цифры
		for _, v := range sl {
			addmap(v, v, mp, smux) // сторимся через функцию
		}
		wg.Done() // -1 к итерации ожидания
	}(sMap, &smux)
	go func(mp map[int]int, smux *sync.Mutex) {
		sl := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19} // закидываем цифры 2 горутиной
		for _, v := range sl {
			addmap(v, v, mp, smux)
		}
		wg.Done()
	}(sMap, &smux)
	wg.Wait() // ждем пока горутины закинуть в map
	for _, v := range sMap {
		fmt.Print(fmt.Sprintf("%v\t", v)) // вывод
	}
}
