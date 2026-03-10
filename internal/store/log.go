package store

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type Entry struct {
	Time    time.Time `json:"time"`
	Src     string    `json:"src"`
	Dst     string    `json:"dst"`
	Bytes   int64     `json:"bytes"`
	Fortune string    `json:"fortune"`
}

func logPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".hikyaku", "log.json"), nil
}

func Append(e Entry) error {
	path, err := logPath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}

	entries, err := Load()
	if err != nil {
		return err
	}

	entries = append(entries, e)

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func Load() ([]Entry, error) {
	path, err := logPath()
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []Entry{}, nil
	}
	if err != nil {
		return nil, err
	}

	var entries []Entry
	if err := json.Unmarshal(data, &entries); err != nil {
		return nil, err
	}
	return entries, nil
}
