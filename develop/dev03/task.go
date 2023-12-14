package main

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"io"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

var (
	column  uint
	numeric bool
	reverse bool
	unique  bool
	args    []string
)

func init() {
	flags := pflag.NewFlagSet("sort", pflag.ExitOnError)
	flags.UintVarP(&column, "column", "k", 0, "column for sorting")
	flags.BoolVarP(&numeric, "numeric", "n", false, "sort by numeric value")
	flags.BoolVarP(&reverse, "reverse", "r", false, "reverse sort")
	flags.BoolVarP(&unique, "unique", "u", false, "don't print duplicate lines")
	flags.Parse(os.Args[1:])
	args = flags.Args()
}

// readLines reads lines from io.Reader passed as argument.
// It returns a slice of strings representing the lines.
func readLines(input io.Reader) ([]string, error) {
	if input == nil {
		return nil, errors.New("input is nil")
	}

	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines passed as slice of strings to io.Writer.
func writeLines(lines []string, output io.Writer) error {
	if output == nil {
		return errors.New("output is nil")
	}
	for _, line := range lines {
		_, err := output.Write([]byte(line + "\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

// determineAndReadInput determines the input source. If no file is passed as
// command line argument, it reads from os.Stdin, otherwise it reads from the
// file/files passed as command line arguments and returns a slice of strings
// representing the lines read.
func determineAndReadInput() ([]string, error) {
	if len(args) == 0 {
		return readLines(os.Stdin)
	}
	var lines []string
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			return nil, err
		}
		fileLines, err := readLines(file)
		if err != nil {
			return nil, err
		}
		lines = append(lines, fileLines...)
	}
	return lines, nil
}

// sortLines sorts the lines passed as slice of strings (default sort).
func sortLines(lines []string) {
	slices.Sort(lines)
}

// numericSort sorts the lines passed as slice of strings by numeric value.
func numericSort(lines []string) {
	var numbers []int
	var sortedLines []string
	for _, line := range lines {
		number, err := strconv.Atoi(line)
		if err != nil {
			sortedLines = append(sortedLines, line)
			continue
		}
		numbers = append(numbers, number)
	}
	slices.Sort(numbers)
	slices.Sort(sortedLines)
	for _, number := range numbers {
		sortedLines = append(sortedLines, strconv.Itoa(number))
	}
	copy(lines, sortedLines)
}

// reverseSort reverses the order of the lines passed as slice of strings.
func reverseSort(lines []string) {
	slices.Reverse(lines)
}

// uiniqeSort removes duplicate lines from the slice of strings passed as
// argument.
func uniqueSort(lines []string) []string {
	uniqueLines := make(map[string]bool)
	for _, line := range lines {
		uniqueLines[line] = true
	}

	if len(uniqueLines) == len(lines) {
		return lines
	}

	res := make([]string, 0, len(uniqueLines))
	for line := range uniqueLines {
		res = append(res, line)
	}
	return res
}

type item struct {
	fields []string
}

type items []*item

func (it items) Len() int {
	return len(it)
}

func (it items) Less(i, j int) bool {
	if len(it[i].fields) <= int(column-1) || len(it[j].fields) <= int(column-1) {
		return true
	}
	return it[i].fields[column-1] < it[j].fields[column-1]
}

func (it items) Swap(i, j int) {
	it[i], it[j] = it[j], it[i]
}

// columnSort sorts the lines passed as slice of strings by the column.
func columnSort(lines []string) {
	items := make(items, len(lines))
	for i, line := range lines {
		items[i] = &item{strings.Fields(line)}
	}
	sort.Sort(items)
	for i, item := range items {
		lines[i] = strings.Join(item.fields, " ")
	}
}

func main() {
	lines, err := determineAndReadInput()
	if err != nil {
		panic(err)
	}

	if unique {
		lines = uniqueSort(lines)
	}

	switch {
	case numeric:
		numericSort(lines)
	case column > 0:
		columnSort(lines)
	default:
		sortLines(lines)
	}

	if reverse {
		reverseSort(lines)
	}

	err = writeLines(lines, os.Stdout)
	if err != nil {
		panic(err)
	}
}
