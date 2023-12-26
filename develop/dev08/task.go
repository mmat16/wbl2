package minishell

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

// Shell запускает интерактивный сеанс
func Shell() {
	fmt.Println("Starting minishell at ", time.Now(), "...\nType \\quit to exit")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(">> ")

		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		input := scanner.Text()
		if input == "\\quit" {
			break
		}

		err = execCmd(input)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
	}
}

// execCmd выполняет команду
func execCmd(input string) error {
	args := strings.Fields(input)
	if len(args) == 0 {
		return nil
	}

	switch args[0] {
	case "cd":
		return cd(args)
	case "pwd":
		return pwd(args)
	case "echo":
		fmt.Println(strings.Join(args[1:], " "))
		return nil
	case "kill":
		return kill(args)
	case "ps":
		return ps(args)
	default:
		return errors.New("Unknown command: " + args[0])
	}
}

// cd смена директории
func cd(args []string) error {
	if len(args) != 2 {
		return errors.New("Error: invalid arguments count for cd")
	}
	return os.Chdir(args[1])
}

// pwd показать путь до текущего каталога
func pwd(args []string) error {
	if len(args) != 1 {
		return errors.New("Error: invalid arguments count for pwd")
	}
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	fmt.Println(dir)
	return nil
}

// kill "убить" процесс, переданный в качесте аргумента
func kill(args []string) error {
	if len(args) != 2 {
		return errors.New("Error: invalid arguments count for kill")
	}
	pid := args[1]
	return exec.Command("kill", pid).Run()
}

// ps выводит общую информацию по запущенным процессам
func ps(args []string) error {
	var output []byte
	var err error
	if len(args) == 1 {
		cmd := exec.Command("ps")
		output, err = cmd.Output()
	} else {
		cmd := exec.Command("ps", args[1:]...)
		output, err = cmd.Output()
	}
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	return nil
}
