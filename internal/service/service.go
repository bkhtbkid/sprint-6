package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectFormat(input string) (string, error) {
	index := strings.IndexFunc(input, func(r rune) bool {
		return r != '-' && r != '.' && r != ' ' && r != '\n' && r != '\r'
	})

	if index == -1 {
		return morse.ToText(input), nil
	}

	upperInput := strings.ToUpper(input)

	for _, r := range upperInput {
		if strings.ContainsRune("\n\r", r) {
			continue
		}
		if _, ok := morse.DefaultMorse[r]; !ok {
			return "", fmt.Errorf("Символ %c не поддерживается", r)
		}
	}

	return morse.ToMorse(upperInput), nil
}
