package spider

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GithubRun() Spider {
	githubSpider := Spider{
		Name: "GitHub Trending",
		Url:  "https://github.com/trending",
		// Url: "http://127.0.0.1:8000/trending",
		Score: 1,
		Out:   make([]OutItem, 0),
	}
	starRegexp := regexp.MustCompile("([0-9,]*)")
	ghDoc := GetDoc(githubSpider.Url)
	ghDoc.Find(".repo-list .col-12.d-block.width-full.py-4.border-bottom").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".d-inline-block.col-9.mb-1").Find("h3 a").Text()
		url, _ := s.Find(".d-inline-block.col-9.mb-1").Find("h3 a").Attr("href")
		desc := s.Find(".py-1").Text()
		stars := s.Find(".f6.text-gray.mt-2").Find(".d-inline-block.float-sm-right").Text()
		title = strings.Trim(title, " \n")
		desc = strings.Trim(desc, " \n")
		stars = starRegexp.FindString(strings.Trim(stars, " \n"))
		stars = strings.Replace(stars, ",", "", -1)
		// fmt.Printf("title:%s,desc:%s,star:%s\n", title, desc, stars)
		out := OutItem{
			Title:    title,
			Subtitle: desc,
			Stars:    stars,
			URL:      fmt.Sprintf("https://github.com%s", url),
		}
		githubSpider.Out = append(githubSpider.Out, out)
	})
	// toJson(githubSpider)
	mux.Lock()
	Out.Data = append(Out.Data, githubSpider)
	mux.Unlock()
	defer wait.Done()
	return githubSpider
}
