package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	calories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0, err
	}
	return calories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 {
		return 0, fmt.Errorf("количество шагов должно быть положительным: %d", steps)
	}
	if weight <= 0 {
		return 0, fmt.Errorf("вес должен быть положительным: %g", weight)
	}
	if height <= 0 {
		return 0, fmt.Errorf("рост должен быть положительным: %g", height)
	}
	if duration <= 0 {
		return 0, fmt.Errorf("продолжительность должна быть положительной: %s", duration)
	}
	return weight * MeanSpeed(steps, height, duration) * duration.Minutes() / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return float64(steps) * (height * stepLengthCoefficient) / mInKm
}
