package scrapings

import (
	"log"
	"testing"
	"time"

	"local.packages/scrapings"
)

func Test_01(t *testing.T) {
	if result, err := scrapings.GetGitHubLatestReleaseDateByUserAndRep("yt-dlp/yt-dlp"); err != nil {
		t.Errorf(err.Error())
	} else {
		log.Println(result)
		if parsedTime, err := time.Parse("2006-01-02T15:04:05Z07:00", result); err != nil {
			t.Errorf(err.Error())
		} else {
			log.Println(parsedTime)
			log.Println(parsedTime.Unix())
		}
	}
	if result, err := scrapings.GetGitHubLatestReleaseDate("yt-dlp", "yt-dlp"); err != nil {
		t.Errorf(err.Error())
	} else {
		log.Println(result)
	}
}
