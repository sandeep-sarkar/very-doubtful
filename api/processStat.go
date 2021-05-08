package api

import "fmt"

type StatCalculator struct {
	Document       []byte
	ColumnsExclude []string
	ColumnsInclude []string
	PrimaryColumn  string
}

func (s *StatCalculator) printDocument() {
	fmt.Println("Document shared : ")
	fmt.Println(string(s.Document))
}
