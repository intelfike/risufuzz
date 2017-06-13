package fuzz

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

func Fuzz(data []byte) int {
	c := http.Client{}
	c.Jar, _ = cookiejar.New(nil)

	_, err := c.Get("http://aji.risumail.net/beta1/src/login.php")
	if err != nil {
		return 0
	}
	resp, err := c.PostForm("http://aji.risumail.net/beta1/src/redirect.php",
		url.Values{
			"login_username":        {string(data)},
			"secretkey":             {string(data)},
			"js_autodetect_results": {string(data)},
			"just_logged_in":        {string(data)},
		},
	)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	defer resp.Body.Close()
	b := new(bytes.Buffer)
	io.Copy(b, resp.Body)
	if resp.StatusCode != 200 {
		fmt.Println(resp.StatusCode)
		fmt.Println(b)
		return 0
	}
	if !strings.Contains(b.String(), "login.php") {
		fmt.Println(b)
		return 0
	}
	return 1
}
