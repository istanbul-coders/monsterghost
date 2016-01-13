package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		fmt.Println("Error :", e)
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
		var rsvp_limit = mySet.String("rsvp_limit", "", "Rsvp Limit")
		var time = mySet.String("time", "", "Time for the event")
		mySet.Parse(os.Args[2:])

		if !mySet.Parsed() {
			fmt.Println("Error parsing arguments:", mySet.Args())
		}

		eventUrl := initiateMeetup(*desc, *apikey, *gid, *name, *vid, *rsvp_limit, *time)
		eventurlfile := []byte(eventUrl)
		err := ioutil.WriteFile("eventurl.md", eventurlfile, 0644)
		check(err)

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
			fmt.Println("Error parsing arguments:", mySet.Args())
		}

		initiateTweet(ckey, csecret, atoken, asecret, subject)
	case "gmail":
		var to string
		var from string
		var subject string
		var body string
		var username string
		var password string

		mySet := flag.NewFlagSet("", flag.ExitOnError)
		mySet.StringVar(&to, "to", "", "email to")
		mySet.StringVar(&from, "from", "", "email from")
		mySet.StringVar(&subject, "subject", "", "email subject")
		mySet.StringVar(&body, "body", "", "email body")
		mySet.StringVar(&username, "username", "", "email username")
		mySet.StringVar(&password, "password", "", "email password")
		mySet.Parse(os.Args[2:])

		if !mySet.Parsed() {
			fmt.Println("Error parsing arguments:", mySet.Args())
		}

		sendEmail(to, from, subject, body, username, password)
	}
}
