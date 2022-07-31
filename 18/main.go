package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. По завершению программа должна выводить итоговое значение счетчика.

type SaveInc struct {
	i       int        // класический счетчик
	i32     int32      // счетчик прикол для atomic
	ichan   int        // счетчик с менеджером
	mux     sync.Mutex // обычный mux
	chaninc chan int   // канал
}

func (si *SaveInc) IncMux() {
	si.mux.Lock()   // лочим
	si.i++          // пишем
	si.mux.Unlock() // разлочим
}

func (si *SaveInc) IncAtomic() {
	atomic.AddInt32(&si.i32, 1) // просто атомиком просим записать
}

func (si *SaveInc) chanman(wg *sync.WaitGroup) {
	for range si.chaninc {
		si.ichan++ // пока канал не закрыт и что то приходит я буду икрементить
	}
	wg.Done()
}

func (si *SaveInc) IncChan() {
	si.chaninc <- 1 // инкримент
}

func main() {
	wg := sync.WaitGroup{}                  // для горутин wg
	wg2 := sync.WaitGroup{}                 // для менеджера каналов
	s := SaveInc{chaninc: (make(chan int))} // создаем стркутуру и канал
	wg2.Add(1)                              // подготовим wg для менеджера канала
	go s.chanman(&wg2)                      // отдаем ему
	wg.Add(3)                               // запускаем горутины и выдаем wg
	go func() {
		for i := 0; i < 20; i++ {
			s.IncMux()
			s.IncAtomic() // используем все 3 метода
			s.IncChan()
		}
		wg.Done() // снимаем
	}()
	go func() {
		for i := 0; i < 10; i++ {
			s.IncMux()
			s.IncAtomic() // используем все 3 метода
			s.IncChan()
		}
		wg.Done() // снимаем
	}()
	go func() {
		for i := 0; i < 5; i++ {
			s.IncMux()
			s.IncAtomic() // используем все 3 метода
			s.IncChan()
		}
		wg.Done() // снимаем
	}()
	wg.Wait()        // ждем горутины
	close(s.chaninc) // закрываем канал
	wg2.Wait()       // ждем пока менеджер завершится
	fmt.Println(s.i) // печать
	fmt.Println(s.i32)
	fmt.Println(s.ichan)
}
