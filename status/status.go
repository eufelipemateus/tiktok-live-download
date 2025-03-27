package status

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eufelipemateus/go-video/browser"
	"github.com/eufelipemateus/go-video/interfaces"
)

func StartLive(liveURLs []string) []*browser.Browser {
	livesCtxlist := []*browser.Browser{}

	for _, url := range liveURLs {
		browser := browser.NewBrowser(url)
		livesCtxlist = append(livesCtxlist, browser)
		time.Sleep(1 * time.Second)
	}

	time.Sleep(5 * time.Second)

	return livesCtxlist
}

func IsLiveOffline(metada interfaces.SigiState) bool {
	status := metada.LiveRoom.LiveRoomUserInfo.LiveRoom.Status
	return (status == 4 || status == 3)
}

func LoadURLsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "//") {
			urls = append(urls, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return urls
}
