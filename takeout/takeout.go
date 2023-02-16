package takeout

import (
	"archive/zip"
	"encoding/json"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/jsawo/loc/helpers"
)

type LocationData struct {
	Order           int
	Date            string
	Address         string
	PlaceName       string
	VisitConfidence int
	Latitude        float64
	Longitude       float64
}

var tableData []LocationData
var mu sync.Mutex
var wg = sync.WaitGroup{}

func GetVisits(history_data_file, dateFilter string) []LocationData {
	r, err := zip.OpenReader(history_data_file)
	helpers.PanicOnError(err)
	defer r.Close()

	for _, file := range r.File {
		if strings.Contains(file.Name, "Semantic Location History") {
			go processCompressedFile(file, file.Name, dateFilter, &wg)
		}
	}

	wg.Wait()

	sort.Slice(tableData, func(i, j int) bool {
		return tableData[i].Order < tableData[j].Order
	})

	return tableData
}

func processCompressedFile(file *zip.File, filename, dateFilter string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	sortNumber := getFileSortNumber(filename)

	rc, err := file.Open()
	helpers.PanicOnError(err)
	defer rc.Close()

	var payload LocationDataFile
	err = json.NewDecoder(rc).Decode(&payload)
	helpers.PanicOnError(err)

	for _, entry := range payload.TimelineObjects {
		// skip activitySegment objects
		if entry.PlaceVisit.CenterLatE7 == 0 {
			continue
		}

		// apply date filter
		if !strings.HasPrefix(entry.PlaceVisit.Duration.StartTimestamp.Format("2006-01-02 15:04:05"), dateFilter) {
			continue
		}

		sortNumber++

		mu.Lock()
		tableData = append(tableData, LocationData{
			Order:           sortNumber,
			Date:            entry.PlaceVisit.Duration.StartTimestamp.Format("2006-01-02 15:04:05"),
			Address:         entry.PlaceVisit.Location.Address,
			PlaceName:       entry.PlaceVisit.Location.Name,
			VisitConfidence: entry.PlaceVisit.VisitConfidence,
			Latitude:        float64(entry.PlaceVisit.Location.LatitudeE7) / 10000000,
			Longitude:       float64(entry.PlaceVisit.Location.LongitudeE7) / 10000000,
		})
		mu.Unlock()
	}
}

func getFileSortNumber(path string) int {
	last := strings.LastIndex(path, "/")
	filename := path[last+1 : len(path)-5]
	parts := strings.Split(filename, "_")
	year := parts[0]
	monthName := parts[1]
	month := ""

	switch monthName {
	case "JANUARY":
		month = "01"
	case "FEBRUARY":
		month = "02"
	case "MARCH":
		month = "03"
	case "APRIL":
		month = "04"
	case "MAY":
		month = "05"
	case "JUNE":
		month = "06"
	case "JULY":
		month = "07"
	case "AUGUST":
		month = "08"
	case "SEPTEMBER":
		month = "09"
	case "OCTOBER":
		month = "10"
	case "NOVEMBER":
		month = "11"
	case "DECEMBER":
		month = "12"
	}

	outInt, _ := strconv.Atoi(year + month + "0000")

	return outInt
}
