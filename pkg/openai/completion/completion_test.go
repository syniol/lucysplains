package completion

import (
	"encoding/json"
	"testing"
)

func TestCompletion_Config(t *testing.T) {
	cfg := &Config{
		Prompt: []string{"Hey Lad! wots happening?"},
	}

	act, _ := json.Marshal(cfg)

	expected := map[string]interface{}{
		"prompt": []string{
			"Hey Lad! wots happening?",
		},
	}

	a, _ := json.Marshal(expected)

	if string(act) != string(a) {
		t.Error("not matching expected")
	}
}
