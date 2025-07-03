package gocodewars

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// MARK: - 8kyu

/*
Even or Odd
*/

func EvenOrOdd(number int) string {

	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

//---------------------------------------------------

/*
Multiply
*/

func Multiply(a, b int) int {
	return a * b
}

//---------------------------------------------------

/*
Return Negative
*/

func MakeNegative(x int) int {
	if x > 0 {
		return -x
	}
	return x
}

//---------------------------------------------------

/*
Sum of positive
*/

func PositiveSum(numbers []int) (sum int) {

	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}
	return
}

//---------------------------------------------------

/*
Reversed Strings
*/

func Solution(word string) (resultWord string) {

	for _, w := range word {
		resultWord = string(w) + resultWord
	}
	return
}

//---------------------------------------------------

/*
Convert a Number to a String!
*/

// import "strconv"

func NumberToString(n int) string {
	var result string = strconv.Itoa(n)
	return result
}

//---------------------------------------------------

/*
Convert boolean values to strings 'Yes' or 'No'.
*/

func BoolToWord(word bool) string {
	if word {
		return "Yes"
	} else {
		return "No"
	}
}

//---------------------------------------------------

/*
Square(n) Sum
*/

func SquareSum(numbers []int) int {
	var summ int

	for _, num := range numbers {
		summ = summ + num*num
	}
	return summ
}

//---------------------------------------------------

/*
Remove First and Last Character
*/

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

//---------------------------------------------------

/*
Opposite number
*/

func Opposite(value int) int {
	return -value
}

//---------------------------------------------------

/*
String repeat
*/

func RepeatStr(repetitions int, value string) string {
	result := strings.Repeat(value, repetitions)
	return result
}

//---------------------------------------------------

/*
Grasshopper - Summation
*/

func Summation(n int) int {
	summ := n * (n + 1) / 2
	return summ
}

//---------------------------------------------------

/*
Find the smallest integer in the array
*/

func SmallestIntegerFinder(numbers []int) int {
	sort.Ints(numbers)
	return numbers[0]
}

//---------------------------------------------------

/*
Remove String Spaces
*/

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

//---------------------------------------------------

/*
Counting sheep...
*/

func CountSheeps(numbers []bool) int {
	var result int
	for _, num := range numbers {
		if num {
			result++
		}
	}
	return result
}

//---------------------------------------------------

/*
Convert a String to a Number!
*/

func StringToNumber(str string) int {
	result, err := strconv.Atoi(str)

	if err != nil {
		fmt.Println("Ошибка преобразования:", err)
	}
	return result
}

//---------------------------------------------------

/*
Abbreviate a Two Word Name
*/

func AbbrevName(name string) string {
	words := strings.Split(name, " ")
	return strings.ToUpper(string(words[0][0])) + "." + strings.ToUpper(string(words[1][0])) + "."
}

//---------------------------------------------------

/*
Keep Hydrated!
*/

func Litres(time float64) int {
	return int(time) / 2
}

//---------------------------------------------------

/*
Century From Year
*/

func Century(year int) int {
	return (year + 99) / 100
}
