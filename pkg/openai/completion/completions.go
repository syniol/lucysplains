package completion

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/syniol/lucysplains/pkg/openai"
)

const Endpoint = "https://api.openai.com/v1/completions"

type Completion struct {
	openAIClient openai.OpenAI
	config       *Config
}

type Config struct {
	Model       openai.Model `json:"model,omitempty"`
	Temperature *Temperature `json:"temperature,omitempty"`
	TopP        *TopP        `json:"top_p,omitempty"`
	Prompt      []string     `json:"prompt,omitempty"`
}

func NewDaVinciCompletion() (*Completion, error) {
	topP, err := NewTopP(0.1)
	if err != nil {
		return nil, err
	}

	temp, err := NewTemperature(0.1)
	if err != nil {
		return nil, err
	}

	return &Completion{
		openAIClient: openai.NewOpenAI(),
		config: &Config{
			Model:       openai.DaVinci,
			Temperature: temp,
			TopP:        topP,
		},
	}, nil
}

func (c *Completion) Create(prompt []string) (interface{}, error) {
	localCfg := *c.config
	localCfg.Prompt = prompt

	input, err := json.Marshal(localCfg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		Endpoint,
		strings.NewReader(string(input)),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.openAIClient.Credentials.APIKey))
	requestResp, err := c.openAIClient.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := ioutil.ReadAll(requestResp.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
