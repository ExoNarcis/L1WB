package main

import (
	"errors"
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.
// Тут я просто примерно показал что без рефлексии адаптер ерунда  и лучше пилить адаптер под каждый класс (что бы меньше ошибок тянуть при изменении структур)

func TestAU(a AU) { // тест мы передали интерфейс AU который просто умеет Login  и ForgotPass
	fmt.Println(a.Login("", "", "Http"))    // просто печать
	fmt.Println(a.Login("", "", "Google"))  // просто печать
	fmt.Println(a.ForgotPass("", "Http"))   // просто печать
	fmt.Println(a.ForgotPass("", "Google")) // просто печать
}

func main() {
	a := NewAdapter(&HttpMainLogin{}, &GoogleMainLogin{})
	TestAU(a)
}

type AU interface { // сам итерфейс
	Login(login string, password string, typ string) (bool, error)
	ForgotPass(email string, typ string) (bool, error)
}

type HttpMainLogin struct { // структура для фантазии на тему как не надо делать
	//...
}

func (h *HttpMainLogin) HttpLogin(login string, password string) (bool, error) { // аля метод входа
	fmt.Println("Http Login!")
	return true, nil
}

func (h *HttpMainLogin) HttpForgotPass(email string) (bool, error) { // метод забыл пароль
	fmt.Println("Http Forgot Pass!")
	return true, nil
}

type GoogleMainLogin struct { // фантазии на тему как работать с гуглом то у него нет login и pass
	//...
}

func (h *GoogleMainLogin) GoogleLogin(email string, jwt string, GoogleID string) (bool, error) { // фантазия входа для гугла
	fmt.Println("Google Login!")
	return true, nil
}

func (h *GoogleMainLogin) GoogleForgotPass(GoogleID string) (bool, error) { // фантазия если пароль забыл
	fmt.Println("Google Forgot Pass!")
	return true, nil
}

type Adapter struct { // сам адаптер тут 2 структуры не нужны по сути ну ладно ...
	*HttpMainLogin
	*GoogleMainLogin
}

func NewAdapter(h *HttpMainLogin, g *GoogleMainLogin) *Adapter { // функция аля конструктор
	return &Adapter{h, g}
}

func (a *Adapter) Login(login string, password string, typ string) (bool, error) { // реализуем интерфейс AU и применим нужную реализацию
	if a.GoogleLogin != nil && typ == "Google" {
		//..
		return a.GoogleLogin("", "", "")
	} else if a.HttpMainLogin != nil && typ == "Http" {
		//...
		return a.HttpLogin("", "")
	} else {
		return false, errors.New("0xc000007b") // ехехе всем известная ошибка 0 ХС 5 нулей 7 бл..
	}
}

func (a *Adapter) ForgotPass(email string, typ string) (bool, error) { // подменим тоже самое для пароля
	if a.GoogleLogin != nil && typ == "Google" {
		//...
		return a.GoogleForgotPass("")
	} else if a.HttpMainLogin != nil && typ == "Http" {
		//...
		return a.HttpForgotPass(email)
	} else {
		return false, errors.New("0xc000007b")
	}
}
