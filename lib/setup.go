package lib

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

func DownloadFile(filepath string, url string) (err error) {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func PlaySong(done chan bool) {
	f, err := os.Open("./assets/song.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	donePlaying := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		donePlaying <- true
	})))

	<-donePlaying // Wait for the song to finish
	done <- true  // Notify main function that song has finished
}

func Setup() {
	fmt.Println("Downloading dependencies...")

	// Make an assets directory, if it doesn't exist
	os.Mkdir("assets", 0755)

	DownloadFile("./assets/song.mp3", "https://drive.google.com/file/d/1VBiJPLfAUyCBKyCYuJ6b8ZdcL0A8K3me/view?usp=sharing")
	DownloadFile("./assets/flag.gif", "https://upload.wikimedia.org/wikipedia/commons/2/2c/Animated_China_Flag.gif")
	DownloadFile("./assets/xi.jpg", "https://i.kym-cdn.com/entries/icons/facebook/000/031/452/cover7.jpg")
}
