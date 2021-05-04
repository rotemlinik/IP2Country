package db

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"
)

type CsvImpl struct {
	ip2LocationMap map[string]Location
}

func NewDb() Db {
	dbConfig := newDbConfig()
	csvImpl := CsvImpl{}
	csvImpl.mapCsv(getRecordsFromCsv(dbConfig.dbPath))

	return &csvImpl
}

func (dbCsvImpl *CsvImpl) GetLocation(ip string) Location {
	return dbCsvImpl.ip2LocationMap[ip]
}

func (dbCsvImpl *CsvImpl) mapCsv(records [][]string) {
	ip2LocationMap := make(map[string]Location)

	for i := 1; i < len(records); i++ {
		record := records[i]
		ip2LocationMap[record[0]] = Location{Country: record[1], City: record[2]}
	}

	dbCsvImpl.ip2LocationMap = ip2LocationMap
}

func getRecordsFromCsv(path string) [][]string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(filepath.Join(pwd, "/db", path))
	if err != nil {
		log.Fatal(err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return records
}