package timex

import "fmt"

type PresTracker struct {
	Years   int `json:"years"`
	Months  int `json:"months"`
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

func (p PresTracker) String() string {
	return fmt.Sprintf("years: %d, months: %d days: %d, hours: %d, minutes: %d, seconds: %d",
		p.Years, p.Months, p.Days, p.Hours, p.Minutes, p.Seconds)
}
