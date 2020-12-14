package main

import (
	"fmt"
)

func main() {
	var birth_year, birth_month, birth_date int
	var current_year, current_month, current_date int

	var b int
	//fmt.Println("Input your birth year:")
	//fmt.Scanf("%d", &birth_year)
	fmt.Println("Input your birth month:")
	fmt.Scanf("%d", &birth_month)
	//fmt.Println("Input your birth date:")
	//fmt.Scanf("%d", &birth_date)

	fmt.Println("Before :", birth_month)
	//birth_month = ((birth_month-1)*30)

	if birth_month < 9 {
		b = birth_month / 2
	} else {
		b = birth_month/2 + 1
	}

	if birth_month > 2 {
		b -= 1
	}

	birth_month = (birth_month-1)*30 + b

	fmt.Println("After :", birth_month)
	//birth_year = (birth_year-1)*365
	//birth_month = (birth_month-1)*30
	//birth_date = (birth_date-1)

	fmt.Println("Input your current year:")
	fmt.Scanf("%d", &current_year)
	fmt.Println("Input your current month:")
	fmt.Scanf("%d", &current_month)
	fmt.Println("Input your current date:")
	fmt.Scanf("%d", &current_date)

	current_year = (current_year - 1) * 365
	current_month = (current_month - 1) * 30
	current_date = (current_date - 1)

	days := birth_year + birth_month + birth_date
	today := current_year + current_month + current_date
	fmt.Println("This is your days from your birth :", today-days)

}
