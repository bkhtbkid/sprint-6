package handlers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	text := string(data)
	var result string

	isMorse := true
	for _, char := range text {
		if !strings.ContainsRune(".- \n\r", char) {
			isMorse = false
			break
		}
	}

	if isMorse {
		result = morse.ToText(text)
	} else {
		result = morse.ToMorse(text)
	}

	dst, err := os.Create(header.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	dst.WriteString(result)

	w.Write([]byte(result))
}
