package rssFetcher

import (
	"encoding/xml"
	"log"
	"net/http"
	"time"
)

type RSS struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Items []Item `xml:"item"`
}

type Item struct {
	Title   string `xml:"title"`
	Link    string `xml:"link"`
	PubDate string `xml:"pubDate"`
}

func FetchFacebookNews() ([]Item, error) {
	faebookRssURL := "https://about.fb.com/news/feed/" // facebook news feed page url
	resp, err := http.Get(faebookRssURL)
	if err != nil {
		log.Printf("Error while fetching content from facebook rssUrl, Err : %+v\n", err)
		return nil, err
	}
	defer resp.Body.Close()

	var rss RSS
	if err := xml.NewDecoder(resp.Body).Decode(&rss); err != nil {
		log.Printf("Error while decoding response from facebook rss feed, Err : %+v\n", err)
		return nil, err
	}

	// Format the dates in the fetched news items
	for i, item := range rss.Channel.Items {
		rss.Channel.Items[i].PubDate = FormatDate(item.PubDate)
	}

	return rss.Channel.Items, nil
}

func FormatDate(rssDate string) string {
	parsedTime, err := time.Parse(time.RFC1123, rssDate)
	if err != nil {
		log.Printf("Error while parsing article time in yyyy-mm-dd format, Err : %+v\n", err)
		return rssDate // Return original if parsing fails
	}
	return parsedTime.Format("2006-01-02 15:04")
}
