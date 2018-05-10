package spider

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func XZRun() Spider {
	xz := Spider{
		Name:  "先知社区",
		Url:   "https://xz.aliyun.com/",
		Score: 4,
		Out:   make([]OutItem, 0),
	}
	xzDoc := GetDoc(xz.Url)
	xzDoc.Find(".table.topic-list tr").Each(func(i int, s *goquery.Selection) {
		if i > 20 {
			return
		}
		titlelink := s.Find("a.topic-title")
		title := strings.Trim(titlelink.Text(), " \n")
		url, _ := titlelink.Attr("href")

		// stars := s.Find(".badge.badge-hollow.text-center").Text()
		stars := strconv.Itoa(20 - i)
		desc := strings.Trim(s.Find(".topic-info").Text(), " \n")
		descSlice := strings.Split(desc, "\n")
		desc = strings.Join(descSlice[:2], " ")

		// fmt.Printf("title:%s,desc:%s,star:%s\n", title, desc, stars)
		out := OutItem{
			Title:    title,
			Subtitle: desc,
			Stars:    stars,
			URL:      fmt.Sprintf("https://xz.aliyun.com%s", url),
		}
		xz.Out = append(xz.Out, out)
	})
	// toJson(v2ex)
	mux.Lock()
	Out.Data = append(Out.Data, xz)
	mux.Unlock()
	defer wait.Done()
	return xz
}
