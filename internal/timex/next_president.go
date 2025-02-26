package timex

import (
	"time"
	_ "time/tzdata"
)

func DateOfNextPres() time.Time {
	return outTime.In(eastern)
}

func TillNextPresident() PresTracker {
	return timeTill(outTime.In(eastern))
}

func timeTill(t time.Time) PresTracker {
	now := time.Now().In(eastern)
	if t.Before(now) {
		return PresTracker{}
	}
	years := t.Year() - now.Year()
	months := int(t.Month()) - int(now.Month())
	days := t.Day() - now.Day()
	hours := t.Hour() - now.Hour()
	minutes := t.Minute() - now.Minute()
	seconds := t.Second() - now.Second()

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
		prevMonth := t.AddDate(0, -1, 0)
		days += prevMonth.Day() - now.Day()
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
