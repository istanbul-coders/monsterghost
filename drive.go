package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"google.golang.org/api/drive/v2"
	"io"
	"log"
	"net/http"
	"os"
	"golang.org/x/oauth2"
<<<<<<< HEAD
=======
	"golang.org/x/net/context"
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
)

var eventFilePath = os.Getenv("HOME") + "/Event Schedule.xlsx"

const (
	MIN_NECESSARY_CELL_SIZE = 3
	EXPECTED_EXPORT_MIME_TYPE = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

func downloadFile(d *drive.Service, t http.RoundTripper, f *drive.File) {
	downloadUrl := f.DownloadUrl

	if downloadUrl == "" {
		fmt.Printf("An error occurred: File is not downloadable")
	}

	req, _ := http.NewRequest("GET", downloadUrl, nil)

	resp, _ := t.RoundTrip(req)
	defer resp.Body.Close()

	file, _ := os.Create(eventFilePath)
	defer file.Close()

	if _, err := io.Copy(file, resp.Body); err != nil {
		log.Fatal(err)
	}
}

func getEventListInEventFile() []EventInformation {
	eventListFromFile := make([]EventInformation, 20)
	xlFile, err := xlsx.OpenFile(eventFilePath)

	if err != nil {
		log.Fatalf("Can not read file: %v", err)
	}

	for i, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {

			if len(row.Cells) >= MIN_NECESSARY_CELL_SIZE {
				date, _ := row.Cells[0].String()
				title, _ := row.Cells[1].String()
				description, _ := row.Cells[2].String()

				if isConvenientEvent(i, date, description) {

					eventListFromFile = append(eventListFromFile, EventInformation{
						date:        date,
						title:       title,
						description: description,
						location:    EVENT_LOCATION,
					})

				}

			}

		}
	}

	return eventListFromFile
}

func isConvenientEvent(cellIndex int, date, description string) bool {
	return cellIndex == 0 && date != "" && date != "Date" && description != "";
}

func getDownloadUrlByName(driveService *drive.Service, name string) string {
	file, err := driveService.Files.List().Do()

	if err != nil {
		log.Fatalf("Unable to retrieve files.", err)
	}

	for _, f := range file.Items {
		if f.Title == name {
			exportLinks := f.ExportLinks
			return exportLinks[EXPECTED_EXPORT_MIME_TYPE]
		}
	}

	panic("File Not Found in Google Drive")

}

func createDriveService(client *http.Client) (*drive.Service, error) {
	return drive.New(client)
}

func createTransport() *oauth2.Transport {
	cacheFile, err := tokenCacheFile()

	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}

<<<<<<< HEAD
	token, err := tokenFromFile(cacheFile)

	return &oauth2.Transport{
		Source:oauth2.StaticTokenSource(token),
=======
	config := getConfig()
	token, err := tokenFromFile(cacheFile)

	tokenSource := config.TokenSource(context.Background(), token)
	return &oauth2.Transport{
		Source:tokenSource,
>>>>>>> 85cee804b66592fd6f3f52deb27ddef1da20d420
		Base:http.DefaultTransport,
	}
}

