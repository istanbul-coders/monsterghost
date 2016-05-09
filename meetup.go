package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/memgo/api/meetup"
)

func isEventCreated(name string, apikey string) bool {
	url := fmt.Sprintf("https://api.meetup.com/2/events?key=%s&group_urlname=Istanbul-Hackers&sign=true&status=past,upcoming", apikey)

	resp, err := http.Get(url)

	if err != nil {
		log.Fatalln("Error occured during meetup search", err)
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
	api_url, _ := url.Parse("https://api.meetup.com/2/events/")

	query := api_url.Query()
	query.Set("key", apikey)
	query.Set("group_urlname", "Istanbul-Hackers")
	query.Set("group_id", gid)
	query.Set("venue_id", vid)
	query.Set("rsvp_limit", rsvp_limit)
	query.Set("name", url.QueryEscape(name))
	query.Set("description", url.QueryEscape(desc))

	eventTime, err := time.Parse(time.RFC1123Z, epocs)

	if err != nil {
		log.Fatalln("Could NOT parse event date/time:", err)
	}

	query.Set("time", strconv.FormatInt(eventTime.UnixNano()/int64(time.Millisecond), 10))

	api_url.RawQuery = query.Encode()

	resp, err := http.Get(api_url.String())

	if err != nil {
		log.Fatalln("Client error while performing GET request on Meetup API:", err)
	}

	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	log.Println("Meetup API Response", string(respBody))

	if resp.StatusCode != http.StatusOK {
		log.Fatalln("Unsuccessful HTTP Status from Meetup API:", resp.StatusCode)
	}

	event := new(meetup.Event)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(event)

	log.Println("Meetup event has been created successfully:", event)

	return event
}

func initiateMeetup(desc string, apikey string, gid string, name string, vid string, rsvp_limit string, time string) string {
	eventCreated := isEventCreated(name, apikey)
	log.Println("Meetup Event Created? : ", eventCreated)

	if eventCreated {
		os.Exit(0) //FIXME This is very confusing
	}

	log.Println("Creating event with following parameters:")
	log.Println("Desc: ", desc)
	log.Println("Name: ", name)
	log.Println("Time: ", time)
	log.Println("Guest Limit: ", rsvp_limit)
	event := createEvent(apikey, gid, name, desc, vid, rsvp_limit, time)

	return event.EventUrl
}
