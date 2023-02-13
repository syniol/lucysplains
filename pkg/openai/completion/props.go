package completion

import (
	"errors"
)

type Temperature float32

func NewTemperature(temp float32) (*Temperature, error) {
	if temp > 2 || temp < 0 {
		return nil, errors.New("temperature is a float value between 0 - 2")
	}

	output := Temperature(temp)

	return &output, nil
}

type TopP float32

func NewTopP(topP float32) (*TopP, error) {
	if topP > 1 || topP < 0 {
		return nil, errors.New("")
	}

	out := TopP(topP)

	return &out, nil
}
