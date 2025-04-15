package main

import (
	"fmt"

	"github.com/Yandex-Practicum/tracker/internal/actioninfo"
	"github.com/Yandex-Practicum/tracker/internal/daysteps"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/trainings"
)

func main() {
	person := personaldata.Personal{
		Name:       "Витя",
		Weight:     84.6,
		Height:     187,  // Рост в сантиметрах
		HeightUnit: "cm", // Явное указание единиц
	}

	// Дневная активность
	input := []string{
		"678,50m",
		"792,1h14m",
		"1078,1h30m",
		"7830,2h40m",
		",3456",
		"12:40:00,3456",
		"something is wrong",
	}

	fmt.Println("Активность в течение дня")

	dayStepsParser := &daysteps.DaySteps{Personal: person}
	dayStepsParser.Print()

	actioninfo.Info(input, dayStepsParser)

	// Тренировки
	actions := []string{
		"3456,Ходьба,3h00m",
		"something is wrong",
		"678,Бег,5m",
		"1078,Бег,10m",
		",3456,Ходьба",
		"7892,Ходьба,3h10m",
		"15392,Бег,45m",
	}

	fmt.Println("\nЖурнал тренировок")

	trainingParser := &trainings.Training{Personal: person}
	actioninfo.Info(actions, trainingParser)
}
