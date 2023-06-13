package main

import "testing"

func TestGetSubstringLength(t *testing.T) {
	s := "abcabcbb"
	r := getSubstringLength(s)
	if r != 3 {
		t.Errorf("Результат %q для строки %s не соответствует ожидаемому значению 3", r, s)
	}
}
