package browser

import (
	"context"
	"encoding/json"
	"log"

	"github.com/chromedp/chromedp"
	"github.com/eufelipemateus/go-video/interfaces"
)

type Browser struct {
	ctx    context.Context
	cancel context.CancelFunc
	Url    string
	Active bool
}

func NewBrowser(url string) *Browser {

	ctx, _ := context.WithCancel(context.Background())

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),                       // Executar no modo headless
		chromedp.Flag("disable-gpu", true),                     // Desativar GPU
		chromedp.Flag("blink-settings", "imagesEnabled=false"), // Desativa imagens
		chromedp.Flag("proxy-server", "socks5://tor:9050"),  
	)

	// Criando um contexto de chromedp com timeout
	ctx, _ = chromedp.NewExecAllocator(ctx, opts...)

	ctx, cacancelBrpwser := chromedp.NewContext(ctx)

	chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
	)

	return &Browser{ctx, cacancelBrpwser, url, false}
}


func (b *Browser) GetSigState() (interfaces.SigiState, error) {
	var sigiStateContent string

	err := chromedp.Run(b.ctx,
		chromedp.InnerHTML("#SIGI_STATE", &sigiStateContent, chromedp.ByID),
	)

	if err != nil {
		return interfaces.SigiState{}, err
	}

	var data interfaces.SigiState
	if err := json.Unmarshal([]byte(sigiStateContent), &data); err != nil {
		log.Fatalf("Erro ao deserializar o JSON: %v", err)
	}

	return data, nil
}

func (b *Browser) Reload() error {
	return chromedp.Run(b.ctx,
		chromedp.Reload(),
	)
}
