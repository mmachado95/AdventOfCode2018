package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"time"
)

// Record represents the input record data
type Record struct {
	date    time.Time
	command string
}

// RecordSort is used to sort the record values
type RecordSort []Record

func (recordSort RecordSort) Len() int           { return len(recordSort) }
func (recordSort RecordSort) Less(i, j int) bool { return recordSort[i].date.Before(recordSort[j].date) }
func (recordSort RecordSort) Swap(i, j int) {
	recordSort[i], recordSort[j] = recordSort[j], recordSort[i]
}

func getRecordValue(layout, line string) Record {
	re := regexp.MustCompile(`^\[(\d+\-\d+\-\d+\s+\d+\:\d+)\]\s+(.+)`)
	stringValues := re.FindAllStringSubmatch(line, -1)

	date, _ := time.Parse(layout, stringValues[0][1])
	sentence := stringValues[0][2]

	return Record{date, sentence}
}

func parseRecordValues(layout string) []Record {
	var records []Record

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		records = append(records, getRecordValue(layout, line))
	}

	return records
}

func getGuardsSleepTime(records []Record) map[int][]int {
	guards := map[int][]int{}
	var startSleep time.Time
	guardID := 0

	for _, record := range records {
		switch record.command {
		case "falls asleep":
			startSleep = record.date
		case "wakes up":
			for i := startSleep.Minute(); i < record.date.Minute(); i++ {
				guards[guardID][i]++
			}
		default:
			fmt.Sscanf(record.command, "Guard #%d begins shift", &guardID)
			if _, ok := guards[guardID]; !ok {
				guards[guardID] = make([]int, 60)
			}
		}
	}

	return guards
}

func main() {
	dateLayout := "2006-01-02 15:04"

	records := parseRecordValues(dateLayout)
	sort.Sort(RecordSort(records))

	guardsSleepTime := getGuardsSleepTime(records)

	maxTotalTimeSleeping := 0
	minuteSpentMostTimeSleeping := 0
	maxTimeSpentInMinute := 0
	guardID := -1

	for guard, minutes := range guardsSleepTime {
		totalTimeSleeping, minuteID := 0, 0

		for i, minute := range minutes {
			totalTimeSleeping += minute

			if minute > maxTimeSpentInMinute {
				maxTimeSpentInMinute, minuteID = minute, i
			}
		}
		if totalTimeSleeping > maxTotalTimeSleeping {
			maxTotalTimeSleeping, minuteSpentMostTimeSleeping, guardID = totalTimeSleeping, minuteID, guard
		}
	}

	fmt.Printf("Strategy 1: %v x %v = %v\n", guardID, minuteSpentMostTimeSleeping, guardID*minuteSpentMostTimeSleeping)

	maxTimeSpentInMinute = 0
	for guard, minutes := range guardsSleepTime {
		for i, minute := range minutes {
			if minute > maxTimeSpentInMinute {
				maxTimeSpentInMinute, guardID, minuteSpentMostTimeSleeping = minute, guard, i
			}
		}
	}

	fmt.Printf("Strategy 2: %vx%v = %v\n", guardID, minuteSpentMostTimeSleeping, guardID*minuteSpentMostTimeSleeping)
}
