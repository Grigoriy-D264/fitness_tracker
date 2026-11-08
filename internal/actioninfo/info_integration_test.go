package actioninfo_test

import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/Yandex-Practicum/tracker/internal/actioninfo"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
)

// Проверяет продолжение обработки после ошибок парсинга и неизвестного типа.
func TestInfoContinuesAfterInvalidRecords(t *testing.T) {
	var logs bytes.Buffer
	previousLog := log.Writer()
	log.SetOutput(&logs)
	defer log.SetOutput(previousLog)
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	defer reader.Close()
	previousStdout := os.Stdout
	os.Stdout = writer
	defer func() { os.Stdout = previousStdout }()
	parser := &trainings.Training{Personal: personaldata.Personal{Weight: 75, Height: 1.75}}
	actioninfo.Info([]string{"bad", "6000,Плавание,1h", "6000,Бег,1h"}, parser)
	writer.Close()
	output, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}
	want := "Тип тренировки: Бег\nДлительность: 1.00 ч.\nДистанция: 4.72 км.\nСкорость: 4.72 км/ч\nСожгли калорий: 354.38\n"
	if string(output) != want {
		t.Fatalf("вывод = %q, ожидается %q", output, want)
	}
	if strings.Count(strings.TrimSpace(logs.String()), "\n") != 1 {
		t.Fatalf("ожидались две записи об ошибках: %s", logs.String())
	}
}
