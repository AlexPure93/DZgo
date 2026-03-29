package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {

	for {
		fmt.Println("Хотите узнать ИМТ вашего тела? да/нет")
		input := ""
		fmt.Scan(&input)
		if strings.ToLower(input) == "да" {
			kg, height := getInput()
			imt, err := calculate(kg, height)
			if err != nil {
				fmt.Println("Ошибка ввода, введите число больше 0")
				continue
			}
			result(imt)
		} else if strings.ToLower(input) == "нет" {
			fmt.Println("До свидания!")
			break
		} else {
			fmt.Println("Введите да или нет")
			continue
		}
	}
}

func getInput() (float64, float64) {
	var weight float64
	var height float64

	fmt.Println("Расчет индекса массы тела")
	fmt.Println("Введите ваш вес в кг")
	fmt.Scan(&weight)
	fmt.Println("Введите ваш рост в см")
	fmt.Scan(&height)
	return weight, height
}

func calculate(kg, height float64) (float64, error) {
	if kg <= 0 {
		return 0, errors.New("Errors input")
	}
	if height <= 0 {
		return 0, errors.New("Errors input")
	}
	height = height / 100
	imt := kg / (height * height)
	return imt, nil
}

func result(imt float64) {
	switch {
	case imt < 16:
		fmt.Println("У вас сильный дефицит веса")
	case imt < 18.5:
		fmt.Println("У вас недостаток веса")
	case imt < 25:
		fmt.Println("У вас нормальный вес")
	case imt < 30:
		fmt.Println("У вас избыточный вес")
	default:
		fmt.Println("У вас степень ожирения")
	}
}
