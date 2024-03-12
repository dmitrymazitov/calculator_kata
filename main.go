package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func sum(a, b int, op string) int {
	// Функция операций сложения, вычитания, умножения и деления
	if op == "+" {
		return a + b
	} else if op == "-" {
		return a - b
	} else if op == "*" {
		return a * b
	} else if op == "/" {
		if b == 0 {
			panic("На ноль делить нельзя")
		} else {
			return a / b
		}
	} else {
		return 0
	}
}

func go_to_roman(n int) string {
	// Функция конвертации арабских чисел в римские

	res := ""

	elements := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}

	keys := make([]int, len(elements))

	i := 0

	for k := range elements {
		keys[i] = k
		i++
	}

	sort.Ints(keys)
	slices.Reverse(keys)

	for _, p := range keys {
		y := n / p
		res += strings.Repeat(elements[p], y)
		n %= p

	}
	return res
}

func go_to_arabic(n string) int {
	// Функция конвертации римских чисел в арабские

	elements := map[string]int{
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}

	rome_num := n

	res := 0

	for n, i := range rome_num {
		for ro, ar := range elements {
			if string(i) == ro {
				if n == len(rome_num)-1 {
					res += ar
				} else {
					if string(i) == "C" && string(rome_num[n+1]) == "M" {
						res -= ar
						continue
					}
					if string(i) == "C" && string(rome_num[n+1]) == "D" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "C" {
						res -= ar
						continue
					}
					if string(i) == "X" && string(rome_num[n+1]) == "L" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "X" {
						res -= ar
						continue
					}
					if string(i) == "I" && string(rome_num[n+1]) == "V" {
						res -= ar
						continue
					} else {
						res += ar
					}
				}
			}
		}
	}
	return res
}

func check_type(i string) (string, int) {
	// Проверка аргументов на заданные числа. Числа должны быть арабскими (от 1 до 10) или римскими (от I до X)

	ar_list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	rm_list := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

	if slices.Contains(ar_list, i) {
		return i, 0
	} else if slices.Contains(rm_list, i) {
		return i, 1
	} else {
		panic("Неверное значение")
	}
}

func oper_check(op string) string {
	// Функция проверки операнда

	if op == "+" {
		return op
	} else if op == "-" {
		return op
	} else if op == "*" {
		return op
	} else if op == "/" {
		return op
	} else {
		panic("Неверный операнд")
	}
}

func main() {
	//Вводе данных в консоли
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	strp_st := strings.Fields(text)

	if len(strp_st) != 3 {
		fmt.Println("Неверное значение")
	} else {

		first_value_str := strp_st[0]
		second_value_str := strp_st[2]
		operand := strp_st[1]

		a_d, ir := check_type(first_value_str)
		b_d, ir2 := check_type(second_value_str)

		o_c := oper_check(operand)

		if ir == 0 && ir2 == 0 {
			to_digit, err := strconv.Atoi(a_d)
			if err != nil {
				panic(err)
			}
			to_digit2, err := strconv.Atoi(b_d)
			if err != nil {
				panic(err)
			}

			itog := sum(to_digit, to_digit2, o_c)
			fmt.Println(itog)

		} else if ir == 1 && ir2 == 1 {
			s1 := go_to_arabic(a_d)
			s2 := go_to_arabic(b_d)
			itog_ar := sum(s1, s2, o_c)
			if itog_ar < 0 {
				panic("Ошибка: Итог уравнения римских цифр является отрциательное число")
			} else {
				itog := go_to_roman(itog_ar)
				fmt.Println(itog)
			}
		} else {
			panic("Неверное значение")
		}

	}

}
