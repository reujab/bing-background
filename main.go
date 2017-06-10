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

	if err != nil {
		panic(err)
	}

	defer func() {
		err = res.Body.Close()

		if err != nil {
			panic(err)
		}
	}()

	images := new(images)

	err = json.NewDecoder(res.Body).Decode(images)

	if err != nil {
		panic(err)
	}

	err = wallpaper.SetFromURL("https://www.bing.com" + images.Images[0].URL)

	if err != nil {
		panic(err)
	}
}
