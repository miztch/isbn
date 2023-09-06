package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sliceAtoi(sliceArray []string) ([]int, error) {
	sliceInteger := []int{}

	// Convert a slice of strings to a slice of integers
	for _, i := range sliceArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			return sliceInteger, err
		}
		sliceInteger = append(sliceInteger, j)
	}

	return sliceInteger, nil
}

func getIsbn10(isbn13 string) (string, error) {

	// 1. Remove the first 3 characters and the last character of ISBN13 to obtain a 9-digit number.
	body := isbn13[3:12]

	// 2. Multiply each digit of the result from step 1 by 10, 9, 8, ..., 2, and calculate the sum.
	baseStrings := strings.Split(body, "")
	base, err := sliceAtoi(baseStrings)
	if err != nil {
		return "error occured", err
	}

	weight := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	sum := 0
	for i := 0; i < len(base); i++ {
		sum += base[i] * weight[i]
	}

	// 3. Obtain the check digit by subtracting the remainder when dividing the sum from step 2 by 11 from 11.
	// If the result is 11, set the check digit to 0.
	// If it's 10, set it to X.
	digit := strconv.Itoa(11 - sum%11)

	if digit == "11" {
		digit = "0"
	} else if digit == "10" {
		digit = "X"
	}

	// 4. ISBN10 is obtained by appending the check digit from step 3 to the end of the result from step 1.
	isbn10 := body + digit

	return isbn10, nil

}

func getIsbn13(isbn10 string) (string, error) {

	// 1. Add "978" to the beginning of ISBN10 string
	body := "978" + isbn10[:9]

	// 2. Sum the digits at even positions (2nd, 4th, 6th, ..., 12th)
	//    and triple the digits at odd positions (1st, 3rd, 5th, ..., 13th).
	baseStrings := strings.Split(body, "")
	base, err := sliceAtoi(baseStrings)
	if err != nil {
		return "error occured", err
	}

	sum := 0
	for i := 0; i < len(base); i++ {
		if i%2 == 0 {
			sum += base[i]
		} else {
			sum += base[i] * 3
		}
	}

	// 3. Obtain the check digit by subtracting the remainder when dividing the sum from step 2 by 10 from 10.
	// If the result is 10, set the check digit to 0.
	digit := strconv.Itoa(10 - sum%10)

	if digit == "10" {
		digit = "0"
	}

	// 4. ISBN13 string is obtained by appending the check digit from step 3 to the end of the result from step 1.
	isbn13 := body + digit

	return isbn13, nil
}

func main() {
	os.Exit(_main())
}

func _main() int {
	// Get command line arguments
	flag.Parse()
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s ISBNCode\n", os.Args[0])
		flag.PrintDefaults()
	}

	// Check arguments count
	if flag.NArg() != 1 {
		flag.Usage()
		return 1
	}

	// Get input
	var inputIsbnCode string = flag.Args()[0]

	// Check order of input ISBN code
	// If valid, then convert
	var isbnCode string
	var err error

	if len(inputIsbnCode) != 10 && len(inputIsbnCode) != 13 {
		isbnCode, err = "", errors.New("error: the order of input ISBN code is incorrect.(10 or 13 digits are allowed)")
	} else if len(inputIsbnCode) == 10 {
		isbnCode, err = getIsbn13(inputIsbnCode)
	} else if len(inputIsbnCode) == 13 {
		isbnCode, err = getIsbn10(inputIsbnCode)
	}

	// Output
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	} else {
		fmt.Println(isbnCode)
		return 0
	}
}
