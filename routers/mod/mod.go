package mod

import (
	"dont/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"os/exec"
	"time"
	"unsafe"
)

type Vote struct {
	Num  string `form:"num" json:"num"`
	Star int    `form:"star" json:"num"`
}

type Searchmodinfo struct {
	Auth string `form:"auth" json:"auth"`
	Id   string `form:"id" json:"id"`
	Img  string `form:"img" json:"img"`
	Name string `form:"name" json:"name"`
	Sub  int    `form:"sub" json:"sub"`
	Time string `form:"time" json:"time"`
	Vote `form:"vote" json:"vote"`
}
type Moddown struct {
	//token string `form:"token" json:"token" uri:"token" xml:"token" binding:"required"`
	Modid string `form:"modid" json:"modid" uri:"modid" xml:"modid" binding:"required"`
}
type modResult struct {
	status  int    `json:"code"`
	modinfo string `json:"msg"`
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func SearchMod(g *gin.Context) {
	var searchmodinfo [30]Searchmodinfo
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("steamcommunity.com"),
	)
	c.Async = true
	//c.Limit(&colly.LimitRule{
	//	DomainGlob:  "*httpbin.*",
	//	Parallelism: 4,
	//	RandomDelay: 1 * time.Second, // 两次请求 随机延迟5s 内
	//})
	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		//r.Headers.Set("Host", "")
		//r.Headers.Set("Origin", "")
		//r.Headers.Set("Referer", "")
		//r.Headers.Set("Cookie", " sessionid=557ef7d9e971b4fde1efd407; timezoneOffset=28800,0; _ga=GA1.2.1389503378.1669915073; _gid=GA1.2.262271622.1670247728; recentlyVisitedAppHubs=219740%2C322330; steamCountry=HK%7C08d71887c56ddb684406707b272fe261; workshopNumPerPage=9; arp_scroll_position=100; app_impressions=322330@2_9_100013_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|219740@2_9_100013_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_9_100013_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_")
		r.Headers.Set("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8")
		// 這幾行在這iT邦幫忙沒有起到作用，但有些網站會照這些資訊判斷、阻擋其他來源
	})
	//if p, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1080", "http://127.0.0.1:1087"); err == nil {
	//	c.SetProxyFunc(p)
	//}
	q := 0
	c.OnHTML("div[class=workshopItem]", func(e *colly.HTMLElement) {

		searchmodinfo[q].Id = e.ChildAttr("a", "data-publishedfileid")
		searchmodinfo[q].Img = e.ChildAttr("a>div>img", "src")
		searchurl := e.ChildAttr("a[class=ugc]", "href")
		//fileRating := e.ChildAttr("img[class=fileRating]", "src")
		searchmodinfo[q].Name = e.ChildText("a[class=item_link]>div")
		//modauthor := e.ChildText("div>a[class=workshop_author_link]")
		//fmt.Println(modname)
		//fmt.Println(modauthor)
		//fmt.Println(modid)
		//fmt.Println(imgurl)
		//fmt.Println(searchurl)
		if searchurl != "" {
			c.Visit(searchurl)
		}
		//getmodextra(searchurl)
		//fmt.Println(fileRating)
		//link := e.Attr("href")
		//// Print link
		//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// Visit link found on page
		// Only those links are visited which are in AllowedDomains
		//c.Visit(e.Request.AbsoluteURL(link))
		q++

	})
	w := 0
	c.OnHTML("div[class=detailsStatsContainerRight]", func(e *colly.HTMLElement) {
		modsize := e.ChildText("div:nth-child(1)")
		modaddtime := e.ChildText("div:nth-child(2)")
		moduptime := e.ChildText("div:nth-child(3)")
		searchmodinfo[w].Time = moduptime
		fmt.Println(modsize)
		fmt.Println(modaddtime)
		fmt.Println(moduptime)
		w++

	})
	p := 0
	c.OnHTML("table[class=stats_table]", func(e *colly.HTMLElement) {
		modnoresub := e.ChildText("tbody>tr:nth-child(1)>td:nth-child(1)")
		modnowsub := e.ChildText("tbody>tr:nth-child(2)>td:nth-child(1)")
		modaddsub := e.ChildText("tbody>tr:nth-child(3)>td:nth-child(1)")
		searchmodinfo[p].Vote.Num = modnowsub
		p++
		fmt.Println(modnoresub)
		fmt.Println(modnowsub)
		fmt.Println(modaddsub)

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	// Start scraping on https://hackerspaces.org
	c.Visit("https://steamcommunity.com/workshop/browse/?appid=322330&searchtext=棱镜&browsesort=trend&section=&actualsort=trend&p=1&days=-1&numperpage=30")
	c.Wait()
	fmt.Println(searchmodinfo)
	g.JSON(http.StatusOK, searchmodinfo)
}
func AddMod(g *gin.Context) {
	//
	url := "https://workshop8.abcvg.info/archive/322330/1392778117.zip"
	fileName := "1392778117.zip"
	start := time.Now()
	d := models.NewDownloader(url, fileName, "", 5)
	if err := d.Download(); err != nil {
		log.Fatal(err)
	}
	log.Printf("下载%v 耗时：%v s", fileName, time.Since(start).Seconds())
}

func DownloadMod(g *gin.Context) {

	var form Moddown

	if g.Bind(&form) == nil {
		if form.Modid == "" {
		}
	}

	//modinfo := jsonmod(form.Modid)
	modinfo := DownloadMod2(form.Modid)
	//modinfo = strings.Replace(modinfo, "\n", "\\n", -1)
	//fmt.Printf(modinfo)
	data := "{\"status\": 1, \"modinfo\": " + modinfo + "}"
	fmt.Println(data)
	g.String(http.StatusOK, data)

}
func jsonmod(modid string) string {
	ml := "rm -rf /opt/dont-sh/lua-sh/modinfo.lua && cp /root/Steam/steamapps/workshop/content/322330/" + modid + "/modinfo.lua /opt/dont-sh/lua-sh/modinfo.lua"
	checkcmd := exec.Command("bash", "-c", ml)
	_, cherr := checkcmd.CombinedOutput()
	if cherr != nil {
		log.Fatalf("checkcmd.Run() failed with %s\n", cherr)
	}
	ml2 := "cd /opt/dont-sh/lua-sh/ && lua modgetinfo.lua"
	checkcmd2 := exec.Command("bash", "-c", ml2)
	check2, cherr2 := checkcmd2.CombinedOutput()
	//fmt.Println(string(check2))
	//fmt.Println(cherr2)
	if cherr2 != nil {
		log.Fatalf("checkcmd123123.Run() failed with %s\n", cherr2)
	}
	return string(check2)

}

func DownloadMod2(modid string) string {
	var isrun bool
	isrun = true
	var form Moddown

	checkcmd := exec.Command("bash", "-c", "tmux has-session -t DST_MODDOWN")
	cheout, cherr := checkcmd.CombinedOutput()
	fmt.Println("输出:", string(cheout), "错误信息", cherr)
	//fmt.Println(cherr == nil)
	if cherr != nil {
		isrun = false
	}
	fmt.Println(cheout)

	fmt.Print(isrun)
	if isrun == false {
		cmd := exec.Command("bash", "-c", "cd /root && tmux new-session -s DST_MODDOWN -d \"./steamcmd.sh\"")

		out, err := cmd.CombinedOutput()
		fmt.Printf("combined out:\n%s\n", string(out))
		if err != nil {
			log.Fatalf("cmd.Run() failed with %s\n", err)
		}
		cmd1 := exec.Command("bash", "-c", "tmux send-keys -t DST_MODDOWN 'login anonymous' C-m")
		out1, err1 := cmd1.CombinedOutput()
		fmt.Printf("combined out1:\n%s\n", string(out1))
		if err1 != nil {
			log.Fatalf("cmd1.Run() failed with %s\n", err1)
		}
	} else {
		fmt.Printf("已经存在tmux")
	}

	fmt.Println("Modid:", form.Modid)
	ml := "tmux send-keys -t DST_MODDOWN 'workshop_download_item 322330 " + modid + "' C-m"
	cmd2 := exec.Command("bash", "-c", ml)
	out2, err2 := cmd2.CombinedOutput()
	fmt.Printf("combined out2:\n%s\n", string(out2))
	if err2 != nil {
		log.Fatalf("cmd2.Run() failed with %s\n", err2)
	}
	for i := 0; i < 10; i++ {

	}

	ml2 := "rm -rf /opt/dont-sh/lua-sh/modinfo.lua && cp /root/Steam/steamapps/workshop/content/322330/" + modid + "/modinfo.lua /opt/dont-sh/lua-sh/modinfo.lua && cd /opt/dont-sh/lua-sh/ && lua modgetinfo.lua"
	checkcmd2 := exec.Command("bash", "-c", ml2)
	check2, cherr2 := checkcmd2.CombinedOutput()
	fmt.Println(string(check2))
	fmt.Println(cherr2)
	//if cherr2 != nil {
	//	log.Fatalf("checkcmd123123.Run() failed with %s\n", cherr2)
	//}
	return string(check2)
}
func getmodextra(url string) {
	c := colly.NewCollector(
		colly.AllowedDomains("steamcommunity.com"),
	)
	c.Async = true
	//c.Limit(&colly.LimitRule{
	//	DomainGlob:  "*httpbin.*",
	//	Parallelism: 4,
	//	RandomDelay: 1 * time.Second, // 两次请求 随机延迟5s 内
	//})
	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36")
		//r.Headers.Set("Host", "")
		//r.Headers.Set("Origin", "")
		//r.Headers.Set("Referer", "")
		//r.Headers.Set("Cookie", " sessionid=557ef7d9e971b4fde1efd407; timezoneOffset=28800,0; _ga=GA1.2.1389503378.1669915073; _gid=GA1.2.262271622.1670247728; recentlyVisitedAppHubs=219740%2C322330; steamCountry=HK%7C08d71887c56ddb684406707b272fe261; workshopNumPerPage=9; arp_scroll_position=100; app_impressions=322330@2_9_100013_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|219740@2_9_100013_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_9_100013_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_100101_100103|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_|322330@2_100100_230_")
		r.Headers.Set("Accept-Language", " zh-CN,zh;q=0.9,en;q=0.8")
		// 這幾行在這iT邦幫忙沒有起到作用，但有些網站會照這些資訊判斷、阻擋其他來源
	})
	//if p, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1080", "http://127.0.0.1:1087"); err == nil {
	//	c.SetProxyFunc(p)
	//}

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit(url)
	c.Wait()
	return
}
