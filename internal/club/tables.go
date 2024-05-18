package club

import (
	"strconv"

	"github.com/sufetter/pc-club-tracker/internal/event"
	"github.com/sufetter/pc-club-tracker/internal/shift"
)

type Table struct {
	isOccupied bool
	WorkedTime shift.Time
	profit     uint64
}

func (t Table) String() string {
	return strconv.FormatUint(t.profit, 10) + " " + t.WorkedTime.String()
}

func (c *Club) PrepareTable(client *clientState, currentTime shift.Time) event.Event {
	if client.table != 0 {
		curTable := client.table
		c.LeaveTable(client, currentTime)

		for c.clientsQueue.Len() > 0 {
			newClientName := c.clientsQueue.Pop()
			if newClient, ok := c.clients[newClientName]; ok {
				newClient.table = curTable
				newClient.startTime = currentTime
				c.Tables[curTable].isOccupied = true
				c.TablesFree--
				return event.Event{Code: event.ClientSitEmpty, Table: int(newClient.table), Time: currentTime, UserName: newClientName}
			}
		}
	}
	return event.Event{}
}

func (c *Club) LeaveTable(client *clientState, currentTime shift.Time) {
	if client.table == 0 {
		return
	}
	c.Tables[client.table].profit += countProfit(client.startTime, currentTime, c.HourCost)
	c.Tables[client.table].WorkedTime = c.Tables[client.table].WorkedTime.Add(currentTime.Sub(client.startTime))
	c.Tables[client.table].isOccupied = false
	c.TablesFree++
	client.table = 0
	client.startTime = shift.Time{}
}

func (c *Club) WaitForTable(name string, currentTime shift.Time) event.Event {
	if client, ok := c.clients[name]; ok {
		if c.TablesFree > 0 {
			return event.Error("ICanWaitNoLonger!", currentTime)
		}
		if c.clientsQueue.Len() > len(c.Tables) {
			c.KickClient(name, currentTime)
			return event.New(currentTime, event.ClientKick, name, 0, "")
		}
		res := c.PrepareTable(client, currentTime)
		c.clientsQueue.Push(name)
		return res
	}
	return event.Error("ClientUnknown", currentTime)
}
