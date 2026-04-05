package main

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
)

const output string = "Введите валюту для конвертации (USD/EUR/RMB/RUB). Для отмены расчета введите \"стоп\""

var rates = map[string]float64{
	"USD": 81.14,
	"EUR": 93.42,
	"RMB": 11.73,
	"RUB": 1.0,
}

var currencies = []string{"USD", "EUR", "RMB", "RUB"}

func errorInputCurrency(s string) error {
	if slices.Contains(currencies, s) {
		return nil
	}
	return errors.New("Ошибка ввода валюты, пожалуйста введите корректную валюту (USD/EUR/RMB/RUB)\n")
}

func getCurrency(prompt string) (string, error) {
	var value string
	fmt.Println(prompt)
	fmt.Scan(&value)
	value = strings.ToUpper(value)
	if value == "СТОП" {
		os.Exit(0)
	}
	err := errorInputCurrency(value)
	if err != nil {
		return "", err
	}
	return value, nil
}

func getAmount() (float64, error) {
	var num float64
	fmt.Println("Введите необходимую сумму для конвертации")
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, errors.New("Ошибка ввода суммы, пожалуйста введите число\n")
	}
	if num < 0 {
		return 0, errors.New("Ошибка ввода суммы, сумма не может быть отрицательной\n")
	}
	return num, nil
}

func convert(amount float64, from, to string) float64 {
	rubAmount := amount * rates[from]
	return rubAmount / rates[to]
}

func main() {
	fmt.Println("Онлайн калькулятор валюты")
	for {
		from, err := getCurrency(output)
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		sum, err := getAmount()
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		to, err := getCurrency("Введите целевую валюту (USD/EUR/RMB/RUB). Для отмены введите \"стоп\"")
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}
		result := convert(sum, from, to)
		fmt.Printf("%.2f %s = %.2f %s\n", sum, from, result, to)
	}
}
