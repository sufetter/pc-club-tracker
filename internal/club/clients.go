package club

import (
	"slices"

	"github.com/sufetter/pc-club-tracker/internal/event"
	"github.com/sufetter/pc-club-tracker/internal/shift"
)

type clientState struct {
	table     uint64
	startTime shift.Time
}

func (c *Club) NewClient(name string, currentTime shift.Time) event.Event {
	if currentTime.Compare(c.OpenTime) < 0 || currentTime.Compare(c.CloseTime) > 0 {
		return event.Error("NotOpenYet", currentTime)
	}
	if _, ok := c.clients[name]; ok {
		return event.Error("YouShallNotPass", currentTime)
	}
	c.clients[name] = &clientState{
		table: 0,
	}
	return event.Event{}
}

func (c *Club) SetClientTable(name string, table uint64, currentTime shift.Time) event.Event {
	if client, ok := c.clients[name]; ok {
		if c.Tables[table].isOccupied {
			return event.Error("PlaceIsBusy", currentTime)
		}
		c.LeaveTable(client, currentTime)

		client.table = table
		c.Tables[table].isOccupied = true
		client.startTime = currentTime
		c.TablesFree--
		return event.Event{}
	}
	return event.Error("ClientUnknown", currentTime)
}

func (c *Club) KickClient(name string, currentTime shift.Time) event.Event {
	if client, ok := c.clients[name]; ok {
		res := c.PrepareTable(client, currentTime)
		delete(c.clients, name)
		return res
	}
	return event.Error("ClientUnknown", currentTime)
}

func (c *Club) KickClientsSorted() (sortedClients []event.Event) {
	sortedClients = make([]event.Event, 0, len(c.clients))
	for name, client := range c.clients {
		sortedClients = append(sortedClients, event.New(c.CloseTime, event.ClientKick, name, 0, ""))
		c.PrepareTable(client, c.CloseTime)
	}
	slices.SortStableFunc(sortedClients, func(a, b event.Event) int {
		if a.UserName < b.UserName {
			return -1
		} else if a.UserName > b.UserName {
			return 1
		}
		return 0
	})
	return sortedClients
}
