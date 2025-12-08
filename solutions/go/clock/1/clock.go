package clock

import (
	"time"
)

type Clock struct {
	h int
	m int
}

func parseTime(h, m int) time.Time {
	return time.Time{}.Add(time.Hour * time.Duration(h)).Add(time.Minute * time.Duration(m))
}

func New(h, m int) Clock {
	t := parseTime(h,m)
	return Clock{
		t.Hour(),
		t.Minute(),
	}
}

func (c Clock) Add(m int) Clock {
	t := parseTime(c.h, c.m)
	t = t.Add(time.Minute * time.Duration(m))
	c.h = t.Hour()
	c.m = t.Minute()
	return c
}

func (c Clock) Subtract(m int) Clock {
	t := parseTime(c.h, c.m)
	t = t.Add(time.Minute * time.Duration(m) * -1)
	c.h = t.Hour()
	c.m = t.Minute()
	return c
}

func (c Clock) String() string {
	t := parseTime(c.h, c.m)
	return t.Format("15:04")
}
