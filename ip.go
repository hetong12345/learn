package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	URL   = "http://api.ip138.com/query/"
	TOKEN = "532950476e8471adfce7b4aafc52f07c"
	//TOKEN = "bd4c2bf9a38ab06f7cae88c9759ee172"
)

//----------------------------------
// iP地址调用示例代码
//----------------------------------

// xml struct
type xmlinfo struct {
	Ret  string          `xml:"ret"`
	Ip   string          `xml:"ip"`
	Data locationxmlInfo `xml:"data"`
}

type locationxmlInfo struct {
	Country string `xml:"country"`
	Region  string `xml:"region"`
	City    string `xml:"city"`
	Isp     string `xml:"isp"`
	Zip     string `xml:"zip"`
	Zone    string `xml:"zone"`
}

//json struct
type jsoninfo struct {
	Ret  string    `json:"ret"`
	Ip   string    `json:"ip"`
	Data [6]string `json:"data"`
}

//func main() {
//	ipLocation("43.224.45.105", "json")
//}

func ipLocation(ip string, dataType string) {

	queryUrl := fmt.Sprintf("%s?ip=%s&datatype=%s", URL, ip, dataType)
	fmt.Println(queryUrl)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", queryUrl, nil)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	reqest.Header.Add("token", TOKEN)
	response, err := client.Do(reqest)
	defer response.Body.Close()

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	if response.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(bodyByte))
		if dataType == "json" {
			var info jsoninfo
			json.Unmarshal(bodyByte, &info)
			fmt.Println(info)
		} else if dataType == "xml" {
			var info xmlinfo
			xml.Unmarshal(bodyByte, &info)
			fmt.Println(info.Ip)
		}
	}

	return
}
