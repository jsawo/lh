package main

import (
	"flag"
	"fmt"

	"github.com/jsawo/loc/helpers"
	"github.com/jsawo/loc/takeout"
	"github.com/pterm/pterm"
)

func main() {
	inputFile := flag.String("i", "", "path to a takeout zip file")
	dateFilter := flag.String("f", "", "date filter (YYYY-MM-DD HH:MM:SS or left partials)")
	flag.Parse()

	visits := takeout.GetVisits(*inputFile, *dateFilter)

	for _, visit := range visits {
		termw := pterm.GetTerminalWidth()
		fmt.Printf("%v | %s | %s | %s | %v | %v | %v \n",
			visit.Order, visit.Date,
			helpers.Truncate(visit.Address, termw/4),
			helpers.Truncate(visit.PlaceName, termw/4),
			visit.VisitConfidence, visit.Latitude, visit.Longitude,
		)
	}
}
