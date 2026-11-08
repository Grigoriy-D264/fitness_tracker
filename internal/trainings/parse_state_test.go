package trainings

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

func TestParseInvalidRecordPreservesTraining(t *testing.T) {
	want := Training{Steps: 1000, TrainingType: "Бег", Duration: time.Hour, Personal: personaldata.Personal{Name: "Иван", Weight: 75, Height: 1.75}}
	for _, input := range []string{"2000,Ходьба,invalid", "0,Бег,1h", "2000,Бег,0s", "2000,Бег,1h,extra"} {
		got := want
		if err := got.Parse(input); err == nil {
			t.Fatalf("Parse(%q): ожидалась ошибка", input)
		}
		if got != want {
			t.Fatalf("Parse(%q) изменил данные: %+v", input, got)
		}
	}
}
