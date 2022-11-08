package calendar

import (
	"fmt"
	"strings"
)

type CalendarWeek uint
type Day uint
type CurrentDay uint
type EdgeDay uint

type Month [][]interface{}

func (m Month) String() string {
	r := ""

	for i, row := range m {
		sRow := make([]string, len(row))
		for j := range row {
			if row[j] != nil {
				sRow[j] = fmt.Sprintf("%2v", row[j])[:2]
			} else {
				sRow[j] = "  "
			}
		}

		if i > 0 {
			r += "\n"
		}
		r += strings.Join(sRow, " ")
	}

	return r
}
