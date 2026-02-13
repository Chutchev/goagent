package agent

import (
	"bufio"
	"context"
	"fmt"
	"github.com/Chutchev/goagent/pkg/clients/llm"
	"github.com/Chutchev/goagent/pkg/config"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Agent struct {
	systemPrompt string
	Name         string
	Mode         string
}

func NewAgent(promptFile *string, name, mode string) *Agent {
	systemPrompt, err := os.ReadFile(*promptFile)
	if err != nil {
		log.Fatalf("system prompt file read failed: %v", err)
	}
	return &Agent{
		systemPrompt: string(systemPrompt),
		Name:         name,
		Mode:         mode,
	}
}

func (a *Agent) runInteractive() {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Введите текст (пустая строка для завершения):")
		for {
			fmt.Print("> ")
			scanner.Scan()
			text := scanner.Text()

			// Если строка пустая - завершаем ввод
			if text == "" {
				break
			}

			lines = append(lines, text)
		}
		prompt := strings.Join(lines, "\n")
		a.do(prompt)
	}
}

func (a *Agent) runGRPC() {
}

func (a *Agent) runHTTP() {

}

func (a *Agent) survey() {}

func (a *Agent) Run() {
	switch a.Mode {
	case "i":
		go a.runInteractive()
	case "grpc":
		go a.runGRPC()
	case "http":
		go a.runHTTP()
	default:
		log.Fatal("")
	}

	// Сюда вставить запуск HTTP Сервера
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("Получен сигнал завершения, останавливаемся...")
}

func (a *Agent) do(userPrompt string) {
	cfg := config.GetConfig()

	c := llm.NewLLMClient(
		cfg.LLMBaseURL,
		cfg.LLMConfig.LLMToken,
	)

	req := llm.ChatRequest{
		Model:       cfg.LLMConfig.LLMModel,
		Temperature: cfg.LLMConfig.Temperature,
		TopP:        cfg.LLMConfig.TopP,
		Seed:        cfg.LLMConfig.Seed,
		Messages: []llm.Message{
			{
				Role:    "system",
				Content: a.systemPrompt,
			},
			{
				Role:    "user",
				Content: userPrompt,
			},
		},
	}
	r, err := c.CreateChatCompletion(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Choices[0].Message.Content)
}
