package calendar

import (
	"fmt"
	"strings"
)

type CalendarWeek uint
type Day struct {
	current bool
	edge    bool
	special bool
	value   uint
}

func (d Day) IsCurrent() bool { return d.current }
func (d Day) IsEdge() bool    { return d.edge }
func (d Day) IsSpecial() bool { return d.special }
func (d Day) Day() uint       { return d.value }

type Month [][]interface{}

func (m Month) String() string {
	r := ""

	for i, row := range m {
		sRow := make([]string, len(row))
		for j := range row {
			switch v := row[j].(type) {
			case Day:
				sRow[j] = fmt.Sprintf("%2d", v.Day())
			default:
				if row[j] != nil {
					sRow[j] = fmt.Sprintf("%2v", row[j])[:2]
				} else {
					sRow[j] = "  "
				}
			}
		}

		if i > 0 {
			r += "\n"
		}
		r += strings.Join(sRow, " ")
	}

	return r
}
