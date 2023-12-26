package grep

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	after   int
	before  int
	context int
	count   bool
	ignore  bool
	invert  bool
	fixed   bool
	lineNum bool
	flags   *pflag.FlagSet
)

// readFlags считывает флаги командной строки
func readFlags() {
	flags = pflag.NewFlagSet("grep", pflag.ExitOnError)
	flags.IntVarP(&after, "after", "A", 0, "print +N lines after match")
	flags.IntVarP(&before, "before", "B", 0, "print +N lines before match")
	flags.IntVarP(&context, "context", "C", 0, "print ±N lines around match")
	flags.BoolVarP(&count, "count", "c", false, "print count of lines")
	flags.BoolVarP(&ignore, "ignore-case", "i", false, "ignore case")
	flags.BoolVarP(&invert, "invert", "v", false, "invert match")
	flags.BoolVarP(&fixed, "fixed", "F", false, "fixed string")
	flags.BoolVarP(&lineNum, "line num", "n", false, "print line number")
	flags.Parse(os.Args[1:])
	if context > 0 {
		after = context
		before = context
	}
}

// takePatter считывает шаблон для поиска из аргументов командной строки
func takePattern() string {
	return flags.Arg(0)
}

// openFiles открывает файлы, переданные в аргументах командной строки,
// или стандартный ввод, если файлы не указаны и возвращает слайс
// файловых дескрипторов и ошибку
func openFiles() ([]*os.File, error) {
	args := flags.Args()[1:]
	if len(args) == 0 {
		return []*os.File{os.Stdin}, nil
	}
	files := make([]*os.File, len(args))
	for i, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			return nil, err
		}
		files[i] = file
	}
	return files, nil
}

// grep ищет в файле file строки, удовлетворяющие условиям флагов и шаблону
// и выводит их в стандартный вывод
func grep(file *os.File, pattern string) error {
	reader := bufio.NewReader(file)
	lnCounter := 0
	matchCounter := 0
	matcher, err := defineMatcher(pattern)
	if err != nil {
		return err
	}
	linesBefore := newQueue(before)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if before > 0 {
			linesBefore.push(line)
		}
		matcher.addLine(line)
		found := searchInLine(matcher, linesBefore, &lnCounter, &matchCounter)
		if found && after > 0 {
			for i := 0; i < after; i++ {
				line, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					return err
				}
				lnCounter++
				fmt.Print(line)
			}
		}
	}
	if count {
		fmt.Println(matchCounter)
	}
	return nil
}

// defineMatcher определяет какой реализацией интерфейса matcher пользоваться
// для поиска совпадений паттерна и строки
func defineMatcher(pattern string) (matcher, error) {
	if ignore {
		pattern = strings.ToLower(pattern)
	}
	if fixed {
		return &stringMatcher{pattern, ""}, nil
	}
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &reMatcher{re, ""}, nil
}

// searchInLine ищет шаблон в строке по регулярному выражению, или точному совпадению
// и печатает строку в стандартный вывод, если условия флагов совпадают
func searchInLine(
	matcher matcher,
	linesBefore *queue,
	lnCounter,
	matchCounter *int,
) bool {
	*lnCounter++
	found := matcher.match()
	if found != invert {
		*matchCounter++
		if count {
			return found != invert
		}
		if before > 0 {
			for _, line := range linesBefore.lines[:len(linesBefore.lines)-1] {
				fmt.Print(line)
			}
		}
		if lineNum {
			fmt.Printf("%d:", *lnCounter)
		}
		fmt.Print(matcher.getLine())
	}
	return found != invert
}

// closeFiles закрывает файлы, переданные в аргументах командной строки после
// завершения работы программы
func closeFiles(files []*os.File) {
	for _, file := range files {
		file.Close()
	}
}

// Grep точка входа в программу
func Grep() error {
	readFlags()
	pattern := takePattern()
	files, err := openFiles()
	if err != nil {
		return err
	}
	defer closeFiles(files)
	for _, file := range files {
		err = grep(file, pattern)
		if err != nil {
			return err
		}
	}
	return nil
}
