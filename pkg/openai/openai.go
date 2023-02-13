package openai

import (
	"net/http"
	"os"
)

const KeyEnvName = "OPEN_AI_KEY"

type Credentials struct {
	APIKey string
}

type OpenAI struct {
	HttpClient  *http.Client
	Credentials Credentials
}

// NewOpenAI creates a new instance of OpenAI SDK
func NewOpenAI() OpenAI {
	return OpenAI{
		HttpClient: &http.Client{},
		Credentials: Credentials{
			APIKey: os.Getenv(KeyEnvName),
		},
	}
}
