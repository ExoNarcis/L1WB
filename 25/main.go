package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.
func whiletime(onend chan bool, endtime int64) { // будет горутинкой с циклом
	for time.Now().Unix() < endtime { // цикл пустой
	}
	onend <- true // цикл закончился значит можно отправить true что время пройдено
	close(onend)  // закроем канал
}

func sleep(Delta int64) {
	ms := make(chan bool)                // канал для связи с горутиной
	endtime := time.Now().Unix() + Delta // дата конца
	go whiletime(ms, endtime)            // запускаем
	<-ms                                 // ждем отправки в канал
}

func main() {
	fmt.Println("Now")
	sleep(5) // спим ~ 5 сек
	fmt.Println("Now + ~5 sec")
}
