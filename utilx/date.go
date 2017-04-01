package utilx

import "time"

func TrimMinTime(t time.Time)  time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func TrimMaxTime(t time.Time)  time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}

func TrimTime(t time.Time, hour int, minute int, second int)  time.Time {
	if hour < 0 || hour > 23{
		hour = t.Hour()
	}
	if minute < 0 || minute > 59{
		minute = t.Minute()
	}
	if second < 0 || second > 59{
		second = t.Second()
	}
	return time.Date(t.Year(), t.Month(), t.Day(), hour, minute, second, 0, t.Location())
}

func ParserDate(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02", date, time.Local)
}

func ParserTime(date string) (time.Time, error) {
	return time.ParseInLocation("2006-01-02 15:04:05", date, time.Local)
}


func FormatDate(date time.Time) string {
	return date.Local().Format("2006-01-02")
}

func FormatTime(date time.Time) string {
	return date.Local().Format("2006-01-02 15:04:05")
}

var weekydays = [...]string{
	"星期日",
	"星期一",
	"星期二",
	"星期三",
	"星期四",
	"星期五",
	"星期六",
}

func WeekdayName(w time.Weekday) string {
	return weekydays[w]
}
