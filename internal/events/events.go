package events

import (
	"errors"

	"github.com/sufetter/pc-club-tracker/internal/models"
)

func ProcessEvent(event models.Event, tables []models.Table) error {
	switch event.ID {
	case 1: // Клиент пришел
	case 2: // Клиент сел за стол
	case 3: // Клиент ожидает
	case 4: // Клиент ушел
	default:
		return errors.New("unknown event ID")
	}
	return nil
}
