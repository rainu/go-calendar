package main

import (
	"flag"
	"github.com/rainu/go-calendar/internal/calendar"
	"github.com/rainu/go-calendar/internal/ui"
	"log"
	"time"
)

func ParseFlags() {
	var oMonday, oTuesday, oWednesday, oThursday, oFriday, oSaturday, oSunday uint
	var nMonday, nTuesday, nWednesday, nThursday, nFriday, nSaturday, nSunday string
	var nJanuary, nFebruary, nMarch, nApril, nMay, nJune, nJuly, nAugust, nSeptember, nOctober, nNovember, nDecember string

	flag.UintVar(&oMonday, "omo", calendar.DayOrder[time.Monday], "The order for Monday")
	flag.UintVar(&oTuesday, "otu", calendar.DayOrder[time.Tuesday], "The order for Tuesday")
	flag.UintVar(&oWednesday, "owe", calendar.DayOrder[time.Wednesday], "The order for Wednesday")
	flag.UintVar(&oThursday, "oth", calendar.DayOrder[time.Thursday], "The order for Thursday")
	flag.UintVar(&oFriday, "ofr", calendar.DayOrder[time.Friday], "The order for Friday")
	flag.UintVar(&oSaturday, "osa", calendar.DayOrder[time.Saturday], "The order for Saturday")
	flag.UintVar(&oSunday, "osu", calendar.DayOrder[time.Sunday], "The order for Sunday")

	flag.StringVar(&nMonday, "nmo", ui.DayNames[time.Monday], "The name for Monday")
	flag.StringVar(&nTuesday, "ntu", ui.DayNames[time.Tuesday], "The name for Tuesday")
	flag.StringVar(&nWednesday, "nwe", ui.DayNames[time.Wednesday], "The name for Wednesday")
	flag.StringVar(&nThursday, "nth", ui.DayNames[time.Thursday], "The name for Thursday")
	flag.StringVar(&nFriday, "nfr", ui.DayNames[time.Friday], "The name for Friday")
	flag.StringVar(&nSaturday, "nsa", ui.DayNames[time.Saturday], "The name for Saturday")
	flag.StringVar(&nSunday, "nsu", ui.DayNames[time.Sunday], "The name for Sunday")

	flag.StringVar(&nJanuary, "njan", ui.MonthNames[time.January], "The name for January")
	flag.StringVar(&nFebruary, "nfeb", ui.MonthNames[time.February], "The name for February")
	flag.StringVar(&nMarch, "nmar", ui.MonthNames[time.March], "The name for March")
	flag.StringVar(&nApril, "napr", ui.MonthNames[time.April], "The name for April")
	flag.StringVar(&nMay, "nmay", ui.MonthNames[time.May], "The name for May")
	flag.StringVar(&nJune, "njun", ui.MonthNames[time.June], "The name for June")
	flag.StringVar(&nJuly, "njul", ui.MonthNames[time.July], "The name for July")
	flag.StringVar(&nAugust, "naug", ui.MonthNames[time.August], "The name for August")
	flag.StringVar(&nSeptember, "nsep", ui.MonthNames[time.September], "The name for September")
	flag.StringVar(&nOctober, "noct", ui.MonthNames[time.October], "The name for October")
	flag.StringVar(&nNovember, "nnov", ui.MonthNames[time.November], "The name for November")
	flag.StringVar(&nDecember, "ndec", ui.MonthNames[time.December], "The name for December")

	flag.Parse()

	calendar.DayOrder[time.Monday] = oMonday
	calendar.DayOrder[time.Tuesday] = oTuesday
	calendar.DayOrder[time.Wednesday] = oWednesday
	calendar.DayOrder[time.Thursday] = oThursday
	calendar.DayOrder[time.Friday] = oFriday
	calendar.DayOrder[time.Saturday] = oSaturday
	calendar.DayOrder[time.Sunday] = oSunday

	ui.DayNames[time.Monday] = nMonday
	ui.DayNames[time.Tuesday] = nTuesday
	ui.DayNames[time.Wednesday] = nWednesday
	ui.DayNames[time.Thursday] = nThursday
	ui.DayNames[time.Friday] = nFriday
	ui.DayNames[time.Saturday] = nSaturday
	ui.DayNames[time.Sunday] = nSunday

	ui.MonthNames[time.January] = nJanuary
	ui.MonthNames[time.February] = nFebruary
	ui.MonthNames[time.March] = nMarch
	ui.MonthNames[time.April] = nApril
	ui.MonthNames[time.May] = nMay
	ui.MonthNames[time.June] = nJune
	ui.MonthNames[time.July] = nJuly
	ui.MonthNames[time.August] = nAugust
	ui.MonthNames[time.September] = nSeptember
	ui.MonthNames[time.October] = nOctober
	ui.MonthNames[time.November] = nNovember
	ui.MonthNames[time.December] = nDecember

	sum := uint(0)
	for _, u := range calendar.DayOrder {
		if u > 7 {
			log.Fatalf("Invalid day order given: %d", u)
		}
		sum += u
	}
	if sum != 28 {
		log.Fatal("Invalid day order setup given")
	}
}
