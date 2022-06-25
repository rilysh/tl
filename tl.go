package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func req_api(to string, query string) string {
	return "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto" + "&tl=" + to + "&dt=t&q=" + url.QueryEscape(query)
}

func prt(text string) {
	fmt.Print(text)
}

func help() {
	prt(os.Args[0] + " v0.1\n\nUsage: \n" + os.Args[0] + " [options/none] [message]\n\nOptions:\n -h, shows this message\n -l, language code translated to\n -f, file to translate\n")
}

func main() {
	var empty []interface{}
	var to string
	var query []string
	var query_str string
	var client http.Client
	var request *http.Response

	if len(os.Args) < 2 {
		help()
		return
	}

	if os.Args[1] == "-h" {
		help()
		return
	}

	if os.Args[1] != "-l" {
		to = "en"
		if os.Args[1] == "-f" {
			if len(os.Args) < 3 {
				help()
				return
			}
			file, err := ioutil.ReadFile(os.Args[2])
			if err != nil {
				prt("Something went wrong while reading " + os.Args[2] + " Debug: " + err.Error())
				return
			}
			query_str = string(file)
			request, _ = client.Get(req_api(to, query_str))
		} else {
			query = os.Args[1:]
			request, _ = client.Get(req_api(to, strings.Join(query, " ")))
		}
	} else {
		if len(os.Args) < 4 {
			help()
			return
		}
		to = os.Args[2]
		if os.Args[3] == "-f" {
			if len(os.Args) < 5 {
				help()
				return
			}
			file, err := ioutil.ReadFile(os.Args[4])
			if err != nil {
				prt("Something went wrong while reading " + os.Args[2] + " Debug: " + err.Error())
				return
			}
			query_str = string(file)
			request, _ = client.Get(req_api(to, query_str))
		} else {
			query = os.Args[3:]
			request, _ = client.Get(req_api(to, strings.Join(query, " ")))
		}
	}

	client = http.Client{}

	switch request.StatusCode {
	case 404:
		prt("Unable to find any results, make sure you've picked up correct language code.\n")
		break

	case 301:
		prt("Unable to proceed furthur, permanently moved.\n")
		break

	case 429:
		prt("You're being rate limited, try to request with a vpn or proxy to bypass it.\n")
		break
	}

	buffer, err := ioutil.ReadAll(request.Body)
	err = json.Unmarshal([]byte(string(buffer)), &empty)
	if err != nil {
		prt("Failed to parse JSON. Debug: " + err.Error() + "\n")
		return
	}
	defer request.Body.Close()
	var body []string
	first_arr := empty[0]

	for _, second_arr := range first_arr.([]interface{}) {
		for _, reached_tl := range second_arr.([]interface{}) {
			body = append(body, fmt.Sprintf("%s", reached_tl))
		}
	}
	prt(body[0] + "\n")
}
