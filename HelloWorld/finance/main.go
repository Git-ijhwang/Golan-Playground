
package main

import (
	// "time"
    "fmt"
	"os"
	// "io"
	// "io/ioutil"
	"strings"
	"strconv"

    "github.com/360EntSecGroup-Skylar/excelize"
)

var xlsx *excelize.File

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	// dat, err := ioutil.ReadFile("./senario")
    // check(err)
	// fmt.Println(dat)

	f, err := os.Open("./senario")
    check(err)

	b1 := make([]byte, 100)
    len, err := f.Read(b1)
    check(err)

	if len <= 0 {
		fmt.Println("Error: Cannot read file")
		return
	}

	str := strings.Split(string(b1), "\n")
	tkr1 := strings.Split(str[0], ":")
	tkr2 := strings.Split(str[1], ":")
	tkr3 := strings.Split(str[2], ":")

	// fmt.Println(tkr1[0], tkr1[1])
	// fmt.Println(tkr2[0], tkr2[1])
	// fmt.Println(tkr3[0], tkr3[1])

	var tiker = [3]string {tkr1[0], tkr2[0], tkr3[0]}
	var rate [4]float64;

	fileName1 := fmt.Sprintf("%s.xlsx", tiker[0])
	fmt.Println(fileName1)
	fileName2 := fmt.Sprintf("%s.xlsx", tiker[1])
	fmt.Println(fileName2)
	fileName3 := fmt.Sprintf("%s.xlsx", tiker[2])
	fmt.Println(fileName3)

	fmt.Println(tkr1[1])
	fmt.Println(tkr2[1])
	fmt.Println(tkr3[1])

	rate[0], err = strconv.ParseFloat(tkr1[1], 64)
	if err != nil {
		fmt.Println(err)
	}
	rate[1], err = strconv.ParseFloat(tkr2[1], 64)
	if err != nil {
		fmt.Println(err)
	}
	rate[2], err = strconv.ParseFloat(tkr3[1], 64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rate[0])
	fmt.Println(rate[1])
	fmt.Println(rate[2])

	filePos1 := fmt.Sprintf("/Users/root1/Downloads/%s", fileName1)
	filePos2 := fmt.Sprintf("/Users/root1/Downloads/%s", fileName2)
	filePos3 := fmt.Sprintf("/Users/root1/Downloads/%s", fileName3)

	xlsx1, err := excelize.OpenFile (filePos1)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	xlsx2, err := excelize.OpenFile (filePos2)
	if (err != nil) {
		fmt.Println(err)
		return
	}

	xlsx3, err := excelize.OpenFile (filePos3)
	if (err != nil) {
		fmt.Println(err)
		return
	}

}