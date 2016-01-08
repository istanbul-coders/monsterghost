package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
	"github.com/memgo/api/meetup"
)

func InitializeClient(ckey string, csecret string, atoken string, asecret string) anaconda.TwitterApi {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csecret)
	api := anaconda.NewTwitterApi(atoken, asecret)
	return *api
}

func IsTweeted(subject string, api anaconda.TwitterApi) bool {
	tweets, err := api.GetUserTimeline(nil)
	if err != nil {
		fmt.Println("Getting User timeline failed! Error : ", err)
		os.Exit(1)
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
		fmt.Println("searching subject :" + subject)
		found := strings.Contains(tweet.Text, subject)

		if found {
			return true
		}
	}
	return false
}

func UpdateStatus(subject string, api anaconda.TwitterApi) bool {
	_, err := api.PostTweet(subject, nil)
	if err != nil {
		fmt.Println("Posting Tweet failed! Error : ", err)
		os.Exit(1)
	}
	return true
}

func IsEventCreated(desc string, apikey string) bool {
	url := fmt.Sprintf("https://api.meetup.com/2/events?key=%s&group_urlname=Istanbul-Hackers&sign=true&status=past,upcoming", apikey)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error occured during meetup search", err)
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()
	return strings.Contains(responseBody, desc)
}

func CreateEvent(apikey string, gid string, name string, vid string) string {
	url := fmt.Sprintf("https://api.meetup.com/2/event?key=%s&group_urlname=Istanbul-Hackers&group_id=%s&name=%s&sign=true&publish_status=draft&venue_id=%s", apikey, gid, name, vid)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", nil)
	if err != nil {
		fmt.Println("Error occured while creating meetup event", err)
		os.Exit(1)
	}
	fmt.Println("Meetup Create Event Response :", resp)

	event := new(meetup.Event)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(event)
	fmt.Println(event)

	return event.EventUrl
}

func initiateMeetup(desc string, apikey string, gid string, name string, vid string) {
	eventCreated := IsEventCreated(desc, apikey)
	fmt.Println("Meetup Event Created: ", eventCreated)

	CreateEvent(apikey, gid, name, vid)
	if eventCreated {
		os.Exit(0)
	}
}

func initiateTweet(ckey string, csecret string, atoken string, asecret string, subject string) {

	api := InitializeClient(ckey, csecret, atoken, asecret)
	tweeted := IsTweeted(subject, api)
	fmt.Print(tweeted)

	subjectPosted := true
	if !tweeted {
		subjectPosted = UpdateStatus(subject, api)
	}

	if !subjectPosted {
		os.Exit(1)
	}

}
func main() {
	commandType := os.Args[1]

	switch commandType {
	case "meetup":
		mySet := flag.NewFlagSet("", flag.ExitOnError)
		var desc = mySet.String("desc", "", "Description of the meetup event")
		var apikey = mySet.String("apikey", "", "meetup developer apikey")
		var gid = mySet.String("gid", "", "Groug id")
		var name = mySet.String("name", "", "Name of the event")
		var vid = mySet.String("vid", "", "Venue id")
		mySet.Parse(os.Args[2:])

		fmt.Println(mySet.Args())
		initiateMeetup(*desc, *apikey, *gid, *name, *vid)
	case "twitter":
		var ckey string
		var csecret string
		var atoken string
		var asecret string
		var subject string

		mySet := flag.NewFlagSet("", flag.ExitOnError)
		mySet.StringVar(&ckey, "ckey", "", "Consumer Key acquired from dev.twitter")
		mySet.StringVar(&csecret, "csecret", "", "Consumer Secret acquired from dev.twitter")
		mySet.StringVar(&atoken, "atoken", "", "Access token from dev.twitter")
		mySet.StringVar(&asecret, "asecret", "", "Access secret from dev.twitter")
		mySet.StringVar(&subject, "subject", "", "Istanbulcoders subject of the event")
		mySet.Parse(os.Args[2:])

		if !mySet.Parsed() {
			fmt.Println(mySet.Args())
		}

		initiateTweet(ckey, csecret, atoken, asecret, subject)
	}
}