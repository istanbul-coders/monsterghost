package main

import (
	"fmt"
	"os"
	"google.golang.org/api/drive/v2"
	"log"
	"flag"
	"io/ioutil"
)

const (
	EVENT_LOCATION = "Bahcesehir Universitesi Besiktas Kampus Ciragan Caddesi Osmanpasa Mektebi Sokak No: 4 - D504 Nolu sınıf"
	EVENT_FILE_NAME = "Event Schedule"
)

var environmentName = os.Getenv("HOME")

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
<<<<<<< HEAD
		eventurlfile := []byte(eventUrl)
		err := ioutil.WriteFile("eventurl.md", eventurlfile, 0644)
=======
		eventUrlFile := []byte(eventUrl)
		err := ioutil.WriteFile("eventurl.md", eventUrlFile, 0644)
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
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
		var cc string
		var from string
		var subject string
		var body string
		var username string
		var password string

		mySet := flag.NewFlagSet("", flag.ExitOnError)
		mySet.StringVar(&to, "to", "", "email to")
		mySet.StringVar(&cc, "cc", to, "email cc")
		mySet.StringVar(&from, "from", "", "email from")
		mySet.StringVar(&subject, "subject", "", "email subject")
		mySet.StringVar(&body, "body", "", "email body")
		mySet.StringVar(&username, "username", "", "email username")
		mySet.StringVar(&password, "password", "", "email password")
		mySet.Parse(os.Args[2:])

		if !mySet.Parsed() {
			fmt.Println("Error parsing arguments:", mySet.Args())
		}

		sendEmail(to, cc, from, subject, body, username, password)

	case "calendar":
<<<<<<< HEAD

=======
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
		client := createGoogleApiClient()

		driveService, err := createDriveService(client)

		if err != nil {
			log.Fatalf("Can not instantiate google drive service client: %v", err)
		}

		transport := createTransport()

		downloadFile(driveService, transport, &drive.File{
			DownloadUrl:getDownloadUrlByName(driveService, EVENT_FILE_NAME),
		})

		calendarService, err := createCalendarService(client)

		if err != nil {
			log.Fatalf("Can not instantiate google calendar service: %v", err)
		}

		eventListFromFile := getEventListInEventFile()

		for _, event := range eventListFromFile {

			if event.date != "" {

				information := EventInformation{
					date:convertDateStringToInnerFormat(event.date),
					title:event.title,
					description:event.description,
					location: event.location,
				}

				insertEventToCalendar(calendarService, information)

			}
		}
	}
}

