package external

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type DataAPI interface {
	Search() (string, error)
}

type dataAPI struct {
	unit int
}

func (receiver *dataAPI) Search() (string, error) {
	params := url.Values{}
	params.Add("part", "snippet")
	params.Add("key", "<your_api_key>")
	params.Add("q", "뷰티")
	params.Add("type", "channel")
	response, err := http.Get("https://www.googleapis.com/youtube/v3/search?" + params.Encode())

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	data := string(responseData)
	fmt.Println(data)

	return data, nil
}

func NewDataAPI() DataAPI {
	return &dataAPI{
		unit: 10000,
	}
}
