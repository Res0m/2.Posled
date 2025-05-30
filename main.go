package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	// fmt.Scan(&n)
	// for i := 0; i < n; i++ {
	//     fmt.Scan(&nums[i])
	// }
	file, err := os.Open("data_prog_contest_problem_2.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка открытия файла: %v\n", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "Ошибка чтения количества элементов\n")
		return
	}
	line := scanner.Text()
    line = strings.TrimSpace(line)
	n, err = strconv.Atoi(line)
	if err != nil || n <= 0 {
		fmt.Fprintf(os.Stderr, "Неверное количество элементов: %s\n", line)
		return
	}

	if !scanner.Scan() {
		fmt.Fprintf(os.Stderr, "Ошибка чтения второй строки\n")
		return
	}
	numStr := scanner.Text()

	parts := strings.Fields(numStr)
	if len(parts) != n {
		fmt.Fprintf(os.Stderr, "Количество чисел не совпадает\n")
		return
	}

	nums := make([]int, n)
	for i, part := range parts {
		val, err := strconv.Atoi(part)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка преобразования числа: %s\n", part)
			return
		}
		nums[i] = val
	}

	result := find(nums)
	if result == -1 {
		fmt.Println("NONE")
	} else {
		fmt.Println(result) // Ответ: 55
	}
}

func find(nums []int) int {
	count := make([]int, 27)
	uniq := 0
	minLen := len(nums) + 1
	l := 0

	for r, val := range nums {
		if val >= 1 && val <= 26 {
			if count[val] == 0 {
				uniq++
			}
			count[val]++
		}

		for uniq == 26 {
			currentLength := r - l + 1
			if currentLength < minLen {
				minLen = currentLength
			}

			leftNum := nums[l]
			if leftNum >= 1 && leftNum <= 26 {
				count[leftNum]--
				if count[leftNum] == 0 {
					uniq--
				}
			}
			l++
		}
	}
	if minLen == len(nums)+1 {
		return -1
	}
	return minLen
}
