package config

import (
	"context"
	"log"
	"sync"

	"github.com/sethvargo/go-envconfig"
)

type LLMConfig struct {
	LLMToken    string  `env:"LLM_TOKEN,default=token"`
	LLMModel    string  `env:"LLM_MODEL,default=cotype_pro_2.5"`
	Temperature float64 `env:"LLM_TEMPERATURE,default=0.0"`
	TopP        float64 `env:"LLM_TOP_P,default=1.0"`
	Seed        int64   `env:"LLM_SEED,default=42"`
}

type Config struct {
	LLMBaseURL string `env:"LLM_BASE_URL,default=https://demo5-fundres.dev.mts.ai/v1"`
	LLMConfig  LLMConfig
}

var (
	cfg  *Config
	once sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}

		ctx := context.Background()
		if err := envconfig.Process(ctx, cfg); err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	})

	return cfg
}
