package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Implementation for current plane seat capacity interface
// Plane seat capacity can be increased
var rows [20]string = [20]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T"}
var cols [8]string = [8]string{"0", "1", "2", "3", "4", "5", "6", "7"}

var planeSeats = make(map[string]string)
var availableSeat string = "AVAILB"
var bookedSeat string = "BOOKED"

// <summary>
// This function creates the file (DB) of seats booking on a plane, if file does not exist (initial call)
// Loads the booking information for booking, if file already exist (subsequent calls)
// <param name="planeSeats"></param>
// </summary>
func InitBooking(planeSeat map[string]string) (map[string]string, error) {
	// create fileDB filehandle AND check it exists
	// create or read file and update planeSeat data variable
	bookingDB, err := os.OpenFile("booking.csv", os.O_RDWR, 0644)

	if os.IsNotExist(err) {
		// if file does not exist: create
		bookingDB, err = os.Create("booking.csv")
		for i := range rows {
			for j := range cols {
				planeSeats[rows[i]+cols[j]] = availableSeat
			}
		}
		UpdateBookingDB(planeSeats, bookingDB)
	} else if err == nil {
		// if file exist: read
		reader := csv.NewReader(bookingDB)

		record, err := reader.ReadAll()
		if err != nil {
			return nil, err
		}
		// unmarshal into a map k:v data struct
		for _, field := range record {
			planeSeats[field[0]] = field[1]
		}
	} else {
		return nil, err
	}

	defer bookingDB.Close()

	return planeSeats, err
}

// <summary>
// This function writes the booking state of plane seats to storage file
// </summary>
func UpdateBookingDB(planeSeats map[string]string, bookingDB *os.File) error {
	// marshal the map into a csv file
	// write to file

	writer := csv.NewWriter(bookingDB)
	for seat, status := range planeSeats {
		err := writer.Write([]string{seat, status})
		if err != nil {
			return err
		}
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		return err
	}
	return nil
}

// <summary>
// This function books seats specified on a plane; 1 or multiple consecutive seats
// </summary>
func BookSeat(planeSeats map[string]string, seat string, count int) (string, error) {
	// update planeSeats map with booked seat
	seatAvail, err := SeatOpen(planeSeats, seat)
	if err != nil {
		return "FAIL", err
	}
	if !seatAvail {
		return "FAIL", fmt.Errorf("seat not available")
	}

	if count == 1 || count == 0 {
		planeSeats[seat] = bookedSeat
	} else if count > 1 {
		holder := strings.SplitN(seat, "", 2)
		seatRow := holder[0]
		seatNum, err := strconv.Atoi(holder[1])
		if err != nil {
			return "FAIL", err
		}
		for i := 0; i < count; i++ {
			seatAvail, err := SeatOpen(planeSeats, seatRow+strconv.Itoa(seatNum+i))
			if err != nil {
				return "FAIL", fmt.Errorf(err.Error())
			}
			if seatAvail {
				planeSeats[seatRow+strconv.Itoa(seatNum+i)] = bookedSeat
			} else {
				return "FAIL", fmt.Errorf("seat not available")
			}
		}
	} else {
		return "FAIL", fmt.Errorf("invalid number of seats")
	}
	return "SUCCESS", nil
}

// <summary>
// This function cancels seat already booked; 1 at a time
// </summary>
func CancelSeat(planeSeats map[string]string, seat string) (string, error) {
	// Improvement: Add count parameter to cancel multiple seats
	// cancel seat booking in planeSeats map
	seatAvail, err := SeatOpen(planeSeats, seat)
	if err != nil {
		return "FAIL", err
	}
	if seatAvail {
		return "FAIL", fmt.Errorf("seat not booked yet")
	}

	if !seatAvail {
		planeSeats[seat] = availableSeat
	}

	return "SUCCESS", nil
}

// <summary>
// Helper to check the status of plane seat
// Expect little to no memory hit with passing maps on the stack.
// Since maps are passed by reference in Go:
// https://github.com/golang/go/blob/db8142fb8631df3ee56983cbc13db997c16f2f6f/src/runtime/map.go#L298C4-L298C12
// </summary>
func SeatOpen(planeSeats map[string]string, seat string) (bool, error) {
	if planeSeats[seat] == availableSeat {
		return true, nil
	} else if planeSeats[seat] == bookedSeat {
		return false, nil
	} else {
		return false, fmt.Errorf("invalid seat")
	}
}

// Print exmaple usage for application
func PrintDefault() {
	fmt.Fprintf(os.Stderr, "try: %s %s %s %s", os.Args[0], "BOOK", "A0", "1")
}
