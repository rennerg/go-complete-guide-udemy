package utilities

import (
	"encoding/json"
	"errors"
	"os"
)

func ReadFromFile(fileName string, prop string) (float64, error) {
	if _, err := os.Stat(fileName); err != nil {
		WriteToFile(fileName, prop, 0)
	}
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("ERROR: could not read " + prop + " file")
	}
	var result map[string]float64
	err = json.Unmarshal(data, &result)
	if err != nil {
		return 0, errors.New("ERROR: could not parse " + prop + " file")
	}
	balance := result[prop]
	return balance, nil
}

func WriteToFile(fileName string, prop string, value float64) error {
	data, err := json.Marshal(map[string]float64{prop: value})
	if err != nil {
		return errors.New("ERROR: could not encode " + prop + " data")
	}
	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return errors.New("ERROR: could not write " + prop + " file")
	}
	return nil
}
