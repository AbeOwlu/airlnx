package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Improvement:
// Use cobra pkg commands
// cast curent parsed action parameter to uppercase - for usability
// var CommandLine []string = flag.Args()
var Usage = func() {
	fmt.Fprintf(os.Stderr, "FAIL.\nExample Usage of %s:\n", os.Args[0])
	PrintDefault()
}

func main() {
	var CommandLine []string = os.Args

	// Initialize logger
	logger, err := os.OpenFile("airlnx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		os.Exit(2)
	}
	defer logger.Close()
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)
	log.SetOutput(logger)

	// flag.Usage = Usage
	// flag.Parse()

	// manage plane seats:
	if len(CommandLine) < 3 || len(CommandLine) > 4 {
		PrintDefault()
		log.Fatal("Invalid Command")
	}

	planeBooking := make(map[string]string)
	var catch error

	planeBooking, catch = InitBooking(planeBooking)
	if catch != nil {
		fmt.Printf("FAIL")
		log.Fatal(catch)
	}

	bookingFile, err := os.OpenFile("booking.csv", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	var seatCount int = 1

	switch CommandLine[1] {
	case "BOOK":
		if len(CommandLine) == 4 {
			seatCount, err = strconv.Atoi(CommandLine[3])
			if err != nil {
				fmt.Printf("FAIL")
				log.Fatal(err)
			}
		}
		status, err := BookSeat(planeBooking, CommandLine[2], seatCount)
		if err != nil {
			fmt.Print(status)
			log.Fatal(err)
		}
		err = UpdateBookingDB(planeBooking, bookingFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(status)
	case "CANCEL":
		if len(CommandLine) == 4 {
			fmt.Print("FAIL")
			log.Fatal("Invalid Cancellation Command - consecutive seat counts provided")
		}
		status, err := CancelSeat(planeBooking, CommandLine[2])
		if err != nil {
			fmt.Print(status)
			log.Fatal(err)
		}
		err = UpdateBookingDB(planeBooking, bookingFile)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(status)
	default:
		fmt.Printf("FAIL")
		log.Fatal("Invalid Command")
	}
}
