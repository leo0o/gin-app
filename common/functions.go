package common

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

func HttpPostForm(url string, params url.Values) (v interface{}) {
	response, err := http.PostForm(url, params)
	if err != nil {
		return nil
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(body, &v)
	return v
}
