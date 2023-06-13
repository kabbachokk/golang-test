package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

// CLI утилита должна принимать на вход строку (без пробелов) и url, куда надо слать запрос.
// Оба аргумента обязательны.

var (
	url = flag.String("u", "", "Url сервера")
	str = flag.String("s", "", "Строка")
)

func main() {
	flag.Parse()

	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() == "" {
			err := fmt.Errorf("Ошибка. Отсутствует обязательный аргумент -%s", f.Name)
			log.Fatal(err)
		}
	})

	l, err := getSubstringLength()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(l)
}

func getSubstringLength() (int, error) {
	r, err := http.Post(*url, "application/text", bytes.NewBuffer([]byte(*str)))
	if err != nil {
		err = fmt.Errorf("Ошибка соединения. %w", err)
		return 0, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		err = fmt.Errorf("Ошибка. http-статус: %d", r.StatusCode)
		return 0, err
	}

	b, _ := io.ReadAll(r.Body)
	l, _ := strconv.Atoi(string(b))
	return l, nil
}
