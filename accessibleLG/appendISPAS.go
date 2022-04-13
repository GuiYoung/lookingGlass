package accessibleLG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type tempResp struct {
	LgIsp string `json:"isp" gorm:"lg_isp"`
	LgAS  string `json:"as" gorm:"lg_AS"`
}

func appendISPAS(domainName string, status int, url string) {

	des_url := "http://ip-api.com/json/" + domainName + "?fields=isp,as,query"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, des_url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	lg := accessibleLG{}
	lg.LgStatus = status
	lg.LgUrl = url

	temp := tempResp{}

	err = json.Unmarshal(body, &temp)
	if err != nil {
		fmt.Println(err, domainName, string(body))
	}

	lg.LgIsp = temp.LgIsp

	re := regexp.MustCompile("AS[0-9]{1,5}")
	if matchStrs := re.FindStringSubmatch(temp.LgAS); len(matchStrs) > 0 {
		lg.LgAS, _ = strconv.Atoi(strings.Trim(matchStrs[0], "AS"))
	}

	if err := InsertLgUrl(&lg); err != nil {
		fmt.Println(err)
	}

}
