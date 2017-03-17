package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

var units0to19 = [20]string{
	"ноль", "одна", "две", "три", "четыре",
	"пять", "шесть", "семь", "восемь", "девять",
	"десять", "одиннадцать", "двенадцать", "тринадцать", "четырнадцать",
	"пятнадцать", "шестнадцать", "семнадцать", "восемнадцать", "девятнадцать",
}

var tens20to90 = [8]string{
	"двадцать", "тридцать", "сорок", "пятьдесят", "шестьдесят", "семьдесят", "восемьдесят", "девяносто",
}

var hundreds0to900 = [10]string{
	"", "сто", "двести", "триста", "четыреста", "пятьсот", "шестьсот", "семьсот", "восемьсот", "девятьсот",
}

func n2s99(i uint) string {
	if i > 99 {
		panic(fmt.Sprintf("Argument %d is out of range:[0..99]", i))
	}

	if i < 20 {
		return units0to19[i]
	}

	units := i % 10
	unitsStr := units0to19[units]
	if units == 0 {
		unitsStr = ""
	}

	return joinStrings(tens20to90[i/10-2], unitsStr)
}

func n2s999(i uint) string {
	if i > 999 {
		panic(fmt.Sprintf("Argument %d is out of range:[0..999]", i))
	}

	units := i % 100

	unitsStr := n2s99(i % 100)
	hundredsStr := hundreds0to900[i/100]

	if hundredsStr == "" {
		return unitsStr
	}
	if units == 0 {
		unitsStr = ""
	}

	return joinStrings(hundredsStr, unitsStr)
}

func numericToString(i uint) string {
	return n2s999(i)
}

func toTitleCase(s string) string {
	r, i := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[i:]
}

func fmtBottles(i uint) string {
	if i == 0 {
		return "бутылок"
	} else if i == 1 {
		return "бутылка"
	} else if i <= 4 {
		return "бутылки"
	} else if i <= 20 {
		return "бутылок"
	} else {
		return fmtBottles(i % 10)
	}
}

func joinStrings(s1, s2 string) string {
	if s1 != "" {
		if s2 != "" {
			return s1 + " " + s2
		}
		return s1
	}
	return s2
}

func main() {
	// Must be in range [0..999]
	maxBottles := uint(999)
	for i := maxBottles; i > 0; i-- {
		bottles := fmtBottles(i)
		numBottles := toTitleCase(numericToString(i))

		fmt.Println(numBottles, bottles, "пива на стене!")
		fmt.Println(numBottles, bottles, "пива!")
		fmt.Println("Возьми одну, пусти по кругу!")
		fmt.Println(toTitleCase(numericToString(i-1)), fmtBottles(i-1), "пива на стене!\n")
	}

	fmt.Println("Нет бутылок пива на стене!")
	fmt.Println("Нет бутылок пива!")
	fmt.Println("Пойди в магазин и купи ещё,")
	fmt.Println(toTitleCase(numericToString(maxBottles)), "бутылок пива на стене!")
}
