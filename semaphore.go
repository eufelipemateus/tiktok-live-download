package main

import (
	"log"
	"time"

	"github.com/eufelipemateus/go-video/download"
	"github.com/eufelipemateus/go-video/status"
)

func Run(liveURLs []string) {

	onlineLives := []string{}

	// Verifica o status de cada live
	// Iterando pelas URLs e verificando o status
	for {
		for _, url := range liveURLs {
			isOffline, err := status.IsLiveOffline(url)
			if err != nil {
				log.Printf("Erro ao verificar live em '%s': %v\n", url, err)
				continue
			}

			if isOffline {
				log.Printf("A live em '%s' está OFFLINE (LIVE encerrada)\n", url)
				if contains(onlineLives, url) {
					onlineLives = remove(onlineLives, url)
				}
			} else {
				log.Printf("A live em '%s' ainda está ONLINE\n", url)
				if !contains(onlineLives, url) {
					onlineLives = append(onlineLives, url)
					go download.GetVideo(url)
				}
			}
		}

		time.Sleep(10 * time.Second)
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func remove(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
