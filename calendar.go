package calendar

import (
	"io"
	"bytes"
	"fmt"
)

type Calendar struct {
	Version string
	ProId   string
	Events  []Event
}

func NewCalendar(programId string) Calendar {
	return Calendar{Version: "2.0", ProId: programId,}
}

func (c *Calendar) Add(event Event) {
	c.Events = append(c.Events, event)
}

func (c *Calendar) Write(w io.Writer) {
	var buffer bytes.Buffer
	fmt.Fprint(w, "BEGIN:VCALENDAR\r\n")
	fmt.Fprintf(w, "PRODID:%s\r\n", c.ProId)
	fmt.Fprintf(w, "VERSION:%s\r\n", c.Version)

	for _, event := range c.Events {
		event.Write(w)
	}

	fmt.Fprint(w, "END:VCALENDAR\r\n")
	buffer.WriteTo(w)
}
