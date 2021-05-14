package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/cdipaolo/sentiment"
)

func main() {
	// open file
	csvFile, err := os.Open("./data/imdb-dataset.csv")
	if err != nil {
		log.Fatalln("Error:  could not open file", err)
	}
	defer csvFile.Close()
	parseCSV := csv.NewReader(csvFile) // parse dataset

	// createa sentiment model
	model, err := sentiment.Restore()

	if err != nil {
		panic(err)
	}
	var analysis *sentiment.Analysis
	// use a while loop to go throught the data
	for {
		// read the data from csv
		record, err := parseCSV.Read()
		if err == io.EOF {
			// leave the loop
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		analysis = model.SentimentAnalysis(record[0], sentiment.English)
		var sentiment string
		if analysis.Score == 1 {
			sentiment = "positive"
		} else {
			sentiment = "negetive"
		}
		fmt.Printf("Review: %s \n and Sentiment:%s\n", record[0], sentiment)

	}

}
