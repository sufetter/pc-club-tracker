package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/sufetter/pc-club-tracker/internal/club"
	"github.com/sufetter/pc-club-tracker/internal/event"
	"github.com/sufetter/pc-club-tracker/internal/shift"
	"github.com/sufetter/pc-club-tracker/pkg/queue"
)

func ParseTXT(scanner *bufio.Scanner, pcClub *club.Club) (queue.Queue[string], error) {
	checkName, err := regexp.Compile(event.NamePattern)
	if err != nil {
		return queue.Queue[string]{}, fmt.Errorf("error compiling regex: %w", err)
	}

	outputQueue := queue.New[string]()
	prevTime := shift.Time{Hours: 0, Minutes: 0}
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		lineParts := strings.Split(line, " ")
		if len(lineParts) < 3 {
			return queue.Queue[string]{}, fmt.Errorf("wrong line format: %w\n%s", err, line)
		}

		var time shift.Time
		time, err = shift.ParseTime(lineParts[0])
		if err != nil || time.Compare(prevTime) < 0 {
			return queue.Queue[string]{}, fmt.Errorf("wrong time format: %w\n%s", err, line)
		}
		prevTime = time
		if time.Compare(pcClub.CloseTime) > 0 {
			return queue.Queue[string]{}, fmt.Errorf("wrong time format: %w\n%s", err, line)
		}

		var eventCode uint64
		eventCode, err = strconv.ParseUint(lineParts[1], 10, 64)
		if err != nil {
			return queue.Queue[string]{}, fmt.Errorf("wrong event code: %w\n%s", err, line)
		}

		if !checkName.MatchString(lineParts[2]) {
			return queue.Queue[string]{}, fmt.Errorf("wrong client name: %w\n%s", err, line)
		}

		switch eventCode {
		case event.ClientCame:
			outputQueue.Push(line)
			answer := pcClub.NewClient(lineParts[2], time)
			if answer.Code != 0 {
				outputQueue.Push(answer.String())
			}
		case event.ClientSit:
			if len(lineParts) < 4 {
				return queue.Queue[string]{}, fmt.Errorf("wrong line format\n%s", line)
			}
			var table uint64
			table, err = strconv.ParseUint(lineParts[3], 10, 64)
			if err != nil {
				return queue.Queue[string]{}, fmt.Errorf("wrong table number: %w\n%s", err, line)
			}
			outputQueue.Push(line)
			answer := pcClub.SetClientTable(lineParts[2], table, time)

			if answer.Code != 0 {
				outputQueue.Push(answer.String())
			}
		case event.ClientWaiting:
			outputQueue.Push(line)
			answer := pcClub.WaitForTable(lineParts[2], time)
			if answer.Code != 0 {
				outputQueue.Push(answer.String())
			}
		case event.ClientLeft:
			outputQueue.Push(line)
			answer := pcClub.KickClient(lineParts[2], time)
			if answer.Code != 0 {
				outputQueue.Push(answer.String())
			}
		default:
			return queue.Queue[string]{}, fmt.Errorf("wrong event code: %w\n%s", err, line)
		}
	}
	sortedLeaveEvents := pcClub.KickClientsSorted()
	for _, leaveEvent := range sortedLeaveEvents {
		outputQueue.Push(leaveEvent.String())
	}
	pcClub.Tables = append(pcClub.Tables[:0], pcClub.Tables[1:]...)
	return outputQueue, nil
}
