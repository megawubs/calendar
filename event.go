package calendar

import (
	"time"
	"io"
	"fmt"
)

type Event struct {
	UID       int64
	DTSTAMP   time.Time
	ORGANIZER string
	DTSTART   time.Time
	DTEND     time.Time
	SUMMARY   string
}

func NewEvent(uuid int64, organizer string, dateTimeStart time.Time, dateTimeEnd time.Time, summary string) Event{
	return Event{
		UID:       uuid,
		DTSTAMP:   time.Now(),
		ORGANIZER: organizer,
		DTSTART:   dateTimeStart,
		DTEND:     dateTimeEnd,
		SUMMARY:   summary,
	}
}

func (e *Event) Write(w io.Writer){
	dateLayout := "20060102T150405"
	fmt.Fprint(w, "BEGIN:VEVENT\r\n")
	fmt.Fprintf(w, "UID:%d\r\n", e.UID)
	fmt.Fprintf(w, "DTSTAMP:%s\r\n", e.DTSTAMP.Format(dateLayout))
	if e.ORGANIZER != ""{
		fmt.Fprintf(w, "ORGANIZER:%s\r\n", e.ORGANIZER)
	}

	fmt.Fprintf(w, "DTSTART:%s\r\n", e.DTSTART.Format(dateLayout))
	fmt.Fprintf(w, "DTEND:%s\r\n", e.DTEND.Format(dateLayout))
	fmt.Fprintf(w, "SUMMARY:%s\r\n", e.SUMMARY)
	fmt.Fprint(w, "END:VEVENT\r\n")
}