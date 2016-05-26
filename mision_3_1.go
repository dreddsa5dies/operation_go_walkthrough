package main

import (
	"fmt"
	"strings"
)

func main() {
	q := "Ja ja DJaja"
	a := strings.Split(q, " ") // разбивает имя агента на строки по пробелу
	fmt.Println(a)
	b := a[0]                    //
	c := a[1]                    //
	d := a[2]                    //
	x := strings.EqualFold(b, c) // должны быть одинаковы
	y := b != strings.ToLower(c) // сравнение b и c (в нижнем регистре)
	w := strings.Index(d, b+c)
	z := strings.Index(d, b+c) == 1 && len(d) == 5 // false
	fmt.Println(len(d))
	fmt.Println(w)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(x && y && z)
}
