package main

import "encoding/json"
import "fmt"
import "github.com/reujab/wallpaper"
import "io/ioutil"
import "net/http"
import "os"

func main() {
	res, err := http.Get("https://www.bing.com/HPImageArchive.aspx?format=js&n=1")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := res.Body.Close(); err != nil {
			panic(err)
		}
	}()

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{}

	if err := json.Unmarshal(bytes, &data); err != nil {
		panic(err)
	}

	imageURL := "https://www.bing.com" + data["images"].([]interface{})[0].(map[string]interface{})["url"].(string)

	if err := wallpaper.SetFromURL(imageURL); err != nil {
		if err == wallpaper.ErrUnsupportedDE {
			desktop := wallpaper.Desktop

			if desktop == "" {
				desktop = "Your desktop environment"
			}

			fmt.Fprintf(os.Stderr, "%s is not supported.\n", desktop)
			fmt.Fprintf(os.Stderr, "However, you can download %s and set it as your wallpaper manually.\n", imageURL)
			os.Exit(1)
		} else {
			panic(err)
		}
	}
}
