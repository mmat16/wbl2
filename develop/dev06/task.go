package cut

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

/*
Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

var (
	separated bool
	delimiter string
	fields    []int
	flags     *pflag.FlagSet
)

// readFlags - чтение флагов
func readFlags() {
	flags = pflag.NewFlagSet("cut", pflag.ExitOnError)
	flags.BoolVarP(&separated, "separated", "s", false, "только строки с разделителем")
	flags.StringVarP(&delimiter, "delimiter", "d", "\t", "использовать другой разделитель")
	flags.IntSliceVarP(&fields, "fields", "f", []int{}, "выбрать поля (колонки)")
	flags.Parse(os.Args[1:])
}

// openFiles - открытие файлов
func openFiles() ([]*os.File, error) {
	if len(flags.Args()) == 0 {
		return []*os.File{os.Stdin}, nil
	}
	files := make([]*os.File, len(flags.Args()))
	for i, filename := range flags.Args() {
		file, err := os.Open(filename)
		if err != nil {
			return nil, err
		}
		files[i] = file
	}
	return files, nil
}

// closeFiles - закрытие файлов
func closeFiles(files []*os.File) {
	for _, file := range files {
		file.Close()
	}
}

// Cut - реализация утилиты cut. Открывает файлы, если они указаны, иначе читает
// из stdin. Разбивает строки по разделителю (TAB или специализированный в
// флаге) на колонки и выводит запрошенные (или все - по умолчанию).
// Возвращает ошибку, если таковая возникла
func Cut() error {
	readFlags()
	files, err := openFiles()
	if err != nil {
		return err
	}
	defer closeFiles(files)
	for _, file := range files {
		err = cutFile(file)
		if err != nil {
			return err
		}
	}
	return nil
}

// cutFile - реализация утилиты cut для одного файла
func cutFile(file *os.File) error {
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}
		if separated && !strings.Contains(line, delimiter) {
			continue
		}
		splittedFields := strings.Split(line, delimiter)
		if len(fields) == 0 {
			fmt.Print(line)
			continue
		}
		for _, field := range fields {
			if field > len(splittedFields) {
				continue
			}
			fmt.Print(splittedFields[field-1], delimiter)
		}
		fmt.Println()
	}
	return nil
}
