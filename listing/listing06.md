Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```shell
[3 2 3]

```
этот вывод можно объяснить тем, что после изменения первого элемента слайса в
функции `modifySlice(i []string)` происходит выделение новой памяти под слайс при
вызове `i = append(i, "4")`. Соответственно все последующие манипуляции происходят
над слайсом, выделенным в локальной памяти функции `modifySlice`. Печать в консоль
же происходит в основной функции, памяти которая принадлежит оригинальному слайсу,
в которой произошло только изменение первого элемента. Это обусловленно тем, что
слайс представлен структурой с указателем на массив и ещё двумя полями - размером
и ёмкостью. Эта структура передаётся в функцию всегда по значению. Но тк 'базовый
массив' всё таки является указателем внутри этой структуры - изменения в нём возможны.
Но при вызове функции `append` происходит создание новой структуры слайса в случае
когда ёмкости в 'базовом массиве' недостаточно для записи новых элементов (что
и происходит в нашем примере). И получается что создав новый объект слайса в теле
функции, можно потерять его при выходе из её области видимости. Поэтому стоит
всегда возвращать слайс из функции, если в её теле происходят манипуляции над ним.
