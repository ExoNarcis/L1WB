package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"time"
)

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

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

func printer(ctx context.Context, chandata <-chan int) {
	for { // циклично ждем
		select { //либо
		case <-ctx.Done(): // если пришел сигнал завершения
			break // выходим из цикла и завершаем
		case v := <-chandata: // или считали с канала
			fmt.Println(v) // печатаем
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // создаем контекст с функцией отмены
	var workcount int                                       // количество воркеров
	for {                                                   // запрашиваем количество циклично пока не введут верное значение
		fmt.Println("Workers count:")  // выведем что просим
		_, err := fmt.Scan(&workcount) //считываем
		if err != nil {
			fmt.Println("Error To read")
			continue // неверно ввели продолжаем просить
		}
		break // верно ввели выходим
	}
	chandata := make(chan int) // канал
	for i := 0; i < workcount; i++ {
		go printer(ctx, chandata) // запускаем горутинки worker(ы)
	}
	go genchandata(ctx, chandata)      // запускаем горутинку для записи в канал
	osSig := make(chan os.Signal)      // ждем сигнал
	signal.Notify(osSig, os.Interrupt) // просим переслать в канал сигнал os.Interrupt
	<-osSig                            // ждем сигнала
	cancel()                           // вызываем закрытие горутин
	time.Sleep(time.Second)            // может вывод еще успеет 1-2 значения упасть
	fmt.Println("Closing 5 sec...")    // зактроем приложение через 5 сек
	time.Sleep(time.Duration(5) * time.Second)
	close(chandata)
}
