package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
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
		fmt.Println("Error occured during meetup search")
		os.Exit(1)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	responseBody := buf.String()
	return strings.Contains(responseBody, desc)
}

func initiateMeetup(desc string, apikey string) {
	eventCreated := IsEventCreated(desc, apikey)
	fmt.Println("Meetup Event Created: ", eventCreated)
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
		fmt.Println("Meetup")
		mySet := flag.NewFlagSet("", flag.ExitOnError)
		var desc = mySet.String("desc", "", "Description of the meetup event")
		var apikey = mySet.String("apikey", "", "meetup developer apikey")
		mySet.Parse(os.Args[2:])

		initiateMeetup(*desc, *apikey)
	case "twitter":
		fmt.Println("Twitter")
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
