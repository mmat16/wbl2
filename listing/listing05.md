Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
error
```
тип `error` является интерфейсом. хотя его значение по умолчанию является `nil`,
в данном случае после вызова функции `test()`он так же кастится к типу `customError`, 
то есть значение поля тип этого интерфейса уже не будет равно nil и в таком 
случае проверка на `nil` не будет вычисляться в `true`
