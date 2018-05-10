package spider

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func V2exRun() Spider {
	v2ex := Spider{
		Name:  "v2ex Hot",
		Url:   "https://www.v2ex.com/?tab=hot",
		Score: 3,
		Out:   make([]OutItem, 0),
	}
	v2Doc := GetDoc(v2ex.Url)
	v2Doc.Find(".cell.item").Each(func(i int, s *goquery.Selection) {
		titlelink := s.Find(".item_title a")
		title := strings.Trim(titlelink.Text(), " \n")
		url, _ := titlelink.Attr("href")

		stars := s.Find(".count_livid").Text()
		desc := strings.Trim(s.Find(".topic_info .node").Text(), " \n")

		// fmt.Printf("title:%s,desc:%s,star:%s\n", title, desc, stars)
		out := OutItem{
			Title:    title,
			Subtitle: desc,
			Stars:    stars,
			URL:      fmt.Sprintf("https://www.v2ex.com%s", url),
		}
		v2ex.Out = append(v2ex.Out, out)
	})
	// toJson(v2ex)
	mux.Lock()
	Out.Data = append(Out.Data, v2ex)
	mux.Unlock()
	defer wait.Done()
	return v2ex
}
