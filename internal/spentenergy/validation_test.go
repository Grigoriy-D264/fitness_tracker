package spentenergy

import (
	"testing"
	"time"
)

func TestCaloriesRejectNonPositiveDuration(t *testing.T) {
	for _, duration := range []time.Duration{0, -time.Hour} {
		for _, calculate := range []func(int, float64, float64, time.Duration) (float64, error){RunningSpentCalories, WalkingSpentCalories} {
			calories, err := calculate(6000, 75, 1.75, duration)
			if err == nil || calories != 0 {
				t.Fatalf("duration=%v: calories=%v, err=%v", duration, calories, err)
			}
		}
	}
}

func TestMeanSpeedRejectsNegativeSteps(t *testing.T) {
	if speed := MeanSpeed(-1000, 1.75, time.Hour); speed != 0 {
		t.Fatalf("скорость при отрицательных шагах = %v", speed)
	}
}
