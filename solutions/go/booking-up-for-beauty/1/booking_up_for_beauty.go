package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	var time, err = time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return time
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	var parsedTime, err = time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsedTime.Before(time.Now())

}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	var parsedTime, err = time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		panic(err)
	}
	return parsedTime.Hour() >= 12 && parsedTime.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	var parsedTime = Schedule(date)
	return parsedTime.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	var parsedTime, err = time.Parse("2/1/2006", fmt.Sprintf("15/09/%d", time.Now().Year()))
	if err != nil {
		panic(err)
	}
	return parsedTime
}
