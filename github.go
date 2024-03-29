/*
Scraping scripts for github
*/
package scrapings

import (
	"fmt"
	"log"
	"time"

	"github.com/anaskhan96/soup"
)

// get latest release date
//
// param - userAndRep : string as shape "githubUserName/repositoryName"
// result - string as shape "2023-07-06T19:04:14Z"
func GetGitHubLatestReleaseDateByUserAndRep(userAndRep string) (result string, err error) {

	resp, err := soup.Get(fmt.Sprintf("https://github.com/%s/releases", userAndRep))
	//	resp, err := soup.Get("https://github.com/vuejs/vue/releases")
	if err != nil {
		log.Println(err)
		return
	}

	doc := soup.HTMLParse(resp)

	// soup.find("relative-time").attrs['datetime'] from BeautifulSoup
	result = doc.Find("relative-time").Attrs()["datetime"]
	return
}

func GetGitHubLatestReleaseDate(user string, repository string) (result string, err error) {
	userAndRep := fmt.Sprintf("%s/%s", user, repository)
	result, err = GetGitHubLatestReleaseDateByUserAndRep(userAndRep)

	return
}

func GetGitHubLatestReleaseDateUnixtimeByUserAndRep(userAndRep string) (result int64, err error) {
	timeStr, err := GetGitHubLatestReleaseDateByUserAndRep(userAndRep)
	if err != nil {
		return
	} else {
		var parsedTime time.Time
		if parsedTime, err = time.Parse("2006-01-02T15:04:05Z07:00", timeStr); err != nil {
			return
		} else {
			result = parsedTime.Unix()
			return
		}
	}
}

func GetGitHubLatestReleaseDateUnixtime(user string, repository string) (result int64, err error) {
	userAndRep := fmt.Sprintf("%s/%s", user, repository)
	result, err = GetGitHubLatestReleaseDateUnixtimeByUserAndRep(userAndRep)

	return
}
