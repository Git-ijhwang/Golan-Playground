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
var senarioType int

// var myTotalAsset float64 = 0.0
var myInvest float64 = 0.0
var myFinalAsset float64

var numFile int
var fileasName [5]string 
var startDate string
var endDate string

var stkMdd float64
var stkHighVal float64
var stkLowVal float64
var stkDateHigh string
var stkDateLow string
var stkDateDue string

var myMdd float64
var myHighVal float64
var myLowVal float64
var myDateHigh string
var myDateLow string
var myDateDue string

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

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

	str := strings.Split(cel, "-")
	date := "20"+str[2]+"-"+str[0]+"-"+str[1]
	t, _ := time.Parse(layoutISO, date)
	// fmt.Println(t.Day())
	// fmt.Println(int(t.Month()))
	// fmt.Println(t.Year())

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
		// writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myTotalAsset)
		writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myInvest)

		writeExcelString(xlsx, sheet, index, baseCol+2, "Final Asset Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+2, myFinalAsset)

		writeExcelString(xlsx, sheet, index, baseCol+3, "Yield")
		// writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myFinalAsset-myTotalAsset)/myTotalAsset*100))
		writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myFinalAsset-myInvest)/myInvest*100))

		writeExcelString(xlsx, sheet, index, baseCol+4, "MDD Duration")
		writeExcelString(xlsx, sheet, index+1, baseCol+4, stkDateDue)

		writeExcelString(xlsx, sheet, index, baseCol+5, "Stock MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+5, stkMdd)

		writeExcelString(xlsx, sheet, index, baseCol+6, "MDD Duration")
		writeExcelString(xlsx, sheet, index+1, baseCol+6, myDateDue)

		writeExcelString(xlsx, sheet, index, baseCol+7, "My Asset MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+7, myMdd)

	} else {

		// baseCol = baseCol+rune(new*8)
		index = index+(new*4)
		writeExcelString(xlsx, sheet, index-1, baseCol+1, tiker)
		writeExcelString(xlsx, sheet, index, baseCol+1, "Total Investing Value")
		// writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myTotalAsset)
		writeExcelFloat(xlsx, sheet, index+1, baseCol+1, myInvest)

		writeExcelString(xlsx, sheet, index, baseCol+2, "Final Asset Value")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+2, myFinalAsset)

		writeExcelString(xlsx, sheet, index, baseCol+3, "Yield")
		// writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myFinalAsset-myTotalAsset)/myTotalAsset*100))
		// writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myInvest-myTotalAsset)/myInvest*100))
		writeExcelString(xlsx, sheet, index+1, baseCol+3, fmt.Sprintf("%f%%",(myFinalAsset-myInvest)/myInvest*100))

		writeExcelString(xlsx, sheet, index, baseCol+4, "MDD Duration")
		writeExcelString(xlsx, sheet, index+1, baseCol+4, stkDateDue)

		writeExcelString(xlsx, sheet, index, baseCol+5, "Stock MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+5, stkMdd)

		writeExcelString(xlsx, sheet, index, baseCol+6, "MDD Duration")
		writeExcelString(xlsx, sheet, index+1, baseCol+6, myDateDue)

		writeExcelString(xlsx, sheet, index, baseCol+7, "My Asset MDD")
		writeExcelFloat(xlsx, sheet, index+1, baseCol+7, myMdd)
	}
}

