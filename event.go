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
	Allday    bool
}

func NewEvent(uuid int64, organizer string, dateTimeStart time.Time, dateTimeEnd time.Time, summary string) Event {
	return Event{
		UID:       uuid,
		DTSTAMP:   time.Now(),
		ORGANIZER: organizer,
		DTSTART:   dateTimeStart,
		DTEND:     dateTimeEnd,
		SUMMARY:   summary,
		Allday:    false,
	}
}

func NewAllDayEvent(uuid int64, organizer string, dateTimeStart time.Time, dateTimeEnd time.Time, summary string) Event {
	return Event{
		UID:       uuid,
		DTSTAMP:   time.Now(),
		ORGANIZER: organizer,
		DTSTART:   dateTimeStart,
		DTEND:     dateTimeEnd,
		SUMMARY:   summary,
		Allday:    true,
	}
}

func (e *Event) Write(w io.Writer) {
	dateLayout := "20060102T150405"
	allDayLayout := "20060102"
	fmt.Fprint(w, "BEGIN:VEVENT\r\n")
	fmt.Fprintf(w, "DTSTAMP:%s\r\n", e.DTSTAMP.Format(dateLayout))
	if e.ORGANIZER != "" {
		fmt.Fprintf(w, "ORGANIZER:%s\r\n", e.ORGANIZER)
	}
	if e.Allday == false {
		fmt.Fprintf(w, "DTSTART:%s\r\n", e.DTSTART.Format(dateLayout))
		fmt.Fprintf(w, "DTEND:%s\r\n", e.DTEND.Format(dateLayout))
	}
	if e.Allday == true {
		fmt.Fprintf(w, "DTSTART;VALUE=DATE:%s\r\n", e.DTSTART.Format(allDayLayout))
		fmt.Fprintf(w, "DTEND;VALUE=DATE:%s\r\n", e.DTEND.Format(allDayLayout))
	}

	fmt.Fprintf(w, "SUMMARY:%s\r\n", e.SUMMARY)
	fmt.Fprint(w, "END:VEVENT\r\n")
}
