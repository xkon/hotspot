package spider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Spider struct {
	Name  string
	Url   string
	Score int // 重要程度，1为最高
	Out   []OutItem
}
type OutItem struct {
	Title    string
	Subtitle string
	Stars    string
	URL      string
}

type OutPut struct {
	Date string
	Data []Spider
}

var wait sync.WaitGroup
var mux sync.Mutex
var Out OutPut = OutPut{
	Date: time.Now().Format("2006-01-02 15:04:05"),
	Data: make([]Spider, 0),
}

// Stars逆序排序
type OutSlice []OutItem

func (o OutSlice) Len() int {
	return len(o)
}

func (o OutSlice) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o OutSlice) Less(i, j int) bool {
	v1, _ := strconv.Atoi(o[i].Stars)
	v2, _ := strconv.Atoi(o[j].Stars)
	return v1 > v2
}

// spider score 排序
type SpiderSlice []Spider

func (s SpiderSlice) Len() int {
	return len(s)
}

func (s SpiderSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SpiderSlice) Less(i, j int) bool {
	return s[i].Score < s[j].Score
}

func GetDoc(url string) *goquery.Document {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
	}
	return doc
}

func toJson(s Spider) []byte {
	b, err := json.Marshal(s)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func Run() {
	wait.Add(5)
	go GithubRun()
	go HacknewsRun()
	go V2exRun()
	go XZRun()
	go FBRun()
	wait.Wait()
	b, err := json.Marshal(Out)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q", b)
	writeToFile(b)
}

func writeToFile(b []byte) error {
	filename := fmt.Sprintf("data/%s.json", time.Now().Format("2006010215"))
	if err := ioutil.WriteFile(filename, b, 0644); err != nil {
		log.Print(err)
		return err
	} else {
		log.Printf("\n[*] done ")
		return nil
	}
}
