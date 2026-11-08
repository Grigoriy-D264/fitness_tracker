package daysteps

import (
	"testing"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

func TestParseInvalidRecordPreservesDaySteps(t *testing.T) {
	want := DaySteps{Steps: 1000, Duration: time.Hour, Personal: personaldata.Personal{Name: "Иван", Weight: 75, Height: 1.75}}
	for _, input := range []string{"2000,invalid", "0,1h", "2000,0s", "2000,1h,extra"} {
		got := want
		if err := got.Parse(input); err == nil {
			t.Fatalf("Parse(%q): ожидалась ошибка", input)
		}
		if got != want {
			t.Fatalf("Parse(%q) изменил данные: %+v", input, got)
		}
	}
}
