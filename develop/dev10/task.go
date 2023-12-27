package telnet

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"github.com/spf13/pflag"
)

var (
	timeout string
	host    string
	port    string
)

// readFlags читает флаги и аргументы командной строки
func readFlags() {
	flags := pflag.NewFlagSet("telnet", pflag.ExitOnError)
	flags.StringVarP(&timeout, "timeout", "t", "10s", "timeout")
	flags.StringVarP(&host, "host", "h", "", "host")
	flags.StringVarP(&port, "port", "p", "", "port")
	flags.Parse(os.Args[1:])
}

// Telnet выполняет подключение к удалённому хосту по указанному адресу и порту
// и с указанным таймаутом (10 секунд - по умолчанию, если не указан)
func Telnet() {
	readFlags()

	dur := parseTimeout(timeout)
	ctx, cancel := context.WithTimeout(context.Background(), dur)
	defer cancel()
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), dur)
	if err != nil {
		time.Sleep(dur)
		return
	}
	defer conn.Close()

	done := make(chan struct{})
	output := make(chan []byte, 1)

	go handleMessages(ctx, conn, output, done)

	for {
		select {
		case msg := <-output:
			_, err := os.Stdout.Write(msg)
			if err != nil {
				log.Println(err)
				continue
			}
			os.Stdout.Sync()
		case <-done:
			return
		case <-ctx.Done():
			return
		}
	}
}

func handleMessages(
	ctx context.Context,
	conn net.Conn,
	output chan<- []byte,
	done chan<- struct{},
) {
	stdinReader := bufio.NewReader(os.Stdin)
	connReader := bufio.NewReader(conn)

	var wg sync.WaitGroup
	wg.Add(2)

	go sendMessages(ctx, stdinReader, conn, done)

	go recvMessages(connReader, output, done)
	wg.Wait()
}

func recvMessages(reader *bufio.Reader, output chan<- []byte, done chan<- struct{}) {
	for {
		msg, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				close(done)
				return
			}
			log.Println(err)
			return
		}
		output <- msg
	}
}

func sendMessages(ctx context.Context, reader *bufio.Reader, conn net.Conn, done chan<- struct{}) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			msg, err := reader.ReadBytes('\n')
			if err != nil {
				if err == io.EOF {
					close(done)
					return
				}
				log.Println(err)
				continue
			}
			_, err = conn.Write(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

// parseTimeout преобразует строку вида "10s" в тип Duration
func parseTimeout(timeout string) time.Duration {
	d, err := time.ParseDuration(timeout)
	if err != nil {
		log.Fatal(err)
	}
	return d
}
