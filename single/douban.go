package single

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

//定义新的数据类型
type Spider struct {
	url    string
	header map[string]string
}

//定义 Spider get的方法
func (keyword Spider) getHtmlHeader() string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", keyword.url, nil)
	if err != nil {
	}
	for key, value := range keyword.header {
		req.Header.Add(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return string(body)

}
func parse() {
	header := map[string]string{
		"Host":                      "movie.douban.com",
		"Connection":                "keep-alive",
		"Cache-Control":             "max-age=0",
		"Upgrade-Insecure-Requests": "1",
		"User-Agent":                "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.143 Safari/537.36",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8",
		"Referer":                   "https://movie.douban.com/top250",
	}

	//创建excel文件
	f, err := os.Create("./douban.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	//写入标题
	f.WriteString("电影名称" + "\t" + "\t" + "评分" + "\t" + "\t" + "评价人数" + "\t" + "\t" + "评语" + "\r\n")

	//循环每页解析并把结果写入excel
	for i := 0; i < 10; i++ {
		fmt.Println("正在抓取第" + strconv.Itoa(i) + "页......")
		url := "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25) + "&filter="
		spider := &Spider{url, header}
		html := spider.getHtmlHeader()

		//评价人数
		pattern2 := `<span>(.*?)评价</span>`
		rp2 := regexp.MustCompile(pattern2)
		findTxt2 := rp2.FindAllStringSubmatch(html, -1)

		//评分
		pattern3 := `property="v:average">(.*?)</span>`
		rp3 := regexp.MustCompile(pattern3)
		findTxt3 := rp3.FindAllStringSubmatch(html, -1)

		//电影名称
		pattern4 := `img width="100" alt="(.*?)" src=`
		rp4 := regexp.MustCompile(pattern4)
		findTxt4 := rp4.FindAllStringSubmatch(html, -1)

		pattern5 := `<span class="inq">(.*?)</span>`
		rp5 := regexp.MustCompile(pattern5)
		findTxt5 := rp5.FindAllStringSubmatch(html, -1)

		// 写入UTF-8 BOM
		f.WriteString("\xEF\xBB\xBF")
		//  打印全部数据和写入excel文件
		for i := 0; i < len(findTxt2); i++ {
			fmt.Printf("%s %s %s %s\n", findTxt4[i][1], findTxt3[i][1], findTxt2[i][1], findTxt5[i][1])
			f.WriteString(findTxt4[i][1] + "\t" + findTxt3[i][1] + "\t" + findTxt2[i][1] + "\t" + findTxt5[i][1] + "\r\n")

		}
	}
}

//func main() {
//	t1 := time.Now() // get current time
//	parse()
//	elapsed := time.Since(t1)
//	fmt.Println("爬虫结束,总共耗时: ", elapsed)
//
//}
