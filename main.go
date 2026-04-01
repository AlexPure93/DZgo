package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	result := 0.0
	setNums, operation, err := getInput()
	if err != nil {
		fmt.Println("Invalid input nums")
	}
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
	fmt.Printf("%.2f", result)
}

func getInput() ([]string, string, error) {
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
	fmt.Scan(&nums)
	if nums == "" {
		return nil, "", errors.New("Err")
	}
	inputNums = append(inputNums, nums)
	return inputNums, strings.ToUpper(operations), nil
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
	avg := 0.0
	for _, v := range pars {
		avg += float64(v)
	}
	avg = avg / float64(len(pars))
	return (avg)
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
