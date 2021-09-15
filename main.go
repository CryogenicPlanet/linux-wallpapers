package main

import (
	"fmt"
	"time"
    "encoding/json"
	"io/ioutil"
	"strconv"
	"os"
	"github.com/kbinani/screenshot"
	"github.com/reujab/wallpaper"
)

func handleWallpaperChange(num int)  {
	fmt.Println("Number of active monitor(s) : ", num)

	plan, _ := ioutil.ReadFile("data.json")
	var data map[string]interface{}
	err := json.Unmarshal(plan, &data)

	if(err != nil) {
		fmt.Println("Error reading file")
	}

	path, err := os.Getwd()
	if err != nil {
	    fmt.Println(err)
	}

	fmt.Println("JSON data", data)

	key := strconv.Itoa(num)

	wallpaperToSet := data[key].(string)

	
	wallpaperPath := path + "/wallpapers/" +  wallpaperToSet
	
	fmt.Println("Wallpaper to set", wallpaperPath)
	// Can assign wallpapers based on dimensions later
	// for i := 0; i < num; i++ {
	// 	bounds := screenshot.GetDisplayBounds(i)

	// 	x := bounds.Dx()
	// 	y := bounds.Dy()

	// 	fmt.Printf("Display #%d resolution is %d x %d\n", i, x, y)
	// }

	err = wallpaper.SetFromFile(wallpaperPath)	
	if err != nil {
		fmt.Println("Error getting wallpaper", err)
	}
	fmt.Println("Wallpaper Updated")
}

func main() {
	ch := make(chan int)
	exit:= make(chan bool)
	go func() {
		for {
			numDisplays := screenshot.NumActiveDisplays()
			ch <- numDisplays
			time.Sleep(5 * time.Second)
		}
		exit<-true // Notify main() that this goroutine has finished
	}()
	go func() {
		lastNum := 0
		for {
			select {
			case num := <-ch:
				if(num != lastNum) {
					handleWallpaperChange(num)
				}
				lastNum = num
				break
			}
		}
	}()
  	<-exit // This blocks until the exit channel receives some input
    fmt.Println("Done.")
}
