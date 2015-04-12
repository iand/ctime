// Package ctime implements retaions between crisp time intervals
package ctime

import (
	"time"
)

// An Interval is a period of time with a defined start and end
type Interval struct {
	Start time.Time
	End   time.Time
}

// Duration returns the duration of the time interval i.
func (i Interval) Duration() time.Duration {
	return i.End.Sub(i.Start)
}

// Before reports whether the time interval i ends before time interval n starts.
//    ..|------ i ------|.....................
//    .....................|------ n ------|..
func (i Interval) Before(n Interval) bool {
	return i.End.Before(n.Start)
}

// After reports whether the time interval i starts after time interval n ends.
//    .....................|------ i ------|..
//    ..|------ n ------|.....................
func (i Interval) After(n Interval) bool {
	return n.Before(i)
}

// Overlaps reports whether the time interval i starts before time interval n starts
// and ends after n starts but before n ends.
//    ....|------ i ------|...................
//    ...............|------ n ------|........
func (i Interval) Overlaps(n Interval) bool {
	end := i.End
	return i.Start.Before(n.Start) && n.Start.Before(end) && end.Before(n.End)
}

// OverlappedBy reports whether the time interval i ends after time interval n ends
// and starts after time interval n starts but before n ends.
//    ...............|------ i ------|........
//    ....|------ n ------|...................
func (i Interval) OverlappedBy(n Interval) bool {
	return n.Overlaps(i)
}

// During reports whether the time interval i starts after time interval n
// starts and ends before n ends.
//    ......|---- i ----|.....................
//    ....|------ n ------|...................
func (i Interval) During(n Interval) bool {
	return n.Start.Before(i.Start) && i.End.Before(n.End)
}

// Meets reports whether the time interval i ends at the same time that
// time interval n starts.
//    ..|------ i ------|.....................
//    ..................|------ n ------|.....
func (i Interval) Meets(n Interval) bool {
	return i.End == n.Start
}

// MetBy reports whether the time interval i starts at the same time that
// time interval n ends.
//    ..................|------ i ------|.....
//    ..|------ n ------|.....................
func (i Interval) MetBy(n Interval) bool {
	return n.Meets(i)
}

// Starts reports whether the time interval i starts at the same time that
// time interval n starts but ends before n ends.
//    ..|---- i ----|.........................
//    ..|------ n ------|.....................
func (i Interval) Starts(n Interval) bool {
	return i.Start == n.Start && i.End.Before(n.End)
}

// StartedBy reports whether the time interval i starts at the same time that
// time interval n starts but ends after n ends.
//    ..|------ i ------|.........................
//    ..|---- n ----|.....................
func (i Interval) StartedBy(n Interval) bool {
	return n.Starts(i)
}

// Finishes reports whether the time interval i ends at the same time that
// time interval n ends but starts after n starts.
//    ......|---- i ----|.....................
//    ..|------ n ------|.....................
func (i Interval) Finishes(n Interval) bool {
	return n.Start.Before(i.Start) && i.End == n.End
}

// FinishedBy reports whether the time interval i ends at the same time that
// time interval n ends but starts before n starts.
//    ..|------ i ------|.....................
//    ......|---- n ----|.....................
func (i Interval) FinishedBy(n Interval) bool {
	return n.Finishes(i)
}

// Equals reports whether the time interval i ends at the same time that
// time interval n ends and starts at the same time n starts.
//    ..|------ i ------|.....................
//    ..|------ n ------|.....................
func (i Interval) Equals(n Interval) bool {
	return i.Start == n.Start && i.End == n.End
}

// Intersects reports whether the time interval i ends after time interval n
// starts and starts before n ends.
//    ..|------ i ------|.....................
//    ......|------ n ------|.................
func (i Interval) Intersects(n Interval) bool {
	return n.Start.Before(i.End) && i.Start.Before(n.End)
}
