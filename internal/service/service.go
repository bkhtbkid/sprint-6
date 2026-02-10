package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DetectFormat(input string) (string, error) {
	isMorse := true

	index := strings.IndexFunc(input, func(r rune) bool {
		return r != '-' && r != '.' && r != ' '
	})

	if index == -1 {
		isMorse = false
	}

	if isMorse {
		return morse.ToText(input), nil
	}

	upperInput := strings.ToUpper(input)

	for _, r := range upperInput {
		if _, ok := morse.DefaultMorse[r]; !ok {
			return "", fmt.Errorf("Символ %c не поддерживается", r)
		}
	}

	return morse.ToMorse(upperInput), nil
}
