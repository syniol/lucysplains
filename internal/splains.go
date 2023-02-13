package internal

import (
	"github.com/syniol/lucysplains/pkg/openai/completion"
)

type LucySplains struct {
	Client *completion.Completion
}

func NewLucySplains() (*LucySplains, error) {
	client, err := completion.NewDaVinciCompletion()
	if err != nil {
		return nil, err
	}

	return &LucySplains{Client: client}, nil
}

func (l *LucySplains) Splain(msg string) (interface{}, error) {
	return l.Client.Create([]string{msg})
}
