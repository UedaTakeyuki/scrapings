package scrapings

import (
	"log"
	"testing"

	"local.packages/scrapings"
)

func Test_01(t *testing.T) {
	if result, err := scrapings.GetGitHubLatestReleaseDateByUserAndRep("yt-dlp/yt-dlp"); err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
	if result, err := scrapings.GetGitHubLatestReleaseDate("yt-dlp", "yt-dlp"); err != nil {
		log.Println(err)
	} else {
		log.Println(result)
	}
}
