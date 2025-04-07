package trainings

import (
	"fmt"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	// TODO: добавить поля
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	// TODO: реализовать функцию
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("неверный формат данных, ожидается: шаги,тип,длительность")
	}

	if _, err = fmt.Sscanf(parts[0], "%d", &t.Steps); err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}

	t.TrainingType = strings.TrimSpace(parts[1])
	if t.TrainingType != "Бег" && t.TrainingType != "Ходьба" {
		return fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	durationStr := strings.TrimSpace(parts[2])
	var dur time.Duration
	if dur, err = time.ParseDuration(durationStr); err != nil {
		return fmt.Errorf("ошибка парсинга времени: %w", err)
	}
	t.Duration = dur

	return nil
}

func (t Training) ActionInfo() (string, error) {
	// TODO: реализовать функцию
	if t.Duration <= 0 {
		return "", fmt.Errorf("некорректная продолжительность тренировки")
	}

	distance := spentenergy.Distance(t.Steps)
	speed := spentenergy.MeanSpeed(t.Steps, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Personal.Weight,
			t.Personal.Height,
			t.Duration,
		)
	default:
		return "", fmt.Errorf("неподдерживаемый тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	result := fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	)

	return result, nil
}
