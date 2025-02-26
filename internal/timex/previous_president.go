package timex

import (
	"time"
)

func SinceLastPresident() PresTracker {
	return timeSince(inTime.In(eastern))
}

func DateOfPreviousPres() time.Time {
	return inTime.In(eastern)
}

func timeSince(t time.Time) PresTracker {
	now := time.Now()
	years := now.Year() - t.Year()
	months := int(now.Month()) - int(t.Month())

	days := now.Day() - t.Day()
	hours := now.Hour() - t.Hour()
	minutes := now.Minute() - t.Minute()
	seconds := now.Second() - t.Second()

	if seconds < 0 {
		seconds += 60
		minutes--
	}
	if minutes < 0 {
		minutes += 60
		hours--
	}
	if hours < 0 {
		hours += 24
		days--
	}
	if days < 0 {
		lastMonth := now.AddDate(0, 0, -now.Day()) // Get last month's last day
		days += lastMonth.Day()
		months--
	}
	if months < 0 {
		months += 12
		years--
	}

	return PresTracker{
		Years:   years,
		Months:  months,
		Days:    days,
		Hours:   hours,
		Minutes: minutes,
		Seconds: seconds,
	}
}
