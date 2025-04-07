package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                              = 1000  // количество метров в километре
	minInH                             = 60    // количество минут в часе
	lenStep                            = 0.65  // длина одного шага в метрах
	walkingCaloriesWeightMultiplier    = 0.035 // коэффициент для веса при ходьбе
	walkingSpeedHeightMultiplier       = 0.029 // коэффициент для роста при ходьбе
	runningCaloriesMeanSpeedMultiplier = 18    // множитель средней скорости бега
	runningCaloriesMeanSpeedShift      = 1.79  // коэффициент изменения средней скорости
)

// Distance вычисляет дистанцию в километрах
func Distance(steps int) float64 {
	return float64(steps) * lenStep / mInKm
}

// MeanSpeed вычисляет среднюю скорость
func MeanSpeed(steps int, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps) / duration.Hours()
}

// WalkingSpentCalories расчёт калорий для ходьбы
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 || height <= 0 {
		return 0, errors.New("некорректные параметры веса или роста")
	}
	if duration <= 0 {
		return 0, errors.New("некорректная продолжительность")
	}

	speed := MeanSpeed(steps, duration)
	calories := (walkingCaloriesWeightMultiplier*weight +
		(speed*speed/height)*walkingSpeedHeightMultiplier*weight) *
		duration.Hours() * minInH

	return calories, nil
}

// RunningSpentCalories расчёт калорий для бега
func RunningSpentCalories(steps int, weight float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("некорректный вес")
	}
	if duration <= 0 {
		return 0, errors.New("некорректная продолжительность")
	}

	speed := MeanSpeed(steps, duration)
	calories := (runningCaloriesMeanSpeedMultiplier*speed - runningCaloriesMeanSpeedShift) * weight

	return calories, nil
}
