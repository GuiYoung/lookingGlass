package accessibleLG

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func readSRCURL() (results []string) {
	file, _ := os.Open("lgs.txt")

	defer file.Close()

	// 接受io.Reader类型参数 返回一个bufio.Scanner实例
	scanner := bufio.NewScanner(file)

	var count int

	for scanner.Scan() {
		count++

		// 读取当前行内容
		line := scanner.Text()
		results = append(results, line)
	}
	return results
}

func randomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func checkAccessibleLG(url string) (int, error) {

	method := "GET"
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return 500, err
	}

	userAgent := randomString()
	req.Header.Add("User-Agent", userAgent)
	res, err := client.Do(req)
	if err != nil {
		return 500, err
	}
	return res.StatusCode, err
}

func getLG(body []byte) {

	buf := new(bytes.Buffer)
	_, err := buf.Write(body)
	if err != nil {
		panic(err)
	}

	var lgs []string

	re := regexp.MustCompile("\"looking_glass\": \".*?\",")
	if matchStrs := re.FindAllStringSubmatch(buf.String(), -1); len(matchStrs) > 0 {
		for _, v := range matchStrs {
			lgs = append(lgs, v[0][18:len(v[0])-2])
		}
	}

	f, err := os.Create("lgs.txt")
	defer f.Close()
	for _, url := range lgs {
		if url != "" {
			io.WriteString(f, url+"\n")
		}

	}

}

func getAllInfo() []byte {
	url := "https://www.peeringdb.com/api/net"
	method := "GET"
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36 Edg/100.0.1185.29")

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("resp.txt")
	if err != nil {
		panic(err)
	}

	io.Copy(f, bytes.NewReader(body))

	return body
}

func GetAccessibleURL() {
	if err := InitDb(); err != nil {
		panic(err)
	}
	//getLG(getAllInfo())

	for _, url := range readSRCURL() {
		var status int
		var err error
		if status, err = checkAccessibleLG(url); err != nil {
			continue
		}

		t := strings.Split(url, "//")[1]
		domainName := strings.Split(t, "/")[0]

		appendISPAS(domainName, status, url)
	}
}
