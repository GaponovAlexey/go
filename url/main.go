package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

)

type UpdateRes struct {
	ok     bool     `json:"ok"`
	Result []Update `json"result"`
}
type Update struct {
	name string `json:"name"`
}

func main() {
	data, _ := test()
	// var res UpdateRes
	// _ = json.Unmarshal(da, res)
fmt.Println(data)

	// fmt.Println("da", res.Result)

}

func test() (data []byte, err error) {
	c := http.Client{}
	u := url.URL{
		Scheme: "https",
		Host:   "myHost",
		Path:   "path",
	}

	query := url.Values{}
	query.Add("offset", "1")
	query.Add("limit", "1")

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)

	req.URL.RawQuery = query.Encode()

	resp, _ := c.Do(req)

	defer func() { _ = resp.Body.Close() }()

	body, _ := io.ReadAll(resp.Body)
	return body, nil
}
