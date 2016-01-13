package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/memgo/api/meetup"
)

func isEventCreated(name string, apikey string) bool {
	url := fmt.Sprintf("https://api.meetup.com/2/events?key=%s&group_urlname=Istanbul-Hackers&sign=true&status=past,upcoming", apikey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured during meetup search", err)
		os.Exit(1)
	}
	events := new(meetup.Events)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(events)

	for _, event := range events.Results {
		fmt.Println("Response event name: ", event.Name)
		fmt.Println("New event name: ", name)
		if strings.Contains(event.Name, name) {
			return true
		}
	}

	return false
}

func createEvent(apikey string, gid string, name string, desc string, vid string, rsvp_limit string, epocs string) *meetup.Event {
	meetup_url := "https://api.meetup.com/2/event"

	key := fmt.Sprintf("?key=%s", apikey)
	meetup_url = fmt.Sprint(meetup_url, key)

	groupUrlName := "&group_urlname=Istanbul-Hackers"
	meetup_url = fmt.Sprint(meetup_url, groupUrlName)

	groupId := fmt.Sprintf("&group_id=%s", gid)
	meetup_url = fmt.Sprint(meetup_url, groupId)

	venue := fmt.Sprintf("&venue_id=%s", vid)
	meetup_url = fmt.Sprint(meetup_url, venue)

	rsvp_limit = fmt.Sprintf("&rsvp_limit=%s", rsvp_limit)
	meetup_url = fmt.Sprint(meetup_url, rsvp_limit)

	epocs_in_ms, _ := time.Parse(time.RFC1123Z, epocs)
	epocs_txt := fmt.Sprintf("&time=%d", (epocs_in_ms.UnixNano() / int64(time.Millisecond)))
	fmt.Println("Epocs in txt: ", epocs_txt)
	meetup_url = fmt.Sprint(meetup_url, epocs_txt)

	name = fmt.Sprintf("&name=%s", url.QueryEscape(name))
	meetup_url = fmt.Sprint(meetup_url, name)

	description := fmt.Sprintf("&description=%s", url.QueryEscape(desc))
	meetup_url = fmt.Sprint(meetup_url, description)

	fmt.Println("Url :", meetup_url)
	resp, err := http.Post(meetup_url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error occured while creating meetup event", err)
		os.Exit(1)
	}
	fmt.Println("Post Response:", resp)
	event := new(meetup.Event)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(event)
	fmt.Println(event)

	return event
}

func initiateMeetup(desc string, apikey string, gid string, name string, vid string, rsvp_limit string, time string) {
	eventCreated := isEventCreated(name, apikey)
	fmt.Println("Meetup Event Created? : ", eventCreated)

	if eventCreated {
		os.Exit(0)
	}

	fmt.Println("Creating event with following parameters:")
	fmt.Println("Desc: ", desc)
	fmt.Println("Name: ", name)
	fmt.Println("Time: ", time)
	fmt.Println("Guest Limit: ", rsvp_limit)
	createEvent(apikey, gid, name, desc, vid, rsvp_limit, time)
}
