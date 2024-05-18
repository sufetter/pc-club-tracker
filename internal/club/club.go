package club

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/sufetter/pc-club-tracker/internal/shift"
	"github.com/sufetter/pc-club-tracker/pkg/queue"
)

type Club struct {
	TablesFree          uint64
	OpenTime, CloseTime shift.Time
	HourCost            uint64
	Tables              []Table
	clients             map[string]*clientState
	clientsQueue        queue.Queue[string]
}

func Config(scn *bufio.Scanner) (*Club, error) {
	scn.Scan()
	line := scn.Text()
	tablesNum, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", line, err)
	}

	scn.Scan()
	line = scn.Text()
	workingHoursStr := strings.Fields(line)
	if len(workingHoursStr) != 2 {
		return nil, fmt.Errorf("%s: %w", line, err)
	}
	openTimeStr, closeTimeStr := workingHoursStr[0], workingHoursStr[1]
	openTime, err := shift.ParseTime(openTimeStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", line, err)
	}
	closeTime, err := shift.ParseTime(closeTimeStr)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", line, err)
	}
	if openTime.Compare(closeTime) > 0 || openTime.Compare(closeTime) == 0 || closeTime.Hours > 23 {
		return nil, fmt.Errorf("%s: %w", line, err)
	}

	scn.Scan()
	line = scn.Text()
	hourCost, err := strconv.ParseUint(line, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", line, err)
	}

	club := New(tablesNum, openTime.Hours, openTime.Minutes, closeTime.Hours, closeTime.Minutes, hourCost)

	return club, nil
}
func countProfit(startTime, endTime shift.Time, costPerHour uint64) uint64 {
	t := endTime.Sub(startTime)
	if t.Minutes > 0 {
		return (t.Hours + uint64(1)) * costPerHour
	}
	return t.Hours * costPerHour
}

func New(tablesFree, openHours, openMinutes, closeHours, closeMinutes, hourCost uint64) *Club {
	// Конвертация выбрана в угоду типу полей uint64 в Club
	return &Club{
		TablesFree: tablesFree,
		OpenTime: shift.Time{
			Hours:   openHours,
			Minutes: openMinutes,
		},
		CloseTime: shift.Time{
			Hours:   closeHours,
			Minutes: closeMinutes,
		},
		HourCost: hourCost,

		Tables:       make([]Table, tablesFree+1),
		clients:      make(map[string]*clientState),
		clientsQueue: queue.New[string](),
	}
}
