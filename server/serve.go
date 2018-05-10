package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/xkon/hotspot/spider"
)

func ListenAndServe() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(cors.Default())

	router.LoadHTMLGlob("frontend/dist/*.html")
	router.GET("/hotspots", renderIndex)
	router.GET("/api/hotspots", getHotSpots)
	router.StaticFS("/static/", http.Dir("frontend/dist/static/"))
	router.StaticFile("/service-worker.js", "frontend/dist/service-worker.js")

	router.Run("127.0.0.1:9999")
}

func renderIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}

func getHotSpots(c *gin.Context) {
	filename := fmt.Sprintf("data/%s.json", time.Now().Format("2006010215"))
	if _, err := os.Stat(filename); err != nil {
		oneHour, _ := time.ParseDuration("-1h")
		filename = fmt.Sprintf("data/%s.json", time.Now().Add(oneHour).Format("2006010215"))
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		// go spider.Run()
		log.Print(err)
		c.JSON(http.StatusOK, gin.H{"code": 404, "data": "data not found"})
		return
	}
	var spiderOut spider.OutPut
	juerr := json.Unmarshal(content, &spiderOut)
	if juerr != nil {
		log.Print(juerr)
	}
	sort.Sort(spider.SpiderSlice(spiderOut.Data))
	for i := 0; i < len(spiderOut.Data); i++ {
		sort.Sort(spider.OutSlice(spiderOut.Data[i].Out))
		// 截取 前20
		if len(spiderOut.Data[i].Out) > 20 {
			spiderOut.Data[i].Out = spiderOut.Data[i].Out[:20]
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": spiderOut})
}
