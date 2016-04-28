package main

import (
	"google.golang.org/api/calendar/v3"
	"log"
	"time"
	"net/http"
)

const username = "hevalberknevruz@gmail.com"

type EventInformation struct {
	date        string
	title       string
	description string
	location    string
}

func insertEventToCalendar(service *calendar.Service, information EventInformation) {
	eventMap := getEventListInCalendar(service)
	event := &calendar.Event{
		Summary:information.title,
		Location:information.location,
		Description:information.description,
		Start:&calendar.EventDateTime{DateTime:information.date},
		End:&calendar.EventDateTime{DateTime:information.date},
	}

	isFuture := isFutureDate(information.date)

	if isFuture {

		for id, date := range eventMap {
			if date == information.date {
				service.Events.Update(username, id, event).Do()
				return
			}
		}

		service.Events.Insert(username, event).Do()
	}
}

func getEventListInCalendar(service *calendar.Service) map[string]string {
	t := time.Now().Format(time.RFC3339)
	events, err := service.Events.List("primary").ShowDeleted(false).
	SingleEvents(true).TimeMin(t).Do()
	whenMap := make(map[string]string)

	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events. %v", err)
	}

	if len(events.Items) > 0 {

		for _, i := range events.Items {

			var when string

			if i.Start.DateTime != "" {
				when = i.Start.DateTime
			} else {
				when = i.Start.Date
			}
			whenMap[i.Id] = when

		}

	}
	return whenMap
}

func createCalendarService(client *http.Client) (*calendar.Service, error) {
	return calendar.New(client)
}
