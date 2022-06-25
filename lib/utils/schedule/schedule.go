package schedule

import (
	"errors"
	"time"
)

func NextWeekday(now time.Time, weekday time.Weekday) time.Time {
	diff := int(weekday) - int(now.Weekday())

	if diff <= 0 {
		diff = diff + 7
	}

	return now.AddDate(0, 0, diff)
}

func AfterBegginingMonth(now time.Time, month int) time.Time {
	// monthヶ月後の月の１日目を取得
	return time.Date(now.Year(), now.Month()+time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}

func WeekNumber(now time.Time) int {
	return (now.Day()-1)/7 + 1
}

func AddWeek(now time.Time, week int) time.Time {
	return now.AddDate(0, 0, 7*week)
}

func BegginingHour(now time.Time, hour int) (time.Time, error) {
	if hour > 24 {
		return now, errors.New("too big hour")
	}
	return now.Truncate(time.Duration(hour) * time.Hour), nil
}
