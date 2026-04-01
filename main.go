package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := 0.0
	setNums, operation := getInput()
	parsionNum, err := pars(setNums)
	if err != nil {
		fmt.Println(err)
	}
	switch operation {
	case "AVG":
		result = avg(parsionNum)
	case "SUM":
		result = sum(parsionNum)
	default:
		result = med(parsionNum)
	}
	fmt.Println(result)
}

func getInput() ([]string, string) {
	inputNums := []string{}
	operations := ""
	nums := ""
	for {
		fmt.Println("Выберите операцию AVG/SUM/MED")
		fmt.Scan(&operations)
		if strings.ToUpper(operations) == "AVG" || strings.ToUpper(operations) == "SUM" || strings.ToUpper(operations) == "MED" {
			break
		}
	}
	fmt.Println("Введите числа через запятую (для остановки введи n)")
	for {
		fmt.Scan(&nums)
		if nums == "n" {
			break
		}
		inputNums = append(inputNums, nums)
	}

	return inputNums, strings.ToUpper(operations)
}

func pars(setNums []string) ([]int, error) {
	sliceInt := make([]int, 0, len(setNums))
	for _, v := range setNums {
		parts := strings.SplitSeq(v, ",")
		for part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("Ошибка конвертации числа '%s' в число %w", part, err)
			}
			sliceInt = append(sliceInt, n)
		}
	}
	return sliceInt, nil
}

func avg(pars []int) float64 {
	avg := 0
	for _, v := range pars {
		avg += v
	}
	avg = avg / len(pars)
	return float64(avg)
}

func sum(pars []int) float64 {
	sum := 0
	for _, v := range pars {
		sum += (v)
	}
	return float64(sum)
}

func med(pars []int) float64 {
	sorted := make([]int, len(pars))
	copy(sorted, pars)
	sort.Ints(sorted)
	n := len(sorted)
	if n%2 == 1 {
		return float64(sorted[n/2])
	} else {
		left := sorted[n/2-1]
		right := sorted[n/2]
		return float64(left+right) / 2.0
	}
}
