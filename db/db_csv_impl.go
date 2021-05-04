package db

import (
	"encoding/csv"
	"log"
	"os"
)

type CsvImpl struct {
	ip2LocationMap map[string]Location
}

func NewDb() Db {
	dbConfig := newDbConfig()

	f, err := os.Open(dbConfig.dbPath)
	if err != nil {
		log.Fatal("Unable to read input file " + dbConfig.dbPath, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal("Unable to close input file " + dbConfig.dbPath, err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + dbConfig.dbPath, err)
	}

	csvImpl := CsvImpl{}
	csvImpl.ip2LocationMap = mapCsv(records)

	return &csvImpl
}

func (dbCsvImpl *CsvImpl) GetLocation(ip string) Location {
	return dbCsvImpl.ip2LocationMap[ip]
}

func mapCsv(records [][]string) map[string]Location {
	returnMap := make(map[string]Location)

	for i := 1; i < len(records); i++ {
		record := records[i]
		returnMap[record[0]] = Location{Country: record[1], City: record[2]}
	}

	return returnMap
}