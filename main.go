package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter the digits of the card")

	cardNum := ""

	fmt.Fscanln(os.Stdin, &cardNum)

	// read the number from console
	digits := []int{}
	counter := 0
	for _, digitCh := range cardNum {
		if digitCh >= '0' && digitCh <= '9' {
			digit, err := strconv.Atoi(string(digitCh))
			if err != nil {
				log.Fatal("Error: ", err)
			}
			digits = append(digits, digit)
			counter++
		} else {
			if counter > 16 {
				log.Fatal("Error: the count of numbers is more than 16 (unavailable format)")
			}
		}
	}

	// luhn algo
	sum := 0
	reverse := false
	for i := len(digits) - 1; i >= 0; i-- {
		if reverse {
			digits[i] *= 2
			if digits[i] > 9 {
				digits[i] -= 9
			}
		}
		sum += digits[i]
		reverse = !reverse
	}

	// result
	res := sum%10 == 0

	if res {
		fmt.Println("the card number is valid")
	} else {
		fmt.Println("the card number is not valid")
	}
}
