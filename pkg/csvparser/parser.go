package csvparser

import (
	"encoding/csv"
	"log"
	"mime/multipart"
	"strconv"
	"time"
)

const (
	weightTitle = "Weight"
	bpfTitle    = "BFP"
	timeTitle   = "Time"
	layout      = "2006-01-02"
)

type ParsedData struct {
	weight map[time.Time]float32
	bfp    map[time.Time]float32
}

func Parse(file multipart.File) (*ParsedData, error) {
	csvr := csv.NewReader(file)
	data, err := csvr.ReadAll()
	if err != nil {
		log.Println("Read All ", err)
		return nil, err
	}

	result := ParsedData{
		weight: make(map[time.Time]float32),
		bfp:    make(map[time.Time]float32),
	}

	for _, column := range data {
		if weightTitle == column[0] || bpfTitle == column[1] || timeTitle == column[2] {
			continue
		}
		recordedDate, err := time.Parse(layout, column[3])
		if err != nil {
			log.Println("Parse recorded date", err)
			return nil, err
		}

		weightf64tmp, err := strconv.ParseFloat(column[0], 32)
		if err != nil {
			log.Println("CSV Parse weight ", err)
			weightf64tmp = 0.0
		}
		result.weight[recordedDate] = float32(weightf64tmp)
		bfpf64tmp, err := strconv.ParseFloat(column[1], 32)
		if err != nil {
			log.Println("CSV Parse bfp ", err)
			bfpf64tmp = 0.0
		}
		result.bfp[recordedDate] = float32(bfpf64tmp)
	}

	return &result, nil
}
