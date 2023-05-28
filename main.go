package main

import "fmt"

func main() {

	fmt.Println(findAll(32, 10)) // [8 118 334]

}

// Основная функция
func findAll(num, len int) [3]int {
	maxNum := getMaxNum(num, len)
	minNum := 0
	count := 0
	// Копируем для результата
	maxNumCopy := maxNum
	// Если максмальное число не подходит по условиям, то подходящих чисел нет
	if !checkValid(num, maxNum) {
		return [3]int{0}
	}
	// Цикл работает, пока число не будет меньше наименьшего по разрядности числа
	// Можно оптимизировать, если сразу найти наименьшее число
	for maxNum > powNum(10, len-1) {
		// Если число подходит по условиям, то мы увеличиваем счетчик и записываем результат
		// Отнимая 9, мы гарантированно не нарушем условия суммы всех чисел последовательности
		if checkValid(num, maxNum) {
			minNum = maxNum
			count += 1
			maxNum -= 9
			continue
		}
		if maxNum%100 > powNum(10, len) {
			// Если число оказалось неподходящим, то нужно уменьшить наибольший разряд и перенести значение в наименьший
			maxNum -= 9 * powNum(10, len-1)
			continue
		}
		maxNum -= 9
	}
	return [3]int{count, minNum, maxNumCopy}
}

// Получение максимального числа
func getMaxNum(num, len int) int {
	maxFirstNum := num / len
	remainder := getMask(num%len, 1)
	rawMaskNum := getMask(len, maxFirstNum)
	return rawMaskNum + remainder
}

// Маска из единиц
func getMask(len, num int) int {
	if len == 0 {
		return 0
	}
	res := 0
	for i := 1; i < len; i++ {
		res += powNum(10, i)
	}
	return (res + 1) * num
}

// Возведение числа в степень
func powNum(num, pow int) int {
	if pow == 0 {
		return 1
	}
	res := num
	for i := 1; i < pow; i++ {
		res *= num
	}
	return res
}

// Проверка на валидность
func checkValid(validNum, num int) bool {
	res := 0
	oldNum := -1
	flag := true
	for num > 0 && flag {
		res += num % 10
		if oldNum == -1 {
			oldNum = res
		}
		flag = oldNum >= num%10
		oldNum = num % 10
		num /= 10
	}
	return res == validNum && flag
}
