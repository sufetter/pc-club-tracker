package club

import (
	"bufio"
	"strings"
	"testing"

	"github.com/sufetter/pc-club-tracker/internal/shift"
)

func TestClub_Config(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		cfg, err := Config(bufio.NewScanner(strings.NewReader("3\n09:00 19:00\n10\n")))
		if cfg.OpenTime.Hours != 9 {
			t.Error("Expected OpenTime.Hours to be 9, got", cfg.OpenTime.Hours)
		}
		if err != nil {
			t.Error("Expected nil error", err)
		}
	})
	t.Run("Negative", func(t *testing.T) {
		_, err := Config(bufio.NewScanner(strings.NewReader("3\n01:00 11:00 -1 Gigachad\n11\n")))
		if err == nil {
			t.Error("Expected line format error")
		}
	})
	t.Run("Wrong time format", func(t *testing.T) {
		_, err := Config(bufio.NewScanner(strings.NewReader("3\n77:00 09:00\n10\n")))
		if err == nil {
			t.Error("Must return line format error")
		}
	})
	t.Run("Wrong time range", func(t *testing.T) {
		_, err := Config(bufio.NewScanner(strings.NewReader("3\n09:00 25:00\n10\n")))
		if err == nil {
			t.Error("Must return line format error")
		}
	})
}

func TestClub_NewClient(t *testing.T) {
	t.Run("Positive", func(t *testing.T) {
		club := New(3, 9, 0, 19, 0, 10)
		client := club.NewClient("Sufetter", shift.Time{Hours: 9, Minutes: 0})
		if client.Code != 0 {
			t.Error("Expected 0 code, got", client.Code)
		}
		if client.String() != "" {
			t.Error("Expected empty string", client.String())
		}
	})
}
