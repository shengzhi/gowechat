package util

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Invoke http.get to fetch data from specified url
func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)

}

func PostJSON(url string, request, response interface{}) error {
	data, err := json.Marshal(request)
	if err != nil {
		return err
	}
	buf := TextBufferPool.Get().(*bytes.Buffer)
	buf.Reset()
	if _, err = buf.Read(data); err != nil {
		return err
	}
	defer TextBufferPool.Put(buf)
	res, err := http.Post(url, "application/json; charset=utf-8", buf)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if data, err = ioutil.ReadAll(res.Body); err != nil {
		return err
	}
	return json.Unmarshal(data, response)
}
