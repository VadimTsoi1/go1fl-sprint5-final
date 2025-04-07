package actioninfo

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			println("Ошибка парсинга:", err.Error())
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			println("Ошибка обработки:", err.Error())
			continue
		}

		println(info)
	}
}
