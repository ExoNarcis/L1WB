package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.
// ОДЗ: Сложение "этих" чисел коммутативная операция.
func main() {
	arr := []int{2, 4, 6, 8, 10}        // Массив
	chanout := make(chan int, len(arr)) // сегодня в шоу канал можно было result сделать atomic но это скучно и не интересно зато быстрее
	// буфер у него по длине массива что бы не блочить чтение, да это затраты можно и без них но задача маленькая... пойдет.
	result := 0             // переменная результат
	wg := sync.WaitGroup{}  // waitgorup как в 2 задаче только для того что бы убедиться что все все записали в канал
	for _, v := range arr { // сахарный цикл по массиву
		wg.Add(1)        // как и тут счетчик +1
		go func(v int) { //горутинка
			chanout <- v * v // выводим в канал
			wg.Done()        // счетчик - 1
		}(v)
	}
	wg.Wait()      // ожидаем завершения записи в буф. канал
	close(chanout) // закроем его range нормально читает такие каналы
	for v := range chanout {
		result += v // считываем и складываем
	}
	fmt.Println(result) // вывод
	main2()             // atomic
	main3()             // mutex
}

func main2() {
	arr := []int{2, 4, 6, 8, 10} // Массив
	// буфер у него по длине массива что бы не блочить чтение, да это затраты можно и без них но задача маленькая... пойдет.
	var result int32        // переменная результат
	wg := sync.WaitGroup{}  // waitgorup как выше только для просто успешного выполнения
	for _, v := range arr { // сахарный цикл по массиву
		wg.Add(1)        // как и тут счетчик +1
		go func(v int) { //горутинка
			atomic.AddInt32(&result, int32(v*v)) // магия atomic
			wg.Done()                            // счетчик - 1
		}(v)
	}
	wg.Wait()           // ожидаем завершения записи в буф. канал
	fmt.Println(result) // вывод
}

func main3() {
	arr := []int{2, 4, 6, 8, 10} // Массив
	// буфер у него по длине массива что бы не блочить чтение, да это затраты можно и без них но задача маленькая... пойдет.
	var result int
	mux := sync.Mutex{}     // переменная результат
	wg := sync.WaitGroup{}  // waitgorup как выше только для просто успешного выполнения
	for _, v := range arr { // сахарный цикл по массиву
		wg.Add(1)        // как и тут счетчик +1
		go func(v int) { //горутинка
			preresult := v * v          // считаем мы конкурентно
			mux.Lock()                  // Блокируем
			result = result + preresult // а записываем нет
			mux.Unlock()                // разблокируем
			wg.Done()                   // счетчик - 1
		}(v)
	}
	wg.Wait()           // ожидаем завершения записи в буф. канал
	fmt.Println(result) // вывод
}
