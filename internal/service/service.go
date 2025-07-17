package service

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(data string) bool {
	runes := []rune(data)
	for _, c := range runes {
		if c != '.' && c != '-' && c != ' ' {
			return false
		}
	}
	return true
}

func Convert(data string) string {
	if isMorse(data) {
		return morse.ToText(data)
	} else {
		return morse.ToMorse(data)
	}
}
