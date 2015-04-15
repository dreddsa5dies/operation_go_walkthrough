package main

import (
	"fmt"
	"strings"
)

func main() {

	request := Request{agent: "Jason Marshal", mission: "Operation Go"}
	// смотри операции над полями структуры
	request.agent = "Ja ja DJaja"

	location := pingSattelite(request)
	fmt.Println(request.agent)             // мои вставки контроля
	fmt.Println(accessible(request.agent)) // мои вставки контроля
	fmt.Println("Location:", location)
}

type Request struct {
	agent   string
	mission string
}

func revealLocation() string {
	return "16.7333,-169.5274"
}

func pingSattelite(request Request) string {
	if accessible(request.agent) {
		return revealLocation()
	}
	return "ACCESS DENIED"
}

func accessible(agent string) bool {
	a := strings.Split(agent, " ") // разбивает имя агента на строки по пробелу
	if len(a) != 3 {
		return false
	}
	b := a[0]                                      
	c := a[1]                                      
	d := a[2]                                      
	x := strings.EqualFold(b, c)                   // должны быть одинаковы
	y := b != strings.ToLower(c)                   // сравнение b и c (в нижнем регистре)
	z := strings.Index(d, b+c) == 1 && len(d) == 5 // индекс первого вхождения сравнение с длиной третьей строки
	return x && y && z
}
