package main

import (
	"time"
    "fmt"
	"os"
	"strings"
	"strconv"
	"bufio"

    "github.com/360EntSecGroup-Skylar/excelize"
)

var xlsx *excelize.File

var sell float64
var stkCnt int
var curSeed float64
var accumVal float64
var preStkCnt int

var lowBase float64;
var highBase float64;
var seed float64 = 1000.0
var periodSeed float64
var initSeed float64 = 30000.0
var firstSeed float64

var myTotalAsset float64 = 0.0
var myFinalAsset float64

var numFile int
var fileasName [5]string 
var startDate string
var endDate string

var stkMdd float64
var stkHighVal float64
var stkLowVal float64
var myMdd float64
var myHighVal float64
var myLowVal float64

func writeExcelFloat(xlsx *excelize.File, sheetName string, index int, row rune, value float64) {
	pos:=fmt.Sprintf("%s%d", string(row), index)
	xlsx.SetCellValue(sheetName, pos, value)
	return;
}

func writeExcelInt(xlsx *excelize.File, sheetName string, index int, row rune, value int) {
	pos:=fmt.Sprintf("%s%d", string(row), index)
	xlsx.SetCellValue(sheetName, pos, value)
	return;
}

func writeExcelString(xlsx *excelize.File, sheetName string, index int, row rune, value string) {
	pos:=fmt.Sprintf("%s%d", string(row), index)
	xlsx.SetCellValue(sheetName, pos, value)
	return;
}

func convertDate(cel string) time.Time {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)

	str := strings.Split(cel, "-")
	date := "20"+str[2]+"-"+str[0]+"-"+str[1]
	t, _ := time.Parse(layoutISO, date)

	return t
}

func writeBasicInfo (xlsx *excelize.File, sheet string, index int, row []string) {
	baseCol := 'B'
	writeExcelInt(xlsx, sheet, index, baseCol, index-1)
	writeExcelString(xlsx, sheet, index, baseCol+1, row[0])
	// fmt.Println(row[0])
	writeExcelString(xlsx, sheet, index, baseCol+2, row[4])
	// fmt.Println(row[4])
}

