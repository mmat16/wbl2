package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую
повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""
Дополнительное задание: поддержка escape - последовательности
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)


В случае если была передана некорректная строка, функция должна возвращать
ошибку. Написать unit-тесты.

*/

// UnpackString - распаковка строки, переданной в качестве аргумента
func UnpackString(s string) (string, error) {
	var builder strings.Builder
	var prev rune

	for _, char := range s {
		if char >= '0' && char <= '9' {
			if prev == 0 {
				return "", errors.New("invalid string")
			}
			count := int(char - '0')
			builder.WriteString(strings.Repeat(string(prev), count))
			prev = 0
		} else {
			builder.WriteRune(char)
			prev = char
		}
	}
	return builder.String(), nil
}

func main() {
	string := "q2we5"
	fmt.Println(UnpackString(string))
}
