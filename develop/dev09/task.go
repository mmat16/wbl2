package wget

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dnephin/pflag"
)

/*
Реализовать утилиту wget с возможностью скачивать сайты целиком.
*/
var (
	mirror bool
	output string
	flags  *pflag.FlagSet
)

// readFlags - чтение флагов
func readFlags() {
	flags = pflag.NewFlagSet("wget", pflag.ExitOnError)
	flags.BoolVarP(&mirror, "mirror", "m", false, "mirror")
	flags.StringVarP(&output, "output", "O", "", "output")
	flags.Parse(os.Args[1:])
}

// Wget - основная функция, выполняет чтение флагов командной строки и
// скачивание сайта, либо файла с сайта
func Wget() {
	readFlags()
	if flags.NArg() != 1 {
		return
	}

	url := flags.Args()[0]
	if output == "" {
		output = url[strings.LastIndex(url, "/")+1:]
	}

	if mirror {
		output = output + ".html"
	}

	err := mirrorSite(url, output)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Done")
}

// mirrorSite - скачивание сайта целиком
func mirrorSite(url, fileName string) error {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("got status code " + resp.Status)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
