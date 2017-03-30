package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	usernamePwd := url.Values{}
	usernamePwd.Set("username", "pearl")
	usernamePwd.Set("password", "pearl")

	resp, err := http.PostForm("http://127.0.0.1:8081/api/get-token/", usernamePwd)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading body")
	}
	fmt.Println(string(body))
}
