package main

import (
	"fmt"
	"math"
)

// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором

func main() {
	p1 := NewPoint(-2, -2)                 // 1 точка
	p2 := NewPoint(-4, -4)                 // 2 точка
	fmt.Println(BetweenDistanse(*p1, *p2)) // вызов функции и принт
}

func BetweenDistanse(p1 Point, p2 Point) float64 {
	x1, x2, y1, y2 := p1.GetX(), p2.GetX(), p1.GetY(), p2.GetY()        // получаем координаты
	return math.Abs(math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))) // = квадратный корень из разницы x1 и x2 в квадрате плюс разница y1 и y2 в квадрате
}

type Point struct {
	_x, _y float64 // я привык приватных членов класса называть через _
}

func (p *Point) GetX() float64 { // функция получения x писал если кто-то вынесет BetweenDistanse за пакет
	return p._x
}

func (p *Point) GetY() float64 { // функция получения y
	return p._y
}

func NewPoint(x float64, y float64) *Point { // конструктор
	return &Point{_x: x, _y: y}
}
