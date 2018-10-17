package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gocarina/gocsv"
)

/*struct to unmarshall the csv values*/
type CsvLine_csv struct {
	Key   string `csv:"key,omitempty"`
	Value string `csv:"value,omitempty"`
}

func main() {
	csvFile, err := os.OpenFile("Corpus.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println("error", err)
	}
	var clients []*CsvLine_csv
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ','
		return r
	})
	if err := gocsv.UnmarshalFile(csvFile, &clients); err != nil { // Load clients from file
		fmt.Println("error", err)
		return
	}
	goibibo_map := make(map[string]string)
	for _, client := range clients {
		goibibo_map[client.Value] = client.Key
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.EqualFold(goibibo_map[strings.TrimPrefix(r.URL.Path, "/")], "") {
			fmt.Println("search not found", strings.TrimPrefix(r.URL.Path, "/"))
			w.WriteHeader(http.StatusNotFound)
		} else {
			fmt.Println("search found", strings.TrimPrefix(r.URL.Path, "/"))
			var response_csv CsvLine_csv
			response_csv.Key = goibibo_map[strings.TrimPrefix(r.URL.Path, "/")]
			response_csv.Value = strings.TrimPrefix(r.URL.Path, "/")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			response, _ := json.Marshal(response_csv)
			w.Write(response)
		}
	})
	fmt.Println("web server listening @ port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