func writeFinalInfo (xlsx *excelize.File, tiker string, sheet string, index int, new int) {

	baseCol := 'A'
	if (new == 0) {
		writeExcelString(xlsx, sheet, index-1, baseCol+1, tiker)
		writeExcelString(xlsx, sheet, index, baseCol+1, "Total Investing Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myTotalAsset)
		writeExcelString(xlsx, sheet, index, baseCol+2, "Final Asset Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+2, myFinalAsset)
		writeExcelString(xlsx, sheet, index, baseCol+3, "Yield")
		writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myFinalAsset-myTotalAsset)/myTotalAsset*100))

		writeExcelString(xlsx, sheet, index, baseCol+4, "Stock MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+4, stkMdd)

		writeExcelString(xlsx, sheet, index, baseCol+5, "My Asset MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+5, myMdd)

	} else {

		baseCol = baseCol+rune(new*8)
		writeExcelString(xlsx, sheet, index-1, baseCol, tiker)
		writeExcelString(xlsx, sheet, index, baseCol, "Total Investing Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol, myTotalAsset)
		writeExcelString(xlsx, sheet, index, baseCol+1, "Final Asset Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myFinalAsset)
		writeExcelString(xlsx, sheet, index, baseCol+2, "Yield")
		writeExcelString(xlsx, sheet, index+1, baseCol+2, fmt.Sprintf("%f%%",(myFinalAsset-myTotalAsset)/myTotalAsset*100))

		writeExcelString(xlsx, sheet, index, baseCol+3, "Stock MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+3, stkMdd)
		writeExcelString(xlsx, sheet, index, baseCol+4, "My Asset MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+4, myMdd)
	}
}

func calcMdd (curVal float64, newIndex int) {

	if (newIndex == 2) {
		stkMdd = 0
		stkHighVal = curVal
		stkLowVal = curVal

		myMdd = 0
		myHighVal = curVal
		myLowVal = curVal
	} else {

		if stkHighVal < curVal {
			stkHighVal = curVal
			stkLowVal = curVal
		}

		if stkLowVal > curVal {
			stkLowVal = curVal
		} else {
			//Calc MDD
			if (stkHighVal == stkLowVal) {
			}else if (stkHighVal > stkLowVal) {
				tmp:=(stkHighVal - stkLowVal) / stkHighVal
				if (stkMdd == 0) {
					stkMdd = tmp
				} else {
					if (stkMdd < tmp ) {
						stkMdd = tmp
					}
				}
				stkHighVal = stkLowVal
			} 
		}

		if myHighVal < (curVal*float64(preStkCnt)) {
			myHighVal = curVal*float64(preStkCnt)
			myLowVal = curVal*float64(preStkCnt)
		}

		if myLowVal > (curVal * float64(preStkCnt)) {
			myLowVal = curVal*float64(preStkCnt)
			// fmt.Println(curVal, preStkCnt, myLowVal)
		} else {
			//Calc MDD
			if (myHighVal == myLowVal) {
			}else if (myHighVal > myLowVal) {
				// myMdd = (myHighVal - myLowVal) / myHighVal
				tmp:=(myHighVal - myLowVal) / myHighVal

				if (myMdd == 0.0) {
					myMdd = tmp
				}else {
					if (myMdd < tmp) {
						myMdd = tmp
					}
				}
				myHighVal = myLowVal
			} 
		}
	}
}

/* Senario #1*/
func senario_1 (xlsx *excelize.File, curVal float64, newIndex int, sheet string) {

	var base	float64 = 1.02
	const low	= 0.9
	const high	= 1.1

	if newIndex%52 == 0 && newIndex>3 {
		base = base-0.002
		seed = seed*0.9
		// fmt.Println(seed)
	}

	if (newIndex == 2) {
		curSeed = initSeed
		myTotalAsset = myTotalAsset + curSeed
		stkCnt := int(float64(curSeed)/curVal)


		writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

		preStkCnt = stkCnt
		writeExcelInt(xlsx, sheet, newIndex, 'F', preStkCnt)

		writeExcelInt(xlsx, sheet, newIndex, 'G', stkCnt)

		accumVal := float64(stkCnt)*curVal
		myFinalAsset = accumVal
		writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

		lowBase = (accumVal*base)*low
		writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

		highBase = (accumVal*base)*high
		writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase)

		curSeed = seed

	} else {
		// fmt.Println(preStkCnt)
		/* pre Calculate value of my al stock */
		totPreVal := curVal*float64(preStkCnt)
		myTotalAsset = myTotalAsset + seed

		if totPreVal > highBase {
			/* Sell */
			writeExcelString(xlsx, sheet, newIndex, 'A', "SELL")

			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

			stkCnt = int((totPreVal - highBase)/curVal)
			writeExcelInt(xlsx, sheet, newIndex, 'F', 0-stkCnt)

			preStkCnt = int(preStkCnt) - stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			sell = float64(stkCnt)*curVal
			curSeed = sell+curSeed

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase)

		} else {
			/* Buy */
			writeExcelString(xlsx, sheet, newIndex, 'A', "BUY")

		// fmt.Println(curSeed)
			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

			stkCnt = int(float64(curSeed)/curVal)
		// fmt.Println(stkCnt, curSeed, curVal)
			writeExcelInt(xlsx, sheet, newIndex, 'F', stkCnt)

			preStkCnt = int(preStkCnt) + stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase) 
			curSeed = seed
		}
	}
}

/* Senario #2*/
func senario_2 (xlsx *excelize.File, curVal float64, newIndex int, sheet string) {

	var base	float64 = 1.02
	const low	= 0.8
	const high	= 1.2

	if newIndex%52 == 0 && newIndex>3 {
		base = base-0.002
		seed = seed*0.9
		// fmt.Println(seed)
	}

	if (newIndex == 2) {
		curSeed = initSeed
		myTotalAsset = myTotalAsset + curSeed
		stkCnt := int(float64(curSeed)/curVal)

		// writeExcelFloat(xlsx, sheet, newIndex, 'L', myTotalAsset)
		writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

		preStkCnt = stkCnt
		writeExcelInt(xlsx, sheet, newIndex, 'F', preStkCnt)

		writeExcelInt(xlsx, sheet, newIndex, 'G', stkCnt)

		accumVal := float64(stkCnt)*curVal
		myFinalAsset = accumVal
		writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

		lowBase = (accumVal*base)*low
		writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

		highBase = (accumVal*base)*high
		writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase)

		curSeed = (curSeed-(float64(stkCnt)*curVal))+seed

		/*
		stkMdd = 0
		stkHighVal = curVal
		stkLowVal = curVal

		myMdd = 0
		myHighVal = curVal
		myLowVal = curVal
		*/

	} else {
		// fmt.Println(preStkCnt)
		/* pre Calculate value of my al stock */
		totPreVal := curVal*float64(preStkCnt)
		myTotalAsset = myTotalAsset + seed

		writeExcelFloat(xlsx, sheet, newIndex, 'L', myTotalAsset)

		/*
		if stkHighVal < curVal {
			stkHighVal = curVal
			stkLowVal = curVal
		}

		if stkLowVal > curVal {
			stkLowVal = curVal
		} else {
			//Calc MDD
			if (stkHighVal == stkLowVal) {
			}else if (stkHighVal > stkLowVal) {
				tmp:=(stkHighVal - stkLowVal) / stkHighVal
				if (stkMdd == 0) {
					stkMdd = tmp
				} else {
					if (stkMdd < tmp ) {
						stkMdd = tmp
					}
				}
				stkHighVal = stkLowVal
			} 
		}

		if myHighVal < (curVal*float64(preStkCnt)) {
			myHighVal = curVal*float64(preStkCnt)
			myLowVal = curVal*float64(preStkCnt)
		}

		if myLowVal > (curVal * float64(preStkCnt)) {
			myLowVal = curVal*float64(preStkCnt)
			// fmt.Println(curVal, preStkCnt, myLowVal)
		} else {
			//Calc MDD
			if (myHighVal == myLowVal) {
			}else if (myHighVal > myLowVal) {
				// myMdd = (myHighVal - myLowVal) / myHighVal
				tmp:=(myHighVal - myLowVal) / myHighVal

				if (myMdd == 0.0) {
					myMdd = tmp
				}else {
					if (myMdd < tmp) {
						myMdd = tmp
					}
				}
				myHighVal = myLowVal
			} 
		}
		*/

		if totPreVal < highBase && totPreVal > lowBase{
			writeExcelString(xlsx, sheet, newIndex, 'A', "-")
			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)
			writeExcelInt(xlsx, sheet, newIndex, 'F', 0)
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)
			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)
			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase)

			curSeed = curSeed+seed
		}else if totPreVal > highBase {
			/* Sell */
			writeExcelString(xlsx, sheet, newIndex, 'A', "SELL")

			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

			if ((curSeed+(totPreVal-highBase)) > totPreVal/2) {
				tmp := (curSeed+(totPreVal-highBase)) - curVal/2
				stkCnt = int(tmp /curVal)
				// fmt.Println("Case 1", tmp, stkCnt)
			} else {
				stkCnt = int((totPreVal - highBase)/curVal)
				// fmt.Println("Case 2", stkCnt)
			}
			stkCnt = 0;
			writeExcelInt(xlsx, sheet, newIndex, 'F', 0-stkCnt)

			preStkCnt = int(preStkCnt) - stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			sell = float64(stkCnt)*curVal
			curSeed = curSeed+sell+seed

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase)

		} else {
			/* Buy */
			writeExcelString(xlsx, sheet, newIndex, 'A', "BUY")

		// fmt.Println(curSeed)
			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

			if (lowBase - totPreVal) > curSeed {
				stkCnt = int(float64(curSeed)/curVal)
			}else {
				tmp:=lowBase - totPreVal
				if tmp < curVal {
					stkCnt = 0
				} else {
					stkCnt = int(tmp/curVal)
				}

			}

			curSeed = curSeed - float64(stkCnt)*curVal
		// fmt.Println(stkCnt, curSeed, curVal)
			writeExcelInt(xlsx, sheet, newIndex, 'F', stkCnt)

			preStkCnt = int(preStkCnt) + stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase) 
			curSeed = curSeed + seed
		}
	}
}

