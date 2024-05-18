package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sufetter/pc-club-tracker/internal/club"
	"github.com/sufetter/pc-club-tracker/internal/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ", os.Args[0], "<filename>")
		return
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file: %w", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pcClub *club.Club
	pcClub, err = club.Config(scanner)
	if err != nil {
		fmt.Println("error parsing file: %w", err)
		return
	}
	events, err := parser.ParseTXT(scanner, pcClub)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(pcClub.OpenTime)
	for !events.IsEmpty() {
		fmt.Println(events.Pop())
	}
	fmt.Println(pcClub.CloseTime)

	for num, table := range pcClub.Tables {
		fmt.Println(num+1, table)
	}
}
