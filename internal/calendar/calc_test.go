package calendar

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	testDate := time.Date(2022, 11, 2, 0, 0, 0, 0, time.UTC)

	result := NewMonth(testDate)

	println(result.String())

	assert.Equal(t, Month([][]interface{}{
		{nil, time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday, time.Sunday},
		{CalendarWeek(44), Day{edge: true, value: 31}, Day{value: 1}, Day{current: true, value: 2}, Day{value: 3}, Day{value: 4}, Day{value: 5}, Day{value: 6}},
		{CalendarWeek(45), Day{value: 7}, Day{value: 8}, Day{value: 9}, Day{value: 10}, Day{value: 11}, Day{value: 12}, Day{value: 13}},
		{CalendarWeek(46), Day{value: 14}, Day{value: 15}, Day{value: 16}, Day{value: 17}, Day{value: 18}, Day{value: 19}, Day{value: 20}},
		{CalendarWeek(47), Day{value: 21}, Day{value: 22}, Day{value: 23}, Day{value: 24}, Day{value: 25}, Day{value: 26}, Day{value: 27}},
		{CalendarWeek(48), Day{value: 28}, Day{value: 29}, Day{value: 30}, Day{edge: true, value: 1}, Day{edge: true, current: true, value: 2}, Day{edge: true, value: 3}, Day{edge: true, value: 4}},
	}), result)
}
