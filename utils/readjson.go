package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson(filename string, jsonvar interface{}) error {
	f, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error when reading JSON file: %w", err)
	}
	err = json.Unmarshal(f, jsonvar)
	if err != nil {
		return fmt.Errorf("error when unmarshal JSON file: %w", err)
	}

	return nil
}
