package main

import (
	"strconv"
	"time"
	"regexp"
)

const (
	DEFAULT_HOUR = 19
	DEFAULT_MINUTES = 30
)

func convertDateStringToInnerFormat(date string) string {
	splittedDate := regexp.MustCompile("[-/.]").Split(date, 3)
	month, _ := strconv.Atoi(splittedDate[0])
	day, _ := strconv.Atoi(splittedDate[1])
	//Problem of excel parser library
	year, _ := strconv.Atoi("20" + splittedDate[2])
	newDate := time.Date(year, time.Month(month), day, DEFAULT_HOUR, DEFAULT_MINUTES, 0, 0, time.Now().Location())
	return newDate.Format(time.RFC3339)
}

func isFutureDate(date string) bool {
	t := time.Now()
	parsedTime, _ := time.Parse(time.RFC3339, date)
	remainingTime := t.Sub(parsedTime)
	return remainingTime < 1
}
