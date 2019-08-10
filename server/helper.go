package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getRequestData(r *http.Request) map[string]interface{} {
	body := ""
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		body = string(buf)
	}

	// get the form
	r.ParseForm()
	form := r.PostForm

	data := map[string]interface{}{
		"url":     r.URL.Path,
		"method":  r.Method,
		"headers": r.Header,
		"query":   r.URL.Query(),
		"form":    form,
		"body":    body,
	}
	return data

}

func GetDict(r *http.Request, extras map[string]string, keys ...string) map[string]interface{} {

	data := getRequestData(r)

	result := map[string]interface{}{}

	for _, key := range keys {
		if _, ok := data[key]; ok {
			result[key] = data[key]
		}
	}

	if extras != nil && len(extras) > 0 {
		for k, v := range extras {
			result[k] = v
		}
	}

	return result
}
