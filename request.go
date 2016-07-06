package meli

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

//execute the request
func execute(method, path string, params map[string]string, body interface{}) (data []byte, err error) {
	var jsonStr []byte
	if body != nil {
		stJson, err := json.Marshal(body)

		if err != nil {
			log.Println(err)
		} else {
			jsonStr = stJson
		}
	}

	uri := makePath(path, params)

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(jsonStr))
	if err != nil {
		log.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", fmt.Sprintf("MELI-GO-SDK-%s", version))

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		var mlErr MLError
		json.Unmarshal(data, &mlErr)

		err = errors.New(mlErr.Message)
	}

	return
}

//makePath
func makePath(path string, params map[string]string) (uri string) {
	if strings.Contains(path, "http") == false {
		uri = fmt.Sprintf("%s/%s", apiRootUrl, path)
	} else {
		uri = path
	}

	if params != nil {
		query := url.Values{}

		for key, val := range params {
			query.Set(key, val)
		}

		uri += fmt.Sprintf("?%s", query.Encode())
	}

	return
}
