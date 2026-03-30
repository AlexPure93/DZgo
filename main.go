package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const output string = "Введите валюту для конвертации (USD/EUR/RMB/RUB). Для отмены расчета введите \"стоп\""
const EURinRUB float64 = 93.42
const USDinRUB float64 = 81.14
const RMBinRUB float64 = 11.73

var currency = []string{"USD", "EUR", "RMB", "RUB"}

func errorInputCurrency(s string) error {
	for _, v := range currency {
		for range v {
			if s == v {
				return nil
			}
		}
	}
	return errors.New("Ошибка ввода валюты, пожалуйста введите коректную валюту (USD/EUR/RMB/RUB)\n")
}

func main() {

	fmt.Println("Онлайн калькулятор валюты")
	for {

		sourceCurrency, err := getSourceCurrency()
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		sum, err := getNumsInput()
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		targetCurrency, err := getSourceCurrency()
		if err != nil {
			fmt.Printf("%v", err)
			continue
		}

		resultConv := CalculateConv(sourceCurrency, sum, targetCurrency)

		fmt.Printf("Ваша сумма ровна %.2f\n", resultConv)

	}
}

func getSourceCurrency() (string, error) {
	var value1 string
	fmt.Println(output)
	fmt.Scan(&value1)
	value1 = strings.ToUpper(value1)
	if value1 == "СТОП" {
		os.Exit(1)
	}
	err := errorInputCurrency(value1)
	if err != nil {
		return "", err
	}
	return value1, nil
}

func getNumsInput() (float64, error) {
	var num float64
	fmt.Println("Введите необходимую сумму для конвертации")
	_, err := fmt.Scan(&num)
	if err != nil {
		return 0, errors.New("Invalid input\n")
	}
	if num < 0 {
		return 0, errors.New("Ошибка ввода cуммы, пожалуйста введите коректное число\n")
	}
	return num, nil
}

func CalculateConv(a string, b float64, c string) float64 {
	var totalSum float64
	switch {
	case a == "USD" && c == "RUB":
		totalSum = b * USDinRUB
	case a == "USD" && c == "EUR":
		totalSum = b * (USDinRUB / EURinRUB)
	case a == "USD" && c == "RMB":
		totalSum = b * (USDinRUB / RMBinRUB)
	case a == "EUR" && c == "USD":
		totalSum = b * (EURinRUB / USDinRUB)
	case a == "EUR" && c == "RUB":
		totalSum = b * EURinRUB
	case a == "EUR" && c == "RMB":
		totalSum = b * (EURinRUB / RMBinRUB)
	case a == "RMB" && c == "USD":
		totalSum = b * (RMBinRUB / USDinRUB)
	case a == "RMB" && c == "RUB":
		totalSum = b * RMBinRUB
	case a == "RMB" && c == "EUR":
		totalSum = b * (EURinRUB / RMBinRUB)
	case a == "RUB" && c == "USD":
		totalSum = b * ((EURinRUB / USDinRUB) / EURinRUB)
	case a == "RUB" && c == "EUR":
		totalSum = b * ((USDinRUB / EURinRUB) / USDinRUB)
	case a == "RUB" && c == "RMB":
		totalSum = b * (USDinRUB / RMBinRUB) / USDinRUB
	}
	return totalSum
}
