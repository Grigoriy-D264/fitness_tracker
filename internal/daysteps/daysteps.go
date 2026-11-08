package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	fields := strings.Split(datastring, ",")
	if len(fields) != 2 {
		return fmt.Errorf("ожидаются два поля: шаги,длительность")
	}
	steps, err := strconv.Atoi(fields[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}
	duration, err := time.ParseDuration(fields[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать длительность: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("длительность должна быть положительной: %s", duration)
	}
	ds.Steps, ds.Duration = steps, duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
