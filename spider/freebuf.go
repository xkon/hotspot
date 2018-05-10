package spider

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FBRun() Spider {
	fb := Spider{
		Name:  "FreeBuf",
		Url:   "http://www.freebuf.com/tech?pg=",
		Score: 5,
		Out:   make([]OutItem, 0),
	}
	extract := func(i int, s *goquery.Selection) {

		titlelink := s.Find("dt a")
		title := strings.Trim(titlelink.Text(), " \n")
		url, _ := titlelink.Attr("href")

		stars := s.Find("span.look strong").First().Text()
		// stars := strconv.Itoa(20 - i)
		desc := strings.Trim(s.Find("dd.text").Text(), " \n")

		// fmt.Printf("title:%s,desc:%s,star:%s\n", title, desc, stars)
		out := OutItem{
			Title:    title,
			Subtitle: desc,
			Stars:    stars,
			URL:      url,
		}
		fb.Out = append(fb.Out, out)
	}
	fbDoc1 := GetDoc(fb.Url + "1")
	fbDoc1.Find(".news_inner.news-list").Each(extract)
	fbDoc2 := GetDoc(fb.Url + "2")
	fbDoc2.Find(".news_inner.news-list").Each(extract)

	// toJson(v2ex)
	mux.Lock()
	Out.Data = append(Out.Data, fb)
	mux.Unlock()
	defer wait.Done()
	return fb
}
