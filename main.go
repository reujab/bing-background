package main

import (
	"encoding/json"
	"net/http"

	"github.com/reujab/wallpaper"
)

type images struct {
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
}

func main() {
	res, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&n=1")

	check(err)

	defer func() {
		check(res.Body.Close())
	}()

	images := new(images)

	check(json.NewDecoder(res.Body).Decode(images))
	check(wallpaper.SetFromURL("https://www.bing.com" + images.Images[0].URL))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
