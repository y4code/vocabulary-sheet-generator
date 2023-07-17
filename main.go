package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"os"
	"regexp"
)

type Vocabulary struct {
	Name           string
	PhoneticSymbol string
	Translation    template.HTML
}

func main() {
	csvFile, err := os.Open("sheet.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()
	var vocabularies []*Vocabulary
	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`美:[^\s]+`)
	for i, item := range record {
		if i == 0 {
			continue
		}
		vocabularies = append(vocabularies, &Vocabulary{
			Name:           item[1],
			PhoneticSymbol: re.ReplaceAllString(item[2], ""),
			Translation:    template.HTML(item[3]),
		})
	}
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	openFile, err := os.OpenFile("sheet.html", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer openFile.Close()
	err = tmpl.Execute(openFile, vocabularies)
	if err != nil {
		panic(err)
	}
	log.Println("success")
}
