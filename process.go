package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

type Metadata struct {
	Format Format `json:"format"`
}

type Format struct {
	Duration string `json:"duration"`
}

const TIME_SLICE = 120

func generateDrawTextFilters(chatFile string, videoHeight int, scrollSpeed float64, startDuration int) ([]string, error) {
	var filters []string

	// Definir o limite inferior (70% da altura da tela) e o limite superior (40% da altura da tela)
	bottomLimit := float64(videoHeight) * 0.7
	topLimit := float64(videoHeight) * 0.4
	// Onde o fade começa (30% da altura acima do limite inferior)
	fadeStart := bottomLimit - 0.3*float64(videoHeight)

	// Abrir o arquivo chat.txt
	file, err := os.Open(chatFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Extrair timestamp e mensagem do formato "10.63 Maria Espinoza: iluminemos la galeria 💚"
		parts := strings.SplitN(line, "-", 2)
		if len(parts) < 2 {
			continue // Linha inválida
		}

		// Converter timestamp para float
		timestampStr := strings.TrimSpace(parts[0])
		timestamp, err := strconv.ParseFloat(timestampStr, 64)
		if err != nil {
			return nil, err
		}

		if startDuration > int(timestamp) {
			continue
		}
		fmt.Printf("startDuration: %d, timestamp: %f\n", startDuration, timestamp)
		if (startDuration + TIME_SLICE) < int(timestamp) {
			break
		}

		message := parts[1]

		// Calcular a posição Y dinâmica, garantindo que o texto suba até o topo
		yPositionFormula := fmt.Sprintf("if(gte(t,%f),%f-(t-%f)*%f,%f)", timestamp, bottomLimit, timestamp, scrollSpeed, topLimit)

		// Fórmula de fade out (opacidade)
		alphaFormula := fmt.Sprintf("if(gte(t,%f),1,(%f-(t-%f)*%f)/%f)", fadeStart, fadeStart, timestamp, scrollSpeed, fadeStart)

		// Gerar filtro drawtext para cada linha com fade out
		filter := fmt.Sprintf(
			"drawtext=text='%s':fontcolor=white:fontsize=24:x=20:y='%s':box=1:boxcolor=black@0.5:enable='gte(t,%f)':alpha='%s'",
			message, yPositionFormula, timestamp, alphaFormula,
		)

		// Adicionar o filtro à lista de filtros
		filters = append(filters, filter)
	}

	// Checar por erros durante a leitura do arquivo
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return filters, nil
}

func GenerateVideoFinal(inputFile string, outputFile string, chatFile string) {

	destPath := "build/out"

	// Obter a duração do vídeo
	duration := getGurationVideo(inputFile)

	var parts = 1

	for i := 0; i < int(duration); i = i + TIME_SLICE {

		filters, err := generateDrawTextFilters(chatFile, 1280, 20, i)
		if err != nil {
			log.Fatalf("Erro ao gerar filtros: %v", err)
		}

		// Criar a string de filtro complexos
		filterComplex := strings.Join(filters, ",")

		err = ffmpeg_go.Input(inputFile, ffmpeg_go.KwArgs{
			"ss": i, // Define o início do corte
		}).
			Output(fmt.Sprintf("%s/%d.mp4", destPath, parts), ffmpeg_go.KwArgs{
				"t":       int(i + TIME_SLICE),
				"c:v":     "libx264",
				"crf":     "18",
				"preset":  "veryfast",
				"pix_fmt": "yuva420p",
				"vf":      filterComplex,
			}).ErrorToStdOut().
			Run()
		if err != nil {
			log.Fatalf("Erro ao executar FFmpeg: %v", err.Error())
		}
		parts++
	}

	os.Exit(0)

	/*// Gerar filtros de texto a partir do arquivo chat.txt
	filters, err := generateDrawTextFilters(chatFile, 1280, 20,)
	if err != nil {
		log.Fatalf("Erro ao gerar filtros: %v", err)
	}

	// Criar a string de filtro complexos
	filterComplex := strings.Join(filters, ",")

	// Executar o ffmpeg-go com os filtros
	err = ffmpeg_go.Input(inputFile).
		Output(outputFile, ffmpeg_go.KwArgs{
			"c:v":     "libx264",
			"crf":     "18",
			"preset":  "veryfast",
			"pix_fmt": "yuva420p",
			"vf":      filterComplex,
		}).ErrorToStdOut().
		Run()

	if err != nil {
		log.Fatalf("Erro ao executar FFmpeg: %v", err.Error())
	}*/

	log.Println("Processamento concluído! O vídeo com o chat foi gerado em", outputFile)

}

func getGurationVideo(videoPath string) float64 {

	// Executa o ffprobe e obtém os metadados

	out, err := ffmpeg_go.Probe(videoPath)
	if err != nil {
		log.Fatalf("Erro ao executar ffprobe: %v", err)
	}

	// Analisa o JSON retornado
	var meta Metadata
	if err := json.Unmarshal([]byte(out), &meta); err != nil {
		log.Fatalf("Erro ao analisar JSON: %v", err)
	}

	// Converte a duração para segundos
	duration, err := parseDuration(meta.Format.Duration)
	if err != nil {
		log.Fatalf("Erro ao converter duração: %v", err)
	}

	log.Printf("A duração do vídeo é: %.2f segundos\n", duration)

	return duration

}

// Converte a string de duração para um número em segundos
func parseDuration(durationStr string) (float64, error) {

	var duration float64
	_, err := fmt.Sscanf(durationStr, "%f", &duration)
	return duration, err
}
