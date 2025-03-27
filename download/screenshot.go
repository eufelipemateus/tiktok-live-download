// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package download

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"strings"
	"sync"
	"time"

	"github.com/eufelipemateus/go-video/browser"
	"github.com/eufelipemateus/go-video/interfaces"
	"github.com/eufelipemateus/go-video/process"
	"github.com/eufelipemateus/go-video/utils"
	"github.com/eufelipemateus/go-video/youtube"
	"github.com/steampoweredtaco/gotiktoklive"
	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)


var startTime = time.Now()

func GetVideo(browser *browser.Browser) {

	distFolder := "build"

	username := getUsernameFromURL(browser.Url)

	//filename := fmt.Sprintf("%s-%d-%02d-%02d-%02d%02d%02d", username, startTime.Year(), startTime.Month(), startTime.Day(), startTime.Hour(), startTime.Minute(), startTime.Second())
	filename := fmt.Sprintf("%s-%d-%02d-%02d", username, startTime.Year(), startTime.Month(), startTime.Day())

	outputFileVideo := fmt.Sprintf("%s/%s.mp4", distFolder, filename)
	outputFileChat := fmt.Sprintf("%s/%s.txt", distFolder, filename)
	outputFileViews := fmt.Sprintf("%s/%s.view", distFolder, filename)
	outputFinalFile := fmt.Sprintf("%s/%s-final.mp4", distFolder, filename)

	sigStage, err := browser.GetSigState()
	if err != nil {
		log.Printf("Erro ao obter estado do sinal: %v", err)
	}

	urlVideoRaw := getURLVideoRaw(sigStage)

	fmt.Printf("URL: %s \n %s\n", urlVideoRaw, outputFileChat)

	startTime = time.Now()
	var wg sync.WaitGroup
	wg.Add(2)

	go downloadVideo(urlVideoRaw, outputFileVideo, &wg, browser)

	os.Remove(outputFileChat)
	os.Remove(outputFileViews)
	go getChat(&wg, username, outputFileChat, outputFileViews)

	wg.Wait()

	os.Remove(outputFinalFile)
	process.GenerateVideoFinal(outputFileVideo, outputFinalFile, outputFileViews, outputFileChat)

	title := fmt.Sprintf("Live @%s - %s", username, startTime.Format("02/01/2006"))
	youtube.UploadVideo(outputFinalFile, title, "Live TikTok")

}
func getURLVideoRaw(data interfaces.SigiState) string {
	var parsedData interfaces.Root
	err := json.Unmarshal([]byte(data.LiveRoom.LiveRoomUserInfo.LiveRoom.StreamData.PullData.StreamData), &parsedData)
	if err != nil {
		log.Fatalf("Erro ao converter JSON: %v", err)
	}

	return parsedData.Data.Origin.Main.Flv
}

func downloadVideo(urlVideoRaw string, outputFile string, wg *sync.WaitGroup, browser *browser.Browser) {
	defer doneDownloadVideo(wg, browser)
	browser.Active = true
	err := ffmpeg_go.Input(urlVideoRaw).
		Output(outputFile, ffmpeg_go.KwArgs{"c": "copy"}).
		OverWriteOutput().
		Run()

	if err != nil {
		log.Printf("Erro ao baixar o vídeo: %v\n", err)
		return
	}
}

func doneDownloadVideo(wg *sync.WaitGroup, browser *browser.Browser) {
	defer wg.Done()
	log.Printf("Download concluído\n")
	browser.Active = false
}

func getChat(wg *sync.WaitGroup, username string, filePathChat string, filePathView string) {
	defer wg.Done()

	fileChat := openFile(filePathChat)
	fileView := openFile(filePathView)

	tiktok, err := gotiktoklive.NewTikTok()
	if err != nil {
		log.Printf("Erro ao criar o cliente TikTok: %s%v%s\n", utils.Colors.Red, err, utils.Colors.Reset)
		return
	}
	live, err := tiktok.TrackUser(username)
	if err != nil {
		log.Println(err)
	}

	// Receive livestream events through the live.Events channel
	for event := range live.Events {
		switch e := event.(type) {

		/*
			// You can specify what to do for specific events. All events are listed below.
			case gotiktoklive.UserEvent:
				fmt.Printf("%T : %s %s\n", e, e.Event, e.User.Username)

			// List viewer count


			// Specify the action for all remaining events
			default:
				fmt.Printf("%T : %+v\n", e, e)*/
		case gotiktoklive.ViewersEvent:
			_, err = fileView.WriteString(fmt.Sprintf("%.4f -> %d\n", time.Since(startTime).Seconds(), e.Viewers))
			if err != nil {
				log.Fatalf("Failed to append text: %v", err)
			}

		case gotiktoklive.ChatEvent:
			_, err = fileChat.WriteString(fmt.Sprintf("%.4f-%s -> %s\n", time.Since(startTime).Seconds(), e.User.Username, strings.Replace(e.Comment, "'", " ", -1)))

			if err != nil {
				log.Fatalf("Failed to append text: %v", err)
			}
		}
	}

	defer fileChat.Close()
	defer fileView.Close()

}

func getUsernameFromURL(url string) string {
	username := strings.TrimPrefix(url, "https://www.tiktok.com/@")
	username = strings.Split(username, "/")[0]
	return username
}

func openFile(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	return file
}
