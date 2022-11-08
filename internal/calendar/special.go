package calendar

import "time"

var SpecialDays SpecialCollection

type SpecialCollection struct {
	Collection []interface{ Check(t time.Time) bool }
}

type SpecialOnce struct {
	Date time.Time
}

type SpecialEachYear struct {
	Day   uint
	Month time.Month
}

type SpecialEachMonth struct {
	Day uint
}

func (s SpecialCollection) Check(t time.Time) bool {
	for _, i := range s.Collection {
		if i.Check(t) {
			return true
		}
	}
	return false
}

func (s SpecialOnce) Check(t time.Time) bool {
	return s.Date.Day() == t.Day() &&
		s.Date.Month() == t.Month() &&
		s.Date.Year() == t.Year()
}

func (s SpecialEachYear) Check(t time.Time) bool {
	return s.Day == uint(t.Day()) &&
		s.Month == t.Month()
}

func (s SpecialEachMonth) Check(t time.Time) bool {
	return s.Day == uint(t.Day())
}
