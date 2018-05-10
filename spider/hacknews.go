package spider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func HacknewsRun() Spider {
	hacknews := Spider{
		Url: "https://news.ycombinator.com/",
		// Url:  "http://127.0.0.1:8000/hackernews.html",
		Score: 2,
		Name:  "Hacker News",
		Out:   make([]OutItem, 0),
	}
	starRegexp := regexp.MustCompile("([0-9]*)")
	hnDoc := GetDoc(hacknews.Url)
	hnDoc.Find(".itemlist tbody tr").Each(func(i int, s *goquery.Selection) {
		if i%3 == 0 {
			// fmt.Println(i)
			link := s.Find("td.title a.storylink")
			title := link.Text()
			starNode := s.Next()
			url, _ := starNode.Find("a").Last().Attr("href")
			if !strings.HasPrefix(url, "http") {
				url = fmt.Sprintf("https://news.ycombinator.com/%s", url)
			}
			starText := starNode.Find(".score").Text()
			stars := starRegexp.FindString(strings.Trim(starText, " \n"))
			subtitle := strings.Trim(s.Find(".sitestr").Text(), " \n")
			// fmt.Printf("title:%s,link:%s,stars:%s\n", title, url, stars)
			hacknews.Out = append(hacknews.Out, OutItem{Title: title, URL: url, Subtitle: subtitle, Stars: stars})
		}
	})
	// toJson(hacknews)
	mux.Lock()
	Out.Data = append(Out.Data, hacknews)
	mux.Unlock()
	defer wait.Done()
	return hacknews
}
