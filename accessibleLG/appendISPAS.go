package accessibleLG

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func appendISPAS(domainName string, status int, url string) {

	des_url := "http://ip-api.com/json/" + domainName + "?fields=status,message,isp,as,query"
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

	err = json.Unmarshal(body, &lg)
	if err != nil {
		fmt.Println(err, domainName, string(body))
	}

	if err := InsertLgUrl(&lg); err != nil {
		fmt.Println(err)
	}

}
