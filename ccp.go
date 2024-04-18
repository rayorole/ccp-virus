package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
	"virus/ccp/lib"

	"github.com/reujab/wallpaper"
)

func openUrl(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func main() {
	lib.Setup()

	// Change desktop background
	wallpaper.SetFromURL("https://petapixel.com/assets/uploads/2013/06/tankman.jpg")

	wallpaper.SetMode(wallpaper.Crop)

	openUrl("./assets/flag.gif")
	openUrl("./assets/xi.jpg")

	done := make(chan bool)
	go lib.PlaySong(done) // Start playing the song in a new goroutine

	time.Sleep(10 * time.Second) // Wait for 10 seconds
	openUrl("https://s.weibo.com/weibo?q=%23%E8%BF%99%E6%89%8D%E6%98%AF%E7%9C%9F%E6%AD%A3%E7%9A%84%E4%B8%8B%E9%A5%AD%E7%95%8C%E6%89%9B%E6%8A%8A%E5%AD%90%23")

	time.Sleep(10 * time.Second)

	openUrl("https://www.bilibili.com/video/BV1ji421f7Xr/?spm_id_from=333.1007.tianma.1-3-3.click")

	// Clean up
	<-done
	fmt.Println("Ending...")
}
