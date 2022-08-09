package utils

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/RyanMarshall5765/bitlyproject/cfg"
)

func AverageMinutes(constructedTimes []string) int {
	var totalMinutes float64
	// Not sure how to set a date as config varible. Can't set as const in cfg.
	startDate := time.Date(0, 0, 1, 8, 0, 0, 0, time.UTC)
	for _, constructedTime := range constructedTimes {
		deconstructedTime, _ := DeconstructTimeString(constructedTime)
		invalidTime := deconstructedTime.ReturnEndDate().Before(startDate)
		if invalidTime {
			log.Fatalf("Sorry, there was an input found before the start of the race at %s. Please start over.", cfg.RaceStartString)
		}
		totalMinutes += FindDifferenceInMinutes(startDate, deconstructedTime.ReturnEndDate())
	}
	return int(math.Round(totalMinutes / float64(len(constructedTimes))))
}

type deconstructedTime struct {
	hour   int
	minute int
	ampm   string
	day    int
}

func (d *deconstructedTime) ReturnEndDate() (endDate time.Time) {
	endDate = time.Date(0, 0, d.day, d.hour, d.minute, 0, 0, time.UTC)
	return
}

func DeconstructTimeString(constructedTimeString string) (deconstructedTime deconstructedTime, err error) {
	endDatePattern := regexp.MustCompile(cfg.TimeStampPattern)
	matches := endDatePattern.FindStringSubmatch(constructedTimeString)

	if matches == nil {
		log.Fatalf("invalid format! expected format: %s", cfg.TimeStampFormat)
	}

	hourIndex := endDatePattern.SubexpIndex("hh")
	minuteIndex := endDatePattern.SubexpIndex("mm")
	ampmIndex := endDatePattern.SubexpIndex("x")
	dayIndex := endDatePattern.SubexpIndex("n")

	deconstructedTime.day, _ = strconv.Atoi(matches[dayIndex])
	deconstructedTime.hour, _ = strconv.Atoi(matches[hourIndex])
	deconstructedTime.minute, _ = strconv.Atoi(matches[minuteIndex])
	deconstructedTime.ampm = matches[ampmIndex]

	if deconstructedTime.hour < 12 && deconstructedTime.ampm == "PM" {
		deconstructedTime.hour += 12
	} else if deconstructedTime.hour >= 12 && deconstructedTime.ampm == "AM" {
		deconstructedTime.hour -= 12
	}

	return
}

func FindDifferenceInMinutes(firstDate, secondDate time.Time) float64 {
	return secondDate.Sub(firstDate).Minutes()
}

func ReadFileByLine(file string, pattern string) (lineContent []string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	rx, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	lineNumber := 1
	for scanner.Scan() {
		match := rx.Match([]byte(scanner.Text()))
		if !match {
			log.Fatalf("Did not match format %s on line %d", cfg.TimeStampFormat, lineNumber)
		}
		lineContent = append(lineContent, scanner.Text())
		lineNumber += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
