package ctime

import (
	"testing"
	"time"
)

// The following tests use the same set of 5 intervals
// that are arranged as follows:
//    ..|---- a ----|........................
//    ......................|----- b -----|..
//    ..........|------ c ------|............
//    ..............|-- d --|................
//    ..|-------------- e ----------------|..

var a = Interval{Start: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 4, 0, 0, 0, 0, time.UTC)}
var b = Interval{Start: time.Date(2015, 1, 6, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 9, 0, 0, 0, 0, time.UTC)}
var c = Interval{Start: time.Date(2015, 1, 2, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 7, 0, 0, 0, 0, time.UTC)}
var d = Interval{Start: time.Date(2015, 1, 4, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 6, 0, 0, 0, 0, time.UTC)}
var e = Interval{Start: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 9, 0, 0, 0, 0, time.UTC)}

var trueAssertions = []bool{
	a.Before(b),

	b.After(a),

	a.Overlaps(c),

	b.OverlappedBy(c),

	c.During(e),
	d.During(e),
	d.During(c),

	a.Meets(d),
	d.Meets(b),

	d.MetBy(a),
	b.MetBy(d),

	a.Starts(e),

	e.StartedBy(a),

	b.Finishes(e),

	e.FinishedBy(b),

	a.Equals(a),
	b.Equals(b),
	c.Equals(c),
	d.Equals(d),
	e.Equals(e),

	a.Intersects(a),
	a.Intersects(c),
	a.Intersects(e),
	b.Intersects(b),
	b.Intersects(c),
	c.Intersects(a),
	c.Intersects(b),
	c.Intersects(c),
	c.Intersects(d),
	c.Intersects(e),
	d.Intersects(c),
	d.Intersects(d),
	d.Intersects(e),
	e.Intersects(d),
	e.Intersects(e),

	a.Contains(a.Start),
	a.Contains(c.Start),
	a.Contains(e.Start),
}

var falseAssertions = []bool{
	a.Before(a),
	a.Before(c),
	a.Before(d),
	a.Before(e),

	b.After(b),
	b.After(c),
	b.After(d),
	b.After(e),

	a.Overlaps(a),
	a.Overlaps(b),
	a.Overlaps(d),
	a.Overlaps(e),

	b.OverlappedBy(a),
	b.OverlappedBy(b),
	b.OverlappedBy(d),
	b.OverlappedBy(e),

	c.During(a),
	c.During(b),
	c.During(c),
	c.During(d),

	d.During(a),
	d.During(b),
	d.During(d),

	a.Meets(a),
	a.Meets(b),
	a.Meets(c),
	a.Meets(e),

	d.Meets(a),
	d.Meets(c),
	d.Meets(d),
	d.Meets(e),

	d.MetBy(b),
	d.MetBy(c),
	d.MetBy(d),
	d.MetBy(e),

	b.MetBy(a),
	b.MetBy(b),
	b.MetBy(c),
	b.MetBy(e),

	a.Starts(a),
	a.Starts(b),
	a.Starts(c),
	a.Starts(d),

	e.StartedBy(b),
	e.StartedBy(c),
	e.StartedBy(d),
	e.StartedBy(e),

	b.Finishes(a),
	b.Finishes(b),
	b.Finishes(c),
	b.Finishes(d),

	e.FinishedBy(a),
	e.FinishedBy(c),
	e.FinishedBy(d),
	e.FinishedBy(e),

	a.Intersects(b),
	a.Intersects(d),

	a.Intersects(b),
	a.Intersects(d),
	b.Intersects(a),
	b.Intersects(d),
	d.Intersects(a),
	d.Intersects(b),

	a.Contains(a.End),
	a.Contains(b.Start),
	a.Contains(d.Start),
}

func TestPredicates(t *testing.T) {
	for i := range trueAssertions {
		if !trueAssertions[i] {
			t.Errorf("true assertion %d was not satisfied", i)
		}
	}

	for i := range falseAssertions {
		if falseAssertions[i] {
			t.Errorf("false assertion %d was not satisfied", i)
		}
	}
}

func TestDuration(t *testing.T) {
	if a.Duration() != time.Hour*24*3 {
		t.Errorf("got %v, wanted %v", a.Duration(), time.Hour*24*4)
	}
}

func TestZeroDurationIntervals(t *testing.T) {
	testCases := []Interval{
		Interval{},
		Interval{Start: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC), End: time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC)},
	}

	comparisons := []Interval{a, b, c, d, e}

	for i, in := range testCases {
		if in.Duration() != 0 {
			t.Errorf("%d: got duration %v, wanted %v", i, in.Duration(), 0)
		}

		for _, co := range comparisons {
			if in.Before(co) {
				t.Errorf("%d: got before %v, wanted not", i, co)
			}

			if in.After(co) {
				t.Errorf("%d: got after %v, wanted not", i, co)
			}

			if in.Overlaps(co) {
				t.Errorf("%d: got overlaps %v, wanted not", i, co)
			}

			if in.OverlappedBy(co) {
				t.Errorf("%d: got overlapped by %v, wanted not", i, co)
			}

			if in.During(co) {
				t.Errorf("%d: got during %v, wanted not", i, co)
			}

			if in.Meets(co) {
				t.Errorf("%d: got meets %v, wanted not", i, co)
			}

			if in.MetBy(co) {
				t.Errorf("%d: got met by %v, wanted not", i, co)
			}

			if in.Starts(co) {
				t.Errorf("%d: got starts %v, wanted not", i, co)
			}

			if in.StartedBy(co) {
				t.Errorf("%d: got started by %v, wanted not", i, co)
			}

			if in.Finishes(co) {
				t.Errorf("%d: got finishes %v, wanted not", i, co)
			}

			if in.FinishedBy(co) {
				t.Errorf("%d: got finished by %v, wanted not", i, co)
			}

			if in.Intersects(co) {
				t.Errorf("%d: got intersects %v, wanted not", i, co)
			}
		}

	}
}
