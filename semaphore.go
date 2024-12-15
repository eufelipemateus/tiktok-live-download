package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eufelipemateus/go-video/download"
	"github.com/eufelipemateus/go-video/status"
	"github.com/eufelipemateus/go-video/utils"
)

func Run() {
	onlineLives := []string{}

	for {
		liveURLs := loadURLsFromFile("lives.txt")

		for _, url := range liveURLs {
			isOffline, err := status.IsLiveOffline(url)
			if err != nil {
				log.Printf("Erro ao verificar live em  %s '%s' %s:  %s %v %s \n", utils.Colors.Blue, url,  utils.Colors.Reset, utils.Colors.Red, err, utils.Colors.Reset)
				continue
			}

			if isOffline {
				log.Printf("A live em '%s%s%s' está %sOFFLINE%s (LIVE encerrada)\n", utils.Colors.Blue, url, utils.Colors.Reset, utils.Colors.Red, utils.Colors.Reset)
				if contains(onlineLives, url) {
					onlineLives = remove(onlineLives, url)
				}
			} else {
				log.Printf("A live em '%s%s%s' ainda está %sONLINE%s\n", utils.Colors.Blue, url, utils.Colors.Reset, utils.Colors.Green, utils.Colors.Reset)
				if !contains(onlineLives, url) {
					onlineLives = append(onlineLives, url)
					go download.GetVideo(url)
				}
			}
		}

		time.Sleep(1 * time.Minute)
	}
}

func loadURLsFromFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)

	// Lê o arquivo linha por linha
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())         // Remove espaços em branco extras
		if line != "" && !strings.HasPrefix(line, "//") { // Ignora linhas vazias ou comentários
			urls = append(urls, line)
		}
	}

	// Verifica erros durante a leitura
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	return urls
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
