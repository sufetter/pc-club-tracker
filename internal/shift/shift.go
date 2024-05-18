package shift

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hours, Minutes uint64
}

func (t Time) Compare(cmp Time) int {
	if t.Hours < cmp.Hours {
		return -1
	}
	if t.Hours > cmp.Hours {
		return 1
	}
	if t.Minutes < cmp.Minutes {
		return -1
	}
	if t.Minutes > cmp.Minutes {
		return 1
	}
	return 0
}

func (t Time) Add(add Time) Time {
	minutes := t.Minutes + add.Minutes
	if minutes > 59 {
		t.Hours++
		minutes -= 60
	}
	return Time{
		Hours:   t.Hours + add.Hours,
		Minutes: minutes,
	}
}

func (t Time) Sub(sub Time) Time {
	if t.Minutes < sub.Minutes {
		return Time{
			Hours:   t.Hours - 1 - sub.Hours,
			Minutes: t.Minutes + 60 - sub.Minutes,
		}
	}
	return Time{
		Hours:   t.Hours - sub.Hours,
		Minutes: t.Minutes - sub.Minutes,
	}
}

func (t Time) String() string {
	if t.Minutes < 10 {
		return fmt.Sprintf("%02d:0%d", t.Hours, t.Minutes)
	}
	return fmt.Sprintf("%02d:%02d", t.Hours, t.Minutes)
}

func ParseTime(timeStr string) (time Time, err error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return time, fmt.Errorf("invalid time format")
	}

	if len(parts[0]) != 2 || len(parts[1]) != 2 {
		return time, fmt.Errorf("%s: %w", timeStr, err)
	}

	var hours, minutes uint64
	hours, err = strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return time, fmt.Errorf("invalid hours: %v", err)
	}

	minutes, err = strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return time, fmt.Errorf("invalid minutes: %v", err)
	}
	time.Hours = hours
	time.Minutes = minutes
	return
}