func checkConfig() {
	i:=0
	fmt.Println("===============================");
	fmt.Println("Init Seed : ", initSeed)
	fmt.Println("Monthly Seed : ", seed)
	fmt.Println("Ticker File Num : ", numFile)
	for i<numFile {
		fmt.Println("File Name : ", fileasName[i])
		i++
	}
	fmt.Println("Start Date : ", startDate)
	fmt.Println("End Date : ", endDate)
	fmt.Println("===============================");
}

func readFile(name string) {

	f, err := os.Open(name)
    check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

			// scanner.Split(bufio.ScanWords)
	numOfFile:=0
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == '#'{
			continue
		}

		ruleData := strings.Split(line, ":")
		if ( ruleData[0] == "INIT_SEED_MONEY") {
			val, _ := (strconv.Atoi(ruleData[1]))
			firstSeed = float64(val)
			continue;
		}
		if ( ruleData[0] == "PERIOD_SEED_MONEY") {
			val, _ := (strconv.Atoi(ruleData[1]))
			periodSeed = float64(val)
			continue;
		}

		if (ruleData[0] == "NUM_FILE") {
			numFile, _ = strconv.Atoi(ruleData[1])
			continue;
		}

		if (ruleData[0] == "FILE_NAME") {
			// fmt.Println(ruleData)
			fileasName[numOfFile] = ruleData[1]
			numOfFile++

			continue;
		}
		if (ruleData[0] == "start_date") {
			startDate = ruleData[1]
			continue;
		}

		if (ruleData[0] == "end_date") {
			endDate = ruleData[1]
			continue;
		}
	}

	return
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main () {

	goFlag := 0
	skipWeek := 1
	newIndex := 2
	var filePath string
	var resultFile *excelize.File

	readFile("./rule")
	checkConfig()

	if numFile <= 0 {
		fmt.Println("the Number of File has to over 1")
		return
	}

	for i:=0;i<numFile; {
		// fileName := os.Args[1]
		str := strings.Split(fileasName[i], ".")
		sheetName := str[0]

		filePath = "/Users/root1/Downloads/"
		filePos := fmt.Sprintf("%s%s",filePath, fileasName[i])

		// fmt.Println(filePos)
		resultSheetName := "Result"

		newIndex = 2
		goFlag = 0
		myTotalAsset = 0.0
		seed = periodSeed
		initSeed = firstSeed

		// fmt.Println(filePos)
		xlsx, err := excelize.OpenFile (filePos)
		if (err != nil) {
			fmt.Println(err)
			return
		}

		rows , err := xlsx.GetRows(sheetName)
		if (err != nil) {
			fmt.Println(err)
			return
		}

		sheet := xlsx.NewSheet(resultSheetName)

		for _, row := range rows {

			if row[0] == "Date" {
				continue;
			}

		// fmt.Println("###")
			if  goFlag == 0  {
				if startDate == row[0] {
					// fmt.Println(startDate, row[0])
					goFlag = 1
					// fmt.Println("GO LANG FLAG SET 1")
				}
				if  goFlag == 0  {
					// fmt.Println("Not matched date to", startDate)
					continue;
				}
			}

			// fmt.Println(endDate, row[0])
			if endDate == row[0] && goFlag == 1  {
				goFlag = 0
				// fmt.Println("GO LANG FLAG SET 0")
				continue;
			}

			t := convertDate(row[0])
			if t.Weekday() != 4 {
				continue
			}

			if ( skipWeek == 0) {
				skipWeek = 1
				continue
			}
			skipWeek = 0
			// fmt.Println("======", numFile)

			if (numFile == 1) {
			}else {
				//Open new file for result
				if (resultFile == nil) {
					resultFile = excelize.NewFile();
				}
			}

			if (xlsx != nil) {
				writeBasicInfo(xlsx, resultSheetName, newIndex, row);
				curVal, _ := strconv.ParseFloat(row[4], 64)
				calcMdd(curVal, newIndex)
				senario_2(xlsx, curVal, newIndex, resultSheetName)
			}

			newIndex++
    	}

		if (numFile == 1) {
			
			writeFinalInfo (xlsx, fileasName[i], resultSheetName, newIndex+5, 0);
			xlsx.SetActiveSheet(sheet)

			err = xlsx.SaveAs(filePos)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else {
			writeFinalInfo(resultFile, fileasName[i], "Sheet1", 5, i);

			newFilePos := fmt.Sprintf("%sResult.xlsx", filePath)
			// fmt.Println(newFilePos)

			// fmt.Println(newFilePos)
			err = resultFile.SaveAs(newFilePos)
			if err != nil {
				fmt.Println(err)
				return
			}
			xlsx.SetActiveSheet(sheet)

			err = xlsx.SaveAs(filePos)
			if err != nil {
				fmt.Println(err)
				return
			}

		}

		i++
	}
	fmt.Println("Done.")
	return
}