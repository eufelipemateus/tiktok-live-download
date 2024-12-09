package youtube

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

// Caminho do arquivo de credenciais e token
const (
	credentialsFile = "./credentials.json"
	tokenFile       = "./token.json"
)

// Lê o token do arquivo local
func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var token oauth2.Token
	err = json.NewDecoder(f).Decode(&token)
	return &token, err
}

// Salva o token em um arquivo local
func saveToken(file string, token *oauth2.Token) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalf("Unable to create token file: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Obtém o cliente HTTP autorizado
func getClient() *http.Client {
	// Lê o arquivo de credenciais
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// Configura o OAuth2
	config, err := google.ConfigFromJSON(b, youtube.YoutubeUploadScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Verifica se o token já está salvo
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		// Se o token não estiver salvo, obtém via web
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}

	// Renova o token se necessário
	if token.Expiry.Before(time.Now()) {
		tokenSource := config.TokenSource(context.Background(), token)
		token, err = tokenSource.Token()
		if err != nil {
			log.Fatalf("Unable to refresh token: %v", err)
		}
		saveToken(tokenFile, token)
	}

	return config.Client(context.Background(), token)
}

// Obtém o token via URL de autenticação
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("", oauth2.AccessTypeOffline)
	fmt.Println("Visit the URL for the auth code: ", authURL)

	// Lê o código de autorização
	fmt.Println("Enter the authorization code: ")
	var code string
	_, err := fmt.Scan(&code)
	if err != nil {
		log.Fatalf("Unable to read the authorization code: %v", err)
	}

	// Troca o código pelo token
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return token
}

func UploadVideo(videoFilePath string, title string, description string) {
	client := getClient()

	// Cria um novo serviço do YouTube
	service, err := youtube.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// Prepara o arquivo de vídeo para upload
	videoFile, err := os.Open(videoFilePath)
	if err != nil {
		log.Fatalf("Error opening video file: %v", err)
	}
	defer videoFile.Close()

	// Cria a solicitação de upload
	call := service.Videos.Insert([]string{"snippet", "status"}, &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:       title,
			Description: description,
			Tags:        []string{"live", "video", "upload"},
		},
		Status: &youtube.VideoStatus{
			PrivacyStatus: "unlisted", // Pode ser "private", "public" ou "unlisted"
		},
	})
	// Realiza o upload
	_, err = call.Media(videoFile).Do()
	if err != nil {
		log.Fatalf("Error uploading video: %v", err)
	}

	log.Println("Video uploaded successfully!")
}
