package api

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"time"
)

type ColStat struct {
	bankEntryDate   map[string]int
	bankEntryText   map[string]int
	bankEntryAmount map[string]int
	accountName     map[string]int
	accountNumber   map[string]int
	accountTypeName map[string]int
}

func (c *ColStat) init() {
	c.bankEntryAmount = make(map[string]int)
	c.bankEntryDate = make(map[string]int)
	c.bankEntryText = make(map[string]int)
	c.accountName = make(map[string]int)
	c.accountNumber = make(map[string]int)
	c.accountTypeName = make(map[string]int)
}

type StatCalculator struct {
	Document       []byte
	ColumnsExclude []string
	ColumnsInclude []string
}

func (s *StatCalculator) calculateStatistics() (string, error) {

	reader := csv.NewReader(bytes.NewBuffer(s.Document))
	_, err := reader.Read()
	if err != nil {
		log.Fatalf("Error in reading csv %v", err)
		return "", err
	}

	idMap := make(map[string]ColStat)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error in reading csv row by row: %v", err)
			return "", err
		}
		companyId := record[1]
		colStat, ok := idMap[companyId]
		if ok == false {
			colStat = ColStat{}
			colStat.init()
			idMap[companyId] = colStat
		}
		colStat.bankEntryDate[record[2]] += 1
		colStat.bankEntryText[record[3]] += 1
		colStat.bankEntryAmount[record[4]] += 1
		colStat.accountName[record[5]] += 1
		colStat.accountNumber[record[6]] += 1
		colStat.accountTypeName[record[7]] += 1
	}

	output := s.printStatistics(idMap)
	return output, nil
}

func (s *StatCalculator) printStatistics(idMap map[string]ColStat) string {
	var outString string
	outString = fmt.Sprintf(",CompanyId,ColumnName,ColumnValue,Count\n")
	count := 0
	for companyId, colStat := range idMap {
		for key, val := range colStat.bankEntryDate {
			tempString := fmt.Sprintf("%d,%s,BankEntryDate,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}
		for key, val := range colStat.bankEntryText {
			tempString := fmt.Sprintf("%d,%s,BankEntryText,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}

		for key, val := range colStat.bankEntryAmount {
			tempString := fmt.Sprintf("%d,%s,BankEntryAmount,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}

		for key, val := range colStat.accountName {
			tempString := fmt.Sprintf("%d,%s,AccountName,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}

		for key, val := range colStat.accountNumber {
			tempString := fmt.Sprintf("%d,%s,AccountNumber,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}

		for key, val := range colStat.accountTypeName {
			tempString := fmt.Sprintf("%d,%s,AccountTypeName,%s,%d\n", count, companyId, key, val)
			outString = outString + tempString
			count++
		}
	}
	t := time.Now()
	fileName := "result-" + t.Format("0102200615040500000") + ".csv"
	ioutil.WriteFile("result/"+fileName, []byte(outString), 0644)
	return "result/" + fileName
}
