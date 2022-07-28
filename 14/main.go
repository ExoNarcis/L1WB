package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.
// channel type assertion нормально не жует

func trytype2(v interface{}) bool { // под капотом по факту та же рефлексия
	if len(fmt.Sprintf("%T", v)) >= 4 { // %T = reflect
		if fmt.Sprintf("%T", v)[0:4] == "chan" { // отдельно chan так как вернется chan T где T - typename
			return true
		}
	}
	switch fmt.Sprintf("%T", v) { // остальное так
	case "int", "string", "bool":
		return true
	default:
		return false
	}
}

func trytype(v interface{}) bool { // рефлексия
	vkind := reflect.ValueOf(v).Kind() // все типы из примеров базовые
	switch vkind {
	case reflect.Int, reflect.String, reflect.Bool, reflect.Chan: // сверяем
		return true
	default:
		return false
	}
}

// типы
func main() {
	// данные
	i := 12
	s := "12"
	b := true
	ch := make(chan int)
	fl := 1.25
	fu := func() {}
	// тут я решил что пора зпихнуть это в массив any
	an := []any{i, s, b, ch, fl, fu}
	// 1-2 способ
	fmt.Println("1 way")
	for _, v := range an {
		fmt.Println(v, fmt.Sprintf("%T", v), trytype(v))
	}
	fmt.Println("2 way")
	for _, v := range an {
		fmt.Println(v, fmt.Sprintf("%T", v), trytype2(v))
	}
}
