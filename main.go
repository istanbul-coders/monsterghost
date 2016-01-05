package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

func InitializeClient(ckey string, csecret string, atoken string, asecret string) anaconda.TwitterApi {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csecret)
	api := anaconda.NewTwitterApi(atoken, asecret)
	fmt.Println(*api.Credentials)
	return *api
}

func IsTweeted(subject string, api anaconda.TwitterApi) bool {
	tweets, err := api.GetUserTimeline(nil)
	if err != nil {
		fmt.Println("Getting User timeline failed! Error : ", err)
		return true
	}
	for _, tweet := range tweets {
		fmt.Println(tweet.Text)
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
		return false
	}
	return true
}

func main() {
	var ckey string
	var csecret string
	var atoken string
	var asecret string
	var subject string

	flag.StringVar(&ckey, "ckey", "", "Consumer Key acquired from dev.twitter")
	flag.StringVar(&csecret, "csecret", "", "Consumer Secret acquired from dev.twitter")
	flag.StringVar(&atoken, "atoken", "", "Access token from dev.twitter")
	flag.StringVar(&asecret, "asecret", "", "Access secret from dev.twitter")
	flag.StringVar(&subject, "subject", "", "Istanbulcoders subject of the event")
	flag.Parse()

	if !flag.Parsed() {
		fmt.Println(flag.Args())
	}

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
