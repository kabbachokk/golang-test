package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

// HTTP-сервер обслуживает только один endpoint (пускай это будет /api/substring).
// В этот endpoint через CLI надо отправлять строки без пробелов.
// При получении данных, HTTP-сервер должен найти длину самой длинной подстроки без повторяющихся символов и затем отправить ответ,
// который выведет CLI утилита. Пример: abcabcbb -> abc, bbbb -> b, pwwkew -> wke.

// Если понадобится взять адрес с коммандной строки
// var httpAddr = flag.String("http", ":8080", "Адрес")
var httpAddr string = ":8383"

// (×_×)
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	// flag.Parse()
	http.HandleFunc("/api/substring", substringHandler)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}

func substringHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	l := getSubstringLength(string(b))

	_, err = w.Write([]byte(strconv.Itoa(l)))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func getSubstringLength(s string) int {
	var r int = 0
	visited := make(map[byte]int)

	// окном i:j идем по строке
	for i, j := 0, 0; j < len(s); j++ {
		val, ok := visited[s[j]]
		if ok { // нашли повторяющийся символ
			i = max(i, val) // двигаем i до индекса повторяющегося символа
		}

		r = max(r, j-i+1)     // длина текущего окна
		visited[s[j]] = j + 1 // добавляем в пройденные
	}
	return r
}
