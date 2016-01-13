package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ChimeraCoder/anaconda"
)

func initializeClient(ckey string, csecret string, atoken string, asecret string) anaconda.TwitterApi {
	anaconda.SetConsumerKey(ckey)
	anaconda.SetConsumerSecret(csecret)
	api := anaconda.NewTwitterApi(atoken, asecret)
	return *api
}

func isTweeted(subject string, api anaconda.TwitterApi) bool {
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

func updateStatus(subject string, api anaconda.TwitterApi) bool {
	_, err := api.PostTweet(subject, nil)
	if err != nil {
		fmt.Println("Posting Tweet failed! Error : ", err)
		os.Exit(1)
	}
	return true
}

func initiateTweet(ckey string, csecret string, atoken string, asecret string, subject string) {

	api := initializeClient(ckey, csecret, atoken, asecret)
	tweeted := isTweeted(subject, api)
	fmt.Print(tweeted)

	subjectPosted := true
	if !tweeted {
		subjectPosted = updateStatus(subject, api)
	}

	if !subjectPosted {
		os.Exit(1)
	}

}
