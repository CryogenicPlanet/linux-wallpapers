package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/reujab/wallpaper"
)

func handleWallpaperChange(num int)  {
	fmt.Println("Number of active monitor(s) : ", num)


  	ex, err := os.Executable()
    if err != nil {
        panic(err)
    }
    path := filepath.Dir(ex)
    fmt.Println(path)

	plan, _ := ioutil.ReadFile(path + "/data.json")
	var data map[string]interface{}
	err = json.Unmarshal(plan, &data)

	if(err != nil) {
		fmt.Println("Error reading file")
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
