package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// On Windows a ^C (Control-C) or ^BREAK (Control-Break) normally cause the program to exit.
// If Notify is called for os.Interrupt, ^C or ^BREAK will cause os.Interrupt to be sent on the channel,
// and the program will not exit. If Reset is called, or Stop is called on all channels passed to Notify,
// then the default behavior will be restored.
func genchandata(ctx context.Context, chandata chan<- int) {
	for { // циклично ждем
		v := rand.Intn(100)
		select { //либо
		case <-ctx.Done(): // если пришел сигнал завершения
			break // выходим из цикла и завершаем
		default: // иначе
			chandata <- v // пишем в канал
		}
	}
}

func printer(ctx context.Context, chandata <-chan int, wg *sync.WaitGroup) {
	for { // циклично ждем
		select { //либо
		case <-ctx.Done(): // если пришел сигнал завершения
			wg.Done() // говорим что больше ничего не будем писать
			break     // выходим из цикла и завершаем
		case v := <-chandata: // или считали с канала
			fmt.Println(v) // печатаем
		}
	}
}

func main() {
	var secs int // секунд до завершения
	for {        // запрашиваем количество циклично пока не введут верное значение
		fmt.Println("time to end (sec):") // выведем что просим
		_, err := fmt.Scan(&secs)         //считываем
		if err != nil {
			fmt.Println("Error To read")
			continue // неверно ввели продолжаем просить
		}
		break // верно ввели выходим
	}
	wg := sync.WaitGroup{} // waintgorup что бы подождать печать последнего значения
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(secs)*time.Second) // создаем контекст по таймауту
	chandata := make(chan int)                                                           // канал
	go printer(ctx, chandata, &wg)                                                       // запускаем горутинки worker(ы)
	go genchandata(ctx, chandata)                                                        // запускаем горутинку для записи в канал
	wg.Wait()                                                                            // ждем вывода последнего значения
	close(chandata)
}
