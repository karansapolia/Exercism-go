package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var filteredRecords []Record
	for _, value := range in {
		accepted := predicate(value)
		if accepted {
			filteredRecords = append(filteredRecords, value)
		}
	}
	return filteredRecords
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	totalAmount := 0.0
	daysInPeriod := Filter(in, ByDaysPeriod(p))
	for _, record := range daysInPeriod {
		totalAmount += record.Amount
	}
	return totalAmount
}

// CategoryExpenses returns total amount of expenses for records
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	if len(Filter(in, ByCategory(c))) == 0 {
		return 0, fmt.Errorf("unknown category %s", c)
	}
	totalAmount := 0.0
	daysInPeriod := Filter(in, ByDaysPeriod(p))
	daysInCategoryAndPeriod := Filter(daysInPeriod, ByCategory(c))
	for _, record := range daysInCategoryAndPeriod {
		totalAmount += record.Amount
	}
	return totalAmount, nil
}
