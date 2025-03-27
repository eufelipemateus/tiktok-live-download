package main

import (
	"log"
	"sync"
	"time"

	"github.com/eufelipemateus/go-video/browser"
	"github.com/eufelipemateus/go-video/download"
	"github.com/eufelipemateus/go-video/status"
	"github.com/eufelipemateus/go-video/utils"
)

func Run() {
	liveURLs := status.LoadURLsFromFile("lives.txt")
	livesCtxlist := status.StartLive(liveURLs)
	for {
		var wg sync.WaitGroup
		wg.Add(len(liveURLs))

		for _, b := range livesCtxlist {
			go func(browser *browser.Browser) {
				defer wg.Done()

				if browser.Active {
					return
				}

				err := browser.Reload()
				if err != nil {
					log.Printf("Erro ao recarregar a página em %s '%s' %s: %s %v %s \n", utils.Colors.Blue, browser.Url, utils.Colors.Reset, utils.Colors.Red, err, utils.Colors.Reset)
					return
				}

				sigState, err := browser.GetSigState()
				if err != nil {
					log.Printf("Erro ao obter estado do sinal em %s '%s' %s: %s %v %s \n", utils.Colors.Blue, browser.Url, utils.Colors.Reset, utils.Colors.Red, err, utils.Colors.Reset)
					return
				}
				isOffline := status.IsLiveOffline(sigState)


				if isOffline {
					log.Printf("A live em '%s%s%s' está %sOFFLINE%s (LIVE encerrada)\n", utils.Colors.Blue, browser.Url, utils.Colors.Reset, utils.Colors.Red, utils.Colors.Reset)
				} else {
					log.Printf("A live em '%s%s%s' ainda está %sONLINE%s\n", utils.Colors.Blue, browser.Url, utils.Colors.Reset, utils.Colors.Green, utils.Colors.Reset)
					go download.GetVideo(browser)
				}
			}(b)
		}
		wg.Wait()
		time.Sleep(1 * time.Minute)
	}
}
