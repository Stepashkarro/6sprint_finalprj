package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// функция Convert определяет, что пришло: текст или Морзе, и конвертирует в нужный язык
func Convert(data string) (string, error) {
	// убираем пробелы и переносы строк
	data = strings.TrimSpace(data)

	if data == "" {
		return "", errors.New("передан пустой файл")
	}

	// проверка на код Морзе - состоит только из точек, тире и пробелов.
	isMorse := true
	for _, char := range data {
		if char != '.' && char != '-' && char != ' ' && char != '\n' && char != '\r' {
			isMorse = false
			break
		}
	}

	// если это Морзе - переводим в текст
	if isMorse {
		return morse.ToText(data), nil
	}

	// иначе (обычный текст) переводим в Морзе
	return morse.ToMorse(data), nil
}
