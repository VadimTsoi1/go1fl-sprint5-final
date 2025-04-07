package daysteps

import (
	"fmt"
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

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных")
	}

	// Парсинг шагов
	if _, err := fmt.Sscanf(parts[0], "%d", &ds.Steps); err != nil {
		return fmt.Errorf("ошибка парсинга шагов: %w", err)
	}

	// Парсинг длительности
	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("ошибка парсинга времени: %w", err)
	}
	ds.Duration = dur

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("некорректная продолжительность активности")
	}

	distance := spentenergy.Distance(ds.Steps)
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Personal.Weight,
		ds.Personal.Height,
		ds.Duration,
	)

	if err != nil {
		return "", fmt.Errorf("ошибка расчета калорий: %w", err)
	}

	return fmt.Sprintf(
		"Количество шагов: %d\nДистанция составила: %.2f км\nВы сожгли: %.2f ккал\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
