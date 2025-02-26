package timex

import "time"

var (
	inTime  time.Time
	outTime time.Time
	eastern *time.Location
)

func init() {
	tz, err := time.LoadLocation("EST")
	if err != nil {
		panic(err)
	}
	eastern = tz
	inauguration := time.
		Date(2025, 1, 20, 11, 0, 0, 0, tz).
		In(eastern)
	inTime = inauguration
	outTime = inTime
	now := time.Now()
	for {
		if outTime.After(now) {
			break
		}
		outTime = inTime.AddDate(4, 0, 0)
	}
}
