package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// IndexHandler отдает HTML-страницу с формой
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	// чтение файла index.html из проекта и его отправка клиенту
	http.ServeFile(w, r, "index.html")
}

// UploadHandler принимает файл, конвертирует данные и сохраняет результат
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	// парсим multipart-форму (лимит 10 МБ)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// получаем файл из формы (имя поля "myFile" берётся из index.html)
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// чтение содержимого файла
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// передаём данные для конвертации
	result, err := service.Convert(string(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// локальный файл сохраняет результат
	// узнаём расширение оригинального файла
	ext := filepath.Ext(header.Filename)

	// генерируем имя файла по времени. Заменяем ":" на "-",
	// чтобы Windows не среагировал на недопустимые символы
	timeStr := time.Now().UTC().String()
	timeStr = strings.ReplaceAll(timeStr, ":", "-")
	fileName := timeStr + ext

	// создание файла на диске
	outFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	// результат конвертации пишем в файл
	_, err = outFile.WriteString(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// результат конвертации возвращаем браузеру
	_, err = w.Write([]byte(result))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
