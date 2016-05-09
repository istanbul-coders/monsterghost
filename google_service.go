package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"io/ioutil"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/calendar/v3"
)

var clientIdFile = environmentName + "/client_id.json"

func createGoogleApiClient() *http.Client {
<<<<<<< HEAD
	ctx := context.Background()
	b, err := ioutil.ReadFile(clientIdFile)
=======
	return getClient(context.Background(), getConfig())
}

func getConfig() *oauth2.Config {
	clientIdFile, err := ioutil.ReadFile(clientIdFile)
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420

	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

<<<<<<< HEAD
	config, err := google.ConfigFromJSON(b, drive.DriveScope, calendar.CalendarScope)
=======
	config, err := google.ConfigFromJSON(clientIdFile, drive.DriveScope, calendar.CalendarScope)
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420

	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

<<<<<<< HEAD
	return getClient(ctx, config)
=======
	return config
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
}

func getClient(ctx context.Context, config *oauth2.Config) *http.Client {
	cacheFile, err := tokenCacheFile()

	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}

	tok, err := tokenFromFile(cacheFile)
<<<<<<< HEAD
=======

>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(cacheFile, tok)
	}

	return config.Client(ctx, tok)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	fmt.Printf("Go to the following link in your browser then type the " +
	"authorization code: \n%v\n", authURL)

	var code string

	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	tok, err := config.Exchange(oauth2.NoContext, code)

	if err != nil {
		log.Fatalf("Unable to retrieve token from web %v", err)
	}

	return tok
}

func tokenCacheFile() (string, error) {
	usr, err := user.Current()

	if err != nil {
		return "", err
	}

	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)

	return filepath.Join(tokenCacheDir,
		url.QueryEscape("drive-go-quickstart.json")), err
}

func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)

	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()

	return t, err
}

func saveToken(file string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.Create(file)

	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}

	defer f.Close()
	json.NewEncoder(f).Encode(token)
}