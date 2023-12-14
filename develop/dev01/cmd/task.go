package main

/*
Создать программу печатающую точное время с использованием NTP -библиотеки.
Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием
этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR
и возвращать ненулевой код выхода в OS
*/

import (
	"fmt"

	"myTime/pkg/ntp"
)

// main is the entry point of the program. Prints the current time and the exact
// time.
const ntpPool = "0.beevik-ntp.pool.ntp.org"

func main() {
	time, err := ntp.GetPreciseTime(ntpPool)
	if err != nil {
		panic(err)
	}
	fmt.Println("Exact time:", time.Local())
}
