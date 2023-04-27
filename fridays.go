package main

import (
	"fmt"
	time2 "time"
)

const (
	Mom = iota
	Dad = iota
)

type SpecialDay struct {
	parent int
	date   time2.Time
}

func (f *SpecialDay) previous() *SpecialDay {
	nextParent := 0
	if f.parent == 0 {
		nextParent = 1
	}
	nextDate := f.date.Add(time2.Hour * (24 * 21 * -1))
	next := &SpecialDay{
		parent: nextParent,
		date:   nextDate,
	}
	return next
}

func (f *SpecialDay) next() *SpecialDay {
	nextParent := 0
	if f.parent == 0 {
		nextParent = 1
	}
	nextDate := f.date.Add(time2.Hour * 24 * 21)
	next := &SpecialDay{
		parent: nextParent,
		date:   nextDate,
	}
	return next
}

func (f *SpecialDay) print() {
	parentDescription := ""
	if f.parent == Mom {
		parentDescription = "Mom"
	} else {
		parentDescription = "Dad"
	}

	fmt.Println("Parent: ", parentDescription, " Date: ", f.date)
}

func main() {

	seedDate := time2.Date(2022, time2.March, 18, 0, 0, 0, 0, time2.Local)
	occurance := &SpecialDay{date: seedDate, parent: Dad}
	occurance.print()

	for occurance.date.Year() < 2024 {
		occurance = occurance.next()
		occurance.print()
	}

}
