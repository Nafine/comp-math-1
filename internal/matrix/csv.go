package matrix

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSV(filePath string) ([][]float64, error) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
		return nil, err
	}

	fmt.Println(records)

	return nil, nil
}
