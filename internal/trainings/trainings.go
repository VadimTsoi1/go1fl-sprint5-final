package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return fmt.Errorf("invalid data format")
	}

	// Убираем пробелы в начале и в конце строки
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])
	parts[2] = strings.TrimSpace(parts[2])

	// Парсим шаги
	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps format")
	}
	t.Steps = steps

	// Проверяем тип тренировки
	t.TrainingType = parts[1]
	if t.TrainingType == "" {
		return fmt.Errorf("invalid training type")
	}

	// Парсим продолжительность
	duration, err := time.ParseDuration(parts[2])
	if err != nil || duration <= 0 {
		return fmt.Errorf("invalid duration format")
	}
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	default:
		// Возвращаем ошибку для неизвестного типа тренировки
		return "", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\n"+
			"Скорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		speed,
		calories,
	), nil
}
