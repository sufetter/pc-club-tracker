package event

import (
	"strconv"

	"github.com/sufetter/pc-club-tracker/internal/shift"
)

const (
	NamePattern = "^[a-z0-9_-]+$"
	ClientCame  = iota
	ClientSit
	ClientWaiting
	ClientLeft
)

const (
	ClientKick = iota + 11
	ClientSitEmpty
	ErrUnknown
)

type Event struct {
	Time     shift.Time
	Code     int
	UserName string
	Table    int
	Err      string
}

func New(time shift.Time, code int, userName string, table int, err string) Event {
	return Event{
		Time:     time,
		Code:     code,
		UserName: userName,
		Table:    table,
		Err:      err,
	}
}

func Error(err string, time shift.Time) Event {
	return Event{Code: ErrUnknown, Err: err, Time: time}
}

func (r Event) String() string {
	switch r.Code {
	case ClientKick:
		return r.Time.String() + " " + strconv.Itoa(r.Code) + " " + r.UserName
	case ClientSitEmpty:
		return r.Time.String() + " " + strconv.Itoa(r.Code) + " " + r.UserName + " " + strconv.Itoa(r.Table)
	case ErrUnknown:
		return r.Time.String() + " " + strconv.Itoa(r.Code) + " " + r.Err
	default:
		return ""
	}
}
