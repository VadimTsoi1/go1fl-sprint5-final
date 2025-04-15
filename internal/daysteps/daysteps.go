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

func (ds *DaySteps) Parse(datastring string) error {
	// Проверяем наличие пробелов в начале или в конце строки
	if strings.HasPrefix(datastring, " ") || strings.HasSuffix(datastring, " ") {
		return fmt.Errorf("invalid data format: leading or trailing spaces in input")
	}

	// Убираем пробелы в начале и в конце строки
	datastring = strings.TrimSpace(datastring)

	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("invalid data format")
	}

	// Проверяем наличие пробелов в начале или в конце значений
	if strings.HasPrefix(parts[0], " ") || strings.HasSuffix(parts[0], " ") ||
		strings.HasPrefix(parts[1], " ") || strings.HasSuffix(parts[1], " ") {
		return fmt.Errorf("invalid data format: values contain leading or trailing spaces")
	}

	// Убираем пробелы внутри значений
	parts[0] = strings.TrimSpace(parts[0])
	parts[1] = strings.TrimSpace(parts[1])

	// Проверяем, что шаги корректны
	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return fmt.Errorf("invalid steps format")
	}
	ds.Steps = steps

	// Проверяем, что продолжительность корректна
	duration, err := time.ParseDuration(parts[1])
	if err != nil || duration <= 0 {
		return fmt.Errorf("invalid duration format")
	}
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		calories,
	), nil
}
