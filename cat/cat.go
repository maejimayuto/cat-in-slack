package cat

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Cat struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func GetCats() []Cat {
	body, err := httpGetStr("https://api.thecatapi.com/v1/images/search?api_key=3e315e96-8143-4a6c-b32a-47252e6cc242")
	if err != nil {
		return nil
	}
	cats, err := formatCat(body)
	if err != nil {
		return nil
	}
	return cats
}

func httpGetStr(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
		return nil, err
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get this url : http status %d", response.StatusCode)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
		return nil, err
	}
	return body, nil
}

func formatCat(body []byte) ([]Cat, error) {
	var cats []Cat
	if err := json.Unmarshal(body, &cats); err != nil {
		return nil, err
	}

	return cats, nil
}

func (c *Cat) ToString() string {
	url := fmt.Sprintf("url: %s\n", c.Url)
	width := fmt.Sprintf("width: %d\n", c.Width)
	height := fmt.Sprintf("height: %d\n", c.Height)
	result := url + width + height

	return result
}
