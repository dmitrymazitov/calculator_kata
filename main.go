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

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}

func go_to_roman(n int) string {
	res := ""

	elements := map[int]string{
		1000: "М",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}

	keys := make([]int, len(elements))

	i := 0

	for k := range elements {
		keys[i] = k
		i++
	}
	sort.Ints((keys))
	slices.Reverse(keys)

	for _, k := range keys {
		if n/k == 1 {
			res += elements[k]
		}
		n %= k
	}
	return res
}

func go_to_arabic(n string) int {

	elements := map[string]int{
		"М":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
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

func check_digits(a string) int {

	//arab_list := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	//rome_list := []string{"X", "IX", "VIII", "VII", "VI", "V", "IV", "III", "II", "I"}

	if len(a) > 0 && len(a) <= 2 && a != " " {
		value_dgt, _ := strconv.Atoi(a)
		return value_dgt
	} else {
		panic("Значение должно быть от 1 до 10")
	}

}

//first_value_dgt, err := strconv.Atoi(first_value_str)
//if err != nil {
//	panic(err)
//}
//
//second_value_str := strp_st[2]
//second_value_dgt, err := strconv.Atoi(second_value_str)
//if err != nil {
//	panic(err)
//}
//operand := strp_st[1]
//
//if operand == "+" {
//	fmt.Println(addition(first_value_dgt, second_value_dgt))
//}
//if operand == "-" {
//	fmt.Println(subtraction(first_value_dgt, second_value_dgt))
//}
//if operand == "*" {
//	fmt.Println(multiplication(first_value_dgt, second_value_dgt))
//}
//if operand == "/" {
//	fmt.Println(division(first_value_dgt, second_value_dgt))
//}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	strp_st := strings.Fields(text)

	if len(strp_st) != 3 {
		fmt.Println("Неверное значение")
	} else {

		first_value_str := strp_st[0]
		second_value_str := strp_st[2]

		a_d := check_digits(first_value_str)
		b_d := check_digits(second_value_str)

		operand := strp_st[1]

		if operand == "+" {
			fmt.Println(addition(a_d, b_d))
		}
		if operand == "-" {
			fmt.Println(subtraction(a_d, b_d))
		}
		if operand == "*" {
			fmt.Println(multiplication(a_d, b_d))
		}
		if operand == "/" {
			fmt.Println(division(a_d, b_d))
		}

	}

	s := "XCVI"
	n := 96

	fmt.Println(go_to_arabic(s))
	fmt.Println(go_to_roman(n))
}
