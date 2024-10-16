package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/jessevdk/go-flags"
	"loconav.com/projects/internal/config"
	"loconav.com/projects/version"
)

var (
	yes = "Y"
	no  = "N"
)

type Options struct {
	Rows    int  `short:"r" long:"rows" description:"Flag to set value of rows in a board" default:"4"`
	Columns int  `short:"c" long:"columns" description:"Flag to set value of columns in a board" default:"4"`
	Version bool `short:"v" long:"version" description:"Flag to get the app version"`
}

func main() {
	var start string
	var opts Options
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatalf("Error while reading flag values. Err: %+v\n", err)
	}
	if opts.Version {
		version.DisplayVersion("Twenty48")
		return
	}
	if config.HaltIfEmpty(opts.Rows) {
		log.Fatalf("Unexpected Row value found while creating board. Expected: >4, Found: %+v\n", opts.Rows)
	}
	if config.HaltIfEmpty(opts.Columns) {
		log.Fatalf("Unexpected Column value found while creating board. Expected: >4, Found: %+v\n", opts.Columns)
	}
	if opts.Columns != opts.Rows {
		log.Fatal("Expect number of rows is equal to number of columns.")
	}
	log.Println("Use {W A S D} or {h j k l} or Arrow keys to move the board")
	log.Println("Where w - UP, a - Left, s - Down, d - Right")
	fmt.Printf("Do you want to start the game. Press Y/N: ")
	fmt.Scanln(&start)
	switch strings.ToUpper(start) {
	case yes:
		game := config.CreateBoard(opts.Rows, opts.Columns)
		game.AddElement()
		game.AddElement()
		for {
			if game.IsOver() {
				break
			}
			game.AddElement()
			game.Display()
			game.TakeInput()
		}
		log.Println("**** Game Over ****")
		maxTiles, totalTiles := game.CountScore()
		log.Println("Score: Max Tile Value: 	", maxTiles)
		log.Println("Score: Total Tiles Value: 	", totalTiles)
	case no:
		log.Fatalf("User don't want to play the game. Abort...")
	default:
		log.Fatalf("Unexpected user input received. Expected to be 'Y' or 'N'")
	}
}
