package main

import (
	"time"
	"errors"
)

func main() {

}

/* Returns a DateRange containing the range of overlap of ranges passed in
	Or error if ranges do not overlap
 */

func findOverlap(dateRange1 DateRange, dateRange2 DateRange) (DateRange, error) {
	var overlap DateRange
	var err error

	if dateRange1.begin.After(dateRange2.begin) {
		overlap.begin = dateRange1.begin

	} else {
		overlap.begin = dateRange2.begin
	}

	if dateRange1.end.Before(dateRange2.end) {
		overlap.end = dateRange1.end
	} else {
		overlap.end = dateRange2.end
	}

	if overlap.end.Before(overlap.begin) || overlap.end.Equal(overlap.begin) {
		err = errors.New("No Overlap")
	}

	return overlap, err
}

type DateRange struct {
	begin time.Time
	end   time.Time
}
