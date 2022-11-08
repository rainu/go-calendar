package calendar

import "time"

var DayOrder = map[time.Weekday]uint{
	time.Monday:    1,
	time.Tuesday:   2,
	time.Wednesday: 3,
	time.Thursday:  4,
	time.Friday:    5,
	time.Saturday:  6,
	time.Sunday:    7,
}

func NewMonth(now time.Time) Month {
	r := Month{}

	header := make([]interface{}, 8)
	for wd, i := range DayOrder {
		header[i] = wd
	}

	r = append(r, header)

	//set iteration day to start of calendar week
	iDay := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	startCw := isoWeek(iDay)
	for {
		prevDay := iDay.Add(-24 * time.Hour)
		if isoWeek(prevDay) != startCw {
			break
		}
		iDay = prevDay
	}

	for {
		cw := isoWeek(iDay)
		row := []interface{}{CalendarWeek(cw), nil, nil, nil, nil, nil, nil, nil}

		for {
			targetIndex := DayOrder[iDay.Weekday()]
			row[targetIndex] = Day{
				edge:    iDay.Month() != now.Month(),
				current: iDay.Day() == now.Day(),
				special: SpecialDays.Check(iDay),
				value:   uint(iDay.Day()),
			}

			iDay = iDay.Add(24 * time.Hour)
			if nextCw := isoWeek(iDay); nextCw != cw {
				break
			}
		}
		r = append(r, row)

		if iDay.Month() != now.Month() {
			break
		}
	}

	return r
}

func isoWeek(t time.Time) int {
	_, cw := t.ISOWeek()
	return cw
}
