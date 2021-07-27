package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	data := struct {
		Revision        string                 `json:"revision"`
		Parallel        int64                  `json:"parallel"`
		BuildParameters map[string]interface{} `json:"build_parameters"`
	}{
		Revision: "",
		BuildParameters: map[string]interface{}{
			"RUN_EXTRA_TESTS": "true",
		},
	}
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal("failed to marshal:", err)
	}
	body := bytes.NewReader(b)
	url := fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s/tree/%s", "github", "hnnsgstfssn", "gerrittest", "master")
	req, _ := http.NewRequest(http.MethodPost, url, body)
	req.Header.Set("Circle-Token", "8a28e58f6a257e650e780b6ad1eb70052afb84ca")
	req.Header.Set("Content-Type", "application/json")
	b, err = httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal("dump:", err)
	}
	fmt.Printf("%s", b, err)
	res, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal("failed roundtrip:", err)
	}
	b, err = httputil.DumpResponse(res, true)
	if err != nil {
		log.Fatal("dump:", err)
	}
	fmt.Printf("%s", b, err)
}
