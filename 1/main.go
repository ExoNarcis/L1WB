package main

import (
	"fmt"
)

// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type Human struct {
	name     string
	surname  string
	lastname string
	age      int
} // Структура Human

func (h *Human) SayName() {
	fmt.Println(h.name)
}

func (h *Human) SayFullName() {
	fmt.Println(h.name + " " + h.surname + " " + h.lastname)
}

func (h *Human) SayAge() {
	fmt.Println(h.age)
}

type Human2 struct {
	name     string
	surname  string
	lastname string
	age      int
} // Структура Human2

func (h *Human2) SayName() {
	fmt.Println(h.name)
}

func (h *Human2) SayFullName() {
	fmt.Println(h.name + " " + h.surname + " " + h.lastname)
}

func (h *Human2) SayAge() {
	fmt.Println(h.age)
}

type Action struct {
	*Human // Частный случай встраивания - агрегация
	name   string
	// to do ...
} // Структура Action

type Action2 struct {
	Human2 // Частный случай встраивания - композиция
	name   string
}

type Action3 struct {
	h    Human2 // В таком случае прийдется вызывать явно, встривание методов можно описать ручками.
	name string
}
type Act interface {
	SayName()
	SayFullName()
	SayAge() // интерфейс
}

type Action4 struct {
	Act  // встраивание интерфейса
	name string
}

func main() {
	// Пример агрегации
	Hum := &Human{name: "TestAgrName", surname: "TestAgrSurName", lastname: "TestAgrLastName", age: 20}
	// Создаю отдельно для примера
	act := Action{Hum, "Act"} // такой порядок возможен только если у нас нет 2 одинаковых типов в структуре
	// Создаем Action
	act.SayName()
	act.SayFullName()
	act.SayAge()
	// Несмотря на то что в Action нет описания методов SayName, SayFullName, SayAge - они сахарно наследовались от Human
	// пример композиции - следить за Human2 не требуется
	act2 := Action2{Human2{name: "TestKompName", surname: "TestKompSurName", lastname: "TestKompLastName", age: 22}, "Act"}
	act2.SayName()
	act2.SayFullName()
	act2.SayAge()
	// Требуется явный вызов встраивание методов в подобие "наследника" не произошло
	act3 := Action3{Human2{name: "TestKompName", surname: "TestKompSurName", lastname: "TestKompLastName", age: 22}, "Act"}
	// Требуется явный вызов
	act3.h.SayAge()

	act4 := Action4{Human2{name: "TestKompName", surname: "TestKompSurName", lastname: "TestKompLastName", age: 22}, "Act"}
	// Аналогично с интерфейсами
	act4.SayName()
	act4.SayFullName()
	act4.SayAge()
	// Вывод встраивание методов возможно для безымянных членов структуры.
}
