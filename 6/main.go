package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины.
// По факту горутина должна сама себя остановить я так понял тут надо продемонстрировать методы ее "оповещения"
func main() {
	wg := sync.WaitGroup{}
	fmt.Println("1- quit channel")
	ch := make(chan bool) // будем крыть каналом
	go func(ch chan bool) {
		wg.Add(1)
		for { // цикл
			select {
			case <-ch: // в канал что то упало - пришло время закруглятся
				wg.Done()
				fmt.Println("signal stop work 1")
				return
			default: // иначе работаем
				// fmt.Println("work 1") // делаем вид что работаем
			}
		}

	}(ch)
	time.Sleep(time.Duration(2) * time.Second)
	close(ch) // вместо bool можно хоть interface{} использовать

	fmt.Println("2 - context")                              // контекст по фатку та же механика только обернута красиво
	ctx, cancel := context.WithCancel(context.Background()) // берем cancelfunction и сам конекст от другого контекста не обязательно background
	go func(ctx context.Context) {
		wg.Add(1)
		for {
			select {
			case <-ctx.Done(): // ctx.Done() в данном примере сработает после вызова cancel()
				wg.Done()
				fmt.Println("signal stop work 2")
				return
			default: // иначе работаем
				// fmt.Println("work 2") // делаем вид что работаем
			}
		}
	}(ctx)
	time.Sleep(time.Duration(2) * time.Second)
	cancel() // вызов функции в данном примере для остановки горутины

	fmt.Println("3 - так делать ненадо") // хотя в академ целях можно, если не верится что работает можно разкоментить вывод
	var cancfunc func()
	go func(fc *func()) {
		wg.Add(1)
		var tr bool
		tr = true
		*fc = func() {
			tr = false
		}
		for tr {
			//fmt.Println("work 3") // делаем вид что работаем
		}
		fmt.Println("signal stop work 3")
		wg.Done()
	}(&cancfunc)
	time.Sleep(time.Duration(2) * time.Second)
	cancfunc()
	wg.Wait()
	// тут можно издеватся бесконечно включая atomic у меня фанатзии не хватило придумать но методы наверное еще есть но там либо под капотом каналы 1,2
	// либо псевдо колбеки - 3
	// либо atomic если не жалко мучать чтением/записью его :( см. 3 задание там atomic
}
