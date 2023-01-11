package main

import (
	"aoc/input"
	"bufio"
	"bytes"
	"fmt"
	"github.com/spf13/cast"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	//"strings"
	//"time"
)

var lineRegex = regexp.MustCompile(`\[([^\]]+)\] (.+)$`)
var guardShiftRegex = regexp.MustCompile(`Guard #(\d+) begins shift`)

type actionType int

const (
	actionTypeStartShift actionType = iota
	actionTypeFallAsleep
	actionTypeWakeUp
)

type Solution struct{}

func (s *Solution) Date() (int, int) {
	return 2018, 4
}

func (s *Solution) Execute(input []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(input))

	var logs []struct {
		date   time.Time
		action string
	}

	for scanner.Scan() {
		line := scanner.Text()

		result := lineRegex.FindStringSubmatch(line)

		t, err := time.Parse("2006-01-02 15:04", result[1])
		if err != nil {
			return err
		}

		logs = append(logs, struct {
			date   time.Time
			action string
		}{
			date:   t,
			action: result[2],
		})
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].date.Before(logs[j].date)
	})

	type guard struct {
		id         int
		totalSleep time.Duration
		asleepMinute map[int]int
		mostAsleepAt int
		mostAsleepCount int
	}

	var guards = make(map[int]*guard)
	var prevGuard *guard
	var sleepStart time.Time
	var mostID = -1
	for _, log := range logs {
		if guardShiftRegex.MatchString(log.action) {
			parts := guardShiftRegex.FindStringSubmatch(log.action)
			id := cast.ToInt(parts[1])

			if g, ok := guards[id]; ok {
				prevGuard = g
			} else {
				newGuard := &guard{
					id: id,
					mostAsleepAt: -1,
					mostAsleepCount: -1,
					asleepMinute: make(map[int]int),
				}
				guards[id] = newGuard
				prevGuard = newGuard
			}

		} else if strings.EqualFold(log.action, "falls asleep") {
			sleepStart = log.date
		} else if strings.EqualFold(log.action, "wakes up") {
			prevGuard.totalSleep += log.date.Sub(sleepStart)
			if mostID == -1 || (prevGuard.totalSleep > guards[mostID].totalSleep) {
				mostID = prevGuard.id
			}

			for i := sleepStart.Minute(); i < log.date.Minute(); i++ {
				prevGuard.asleepMinute[i]++

				if prevGuard.asleepMinute[i] > prevGuard.mostAsleepCount {
					prevGuard.mostAsleepAt = i
					prevGuard.mostAsleepCount = prevGuard.asleepMinute[i]
				}
			}
		}
	}

	fmt.Println("Slept most Guard ID:", mostID)
	fmt.Println("Most asleep at:", guards[mostID].mostAsleepAt)
	fmt.Println("Answer:", mostID * guards[mostID].mostAsleepAt)

	var asleepMinuteCount = -1
	var asleepMinute = -1
	var mostAsleepGuardID int
	for _, guard := range guards {
		if guard.mostAsleepCount > asleepMinuteCount {
			asleepMinuteCount = guard.mostAsleepCount
			asleepMinute = guard.mostAsleepAt
			mostAsleepGuardID = guard.id
		}
	}

	fmt.Println("Guard ID:", mostAsleepGuardID)
	fmt.Println("Most asleep at:", asleepMinute)
	fmt.Println("Answer:", mostAsleepGuardID * asleepMinute)

	return nil
}

//var exampleIn = []byte(`[1518-11-01 00:00] Guard #10 begins shift
//[1518-11-01 00:05] falls asleep
//[1518-11-01 00:25] wakes up
//[1518-11-01 00:30] falls asleep
//[1518-11-01 00:55] wakes up
//[1518-11-01 23:58] Guard #99 begins shift
//[1518-11-02 00:40] falls asleep
//[1518-11-02 00:50] wakes up
//[1518-11-03 00:05] Guard #10 begins shift
//[1518-11-03 00:24] falls asleep
//[1518-11-03 00:29] wakes up
//[1518-11-04 00:02] Guard #99 begins shift
//[1518-11-04 00:36] falls asleep
//[1518-11-04 00:46] wakes up
//[1518-11-05 00:03] Guard #99 begins shift
//[1518-11-05 00:45] falls asleep
//[1518-11-05 00:55] wakes up`)

func main() {
	s := Solution{}
	in, err := input.GetInput(s.Date())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = s.Execute(in)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
