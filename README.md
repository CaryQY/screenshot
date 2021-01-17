- code reference [https://github.com/vova616/screenshot](https://github.com/vova616/screenshot) 
- code reference [https://github.com/vova616/screenshot](https://github.com/vova616/screenshot) 
- code reference [https://github.com/vova616/screenshot](https://github.com/vova616/screenshot) 
- for linux platform
- handle  [Bad things happenedx protocol authentication refused: Maximum number of clients reached](https://github.com/BurntSushi/xgb/issues/40) error
- add GetScreen() function, dynamic image quality and encode image to jpg

## Basic Usage

> more usage please check  "screenshot_linux_test.go" file

```
package main

import (
	"fmt"
	"github.com/CaryQY/screenshot"
	"image/png"
	"os"
	"time"
)

func main() {
	shot, err := screenshot.New()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer shot.Close()

	// custom save img to png
	{
		img, err := shot.CaptureScreen()
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Create("./" + time.Now().Format("200601021504") + ".png")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		err = png.Encode(f, img)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// save img to jpg
	{
		img, err := shot.GetScreen(75)
		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Create("./" + time.Now().Format("200601021504") + ".jpg")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, err = f.Write(img)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
```
