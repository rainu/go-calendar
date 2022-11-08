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
		{CalendarWeek(44), EdgeDay(31), Day(1), CurrentDay(2), Day(3), Day(4), Day(5), Day(6)},
		{CalendarWeek(45), Day(7), Day(8), Day(9), Day(10), Day(11), Day(12), Day(13)},
		{CalendarWeek(46), Day(14), Day(15), Day(16), Day(17), Day(18), Day(19), Day(20)},
		{CalendarWeek(47), Day(21), Day(22), Day(23), Day(24), Day(25), Day(26), Day(27)},
		{CalendarWeek(48), Day(28), Day(29), Day(30), EdgeDay(1), EdgeDay(2), EdgeDay(3), EdgeDay(4)},
	}), result)
}