func calcMdd (curVal float64, newIndex int, strDate string) {

	if (newIndex == 2) {
		stkMdd = 0
		stkHighVal = curVal
		stkLowVal = curVal
		stkDateHigh = ""
		stkDateLow = ""

		myMdd = 0
		myHighVal = curVal
		myLowVal = curVal
		myDateHigh = ""
		myDateLow = ""
	} else {
		if stkHighVal < curVal {
			stkHighVal = curVal
			stkLowVal = curVal
			stkDateHigh = strDate
			stkDateLow = strDate
		}

		if stkLowVal > curVal {
			stkLowVal = curVal
			stkDateLow = strDate
		} else {
			//Calc MDD
			if (stkHighVal == stkLowVal) {
			}else if (stkHighVal > stkLowVal) {
				tmp:=(stkHighVal - stkLowVal) / stkHighVal
				if (stkMdd == 0) {
					stkMdd = tmp
					stkDateDue = stkDateHigh+ " - " + stkDateLow
				} else {
					if (stkMdd < tmp ) {
						stkMdd = tmp
						stkDateDue = stkDateHigh+ " - " + stkDateLow
					}
				}
				stkHighVal = stkLowVal
			} 
		}

		if myHighVal < (curVal*float64(preStkCnt)) {
			myHighVal	= curVal*float64(preStkCnt)
			myLowVal	= curVal*float64(preStkCnt)
			myDateHigh	= strDate
			myDateLow	= strDate
		}

		if myLowVal > (curVal * float64(preStkCnt)) {
			myLowVal = curVal*float64(preStkCnt)
			myDateLow = strDate
			// fmt.Println(curVal, preStkCnt, myLowVal)
		} else {
			//Calc MDD
			if (myHighVal == myLowVal) {
			}else if (myHighVal > myLowVal) {
				// myMdd = (myHighVal - myLowVal) / myHighVal
				tmp:=(myHighVal - myLowVal) / myHighVal

				if (myMdd == 0.0) {
					myMdd = tmp
					myDateDue = stkDateHigh+ " - " + stkDateLow
				}else {
					if (myMdd < tmp) {
						myMdd = tmp
						myDateDue = stkDateHigh+ " - " + stkDateLow
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
	}

	if (newIndex == 2) {
		curSeed = initSeed
		myInvest = myInvest + curSeed
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

		// curSeed = seed
		curSeed = (curSeed-(float64(stkCnt)*curVal))+seed

	} else {
		/* pre Calculate value of my al stock */
		totPreVal := curVal*float64(preStkCnt)
		// myInvest = myInvest + curSeed

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

			curSeed = curSeed + seed

		} else {
			/* Buy */
			writeExcelString(xlsx, sheet, newIndex, 'A', "BUY")

			writeExcelFloat(xlsx, sheet, newIndex, 'E', curSeed)

			stkCnt = int(curSeed/curVal)
			writeExcelInt(xlsx, sheet, newIndex, 'F', stkCnt)

			preStkCnt = int(preStkCnt) + stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			if curSeed > seed {
				myInvest = myInvest + seed
			} else {
				// myInvest = myInvest + (float64(stkCnt)*curVal)
			}
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase) 

			curSeed = curSeed - float64(stkCnt)*curVal
			fmt.Println(curVal, " --- ", curSeed)
			curSeed = curSeed+seed
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
	}

	if (newIndex == 2) {
		curSeed = initSeed
		myInvest = myInvest + curSeed
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

		curSeed = (curSeed-(float64(stkCnt)*curVal))+seed

	} else {
		/* pre Calculate value of my al stock */
		totPreVal := curVal*float64(preStkCnt)
		myInvest = myInvest + curSeed

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
			} else {
				stkCnt = int((totPreVal - highBase)/curVal)
			}
			stkCnt = 0; //Temporary
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

			writeExcelInt(xlsx, sheet, newIndex, 'F', stkCnt)

			preStkCnt = int(preStkCnt) + stkCnt
			writeExcelInt(xlsx, sheet, newIndex, 'G', preStkCnt)

			accumVal = float64(preStkCnt)*curVal
			myFinalAsset = accumVal
			myInvest = myInvest + (float64(stkCnt)*curVal)
			writeExcelFloat(xlsx, sheet, newIndex, 'H', accumVal)

			lowBase = (accumVal*base)*low
			writeExcelFloat(xlsx, sheet, newIndex, 'I', lowBase)

			highBase = (accumVal*base)*high
			writeExcelFloat(xlsx, sheet, newIndex, 'J', highBase) 

			curSeed = curSeed - float64(stkCnt)*curVal
			curSeed = curSeed + seed
		}
	}
}

// func senario_2
var triger int
func senario_3 (xlsx *excelize.File, curVal float64, newIndex int, sheet string, time.Time t) {
	
	prevMonth := 0
	curMonth := int(t.Month())
	preVal := 0.0

	if prevMonth== 0 {
		prevMonth = curMonth
	}

	if prevMonth != curMonth {
		triger = 1
		prevMonth = curMonth
	}

	if prevMonth == curMonth {
		triger = 0
	}

	if triger == 1 {
	} else {
	}

	if (preVal == 0) {
		preVal = curVal
		return
	}
	
	if (preVal > 0) {
		// rate := curVal - preVal
	}

}

func checkConfig() {
	i:=0
	fmt.Println("===============================");
	fmt.Println("Init Seed : ", initSeed)
	fmt.Println("Monthly Seed : ", seed)
	fmt.Println("Senario Type : ", senarioType)
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

	numOfFile:=0
	for scanner.Scan() {
		line := scanner.Text()

		if (line == "") {
			continue;
		}

		if line[0] == '#'{
			continue
		}

		ruleData := strings.Split(line, ":")
		if ( ruleData[0] == "SENARIO_TYPE") {
			val, _ := (strconv.Atoi(ruleData[1]))
			senarioType = val
			continue;
		}

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

	for i := 0; i<numFile; {
		str := strings.Split(fileasName[i], ".")
		sheetName := str[0]

		filePath = "/Users/root1/Downloads/"
		filePos := fmt.Sprintf("%s%s",filePath, fileasName[i])
		fmt.Println(filePos)

		resultSheetName := "Result"

		newIndex = 2
		goFlag = 0
		seed = periodSeed
		initSeed = firstSeed

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

			if  goFlag == 0  {
				if startDate == row[0] {
					goFlag = 1
				}
				if  goFlag == 0  {
					continue;
				}
			}

			if endDate == row[0] && goFlag == 1  {
				goFlag = 0
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
				strDate := t.Format(layoutISO)

				calcMdd(curVal, newIndex, strDate)
				if senarioType == 1 {
					senario_1(xlsx, curVal, newIndex, resultSheetName)
				} else {
					senario_2(xlsx, curVal, newIndex, resultSheetName)
				}
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

			newFilePos := fmt.Sprintf("%sResult_%d.xlsx", filePath, senarioType)

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