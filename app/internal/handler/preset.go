package handler

import (
	"encoding/json"
	"os"
	"path/filepath"
	"time"
)

type PresetFile struct {
	Address uint16       `json:"address"`
	Items   []PresetItem `json:"items"`
}

type PresetItem struct {
	Command      uint16          `json:"command"`
	Action       IrActionId      `json:"action"`
	Repeatable   bool            `json:"repeatable"`
	HoldDuration uint16          `json:"hold_duration"` // in milliseconds
	Params       json.RawMessage `json:"params"`
}

func ParsePresetDir(dirPath string) (map[string]BoundAction, error) {
	newMap := make(map[string]BoundAction)

	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			fullPath := filepath.Join(dirPath, file.Name())

			data, err := os.ReadFile(fullPath)
			if err != nil {
				return nil, err
			}

			var fileData PresetFile
			if err := json.Unmarshal(data, &fileData); err != nil {
				return nil, err
			}

			address := fileData.Address

			for _, item := range fileData.Items {
				key := IrKey{
					Address: address,
					Command: item.Command,
				}

				holdMs := item.HoldDuration
				if holdMs > 10000 {
					holdMs = 10000
				}

				newMap[key.String()] = BoundAction{
					ActionId:     item.Action,
					Repeatable:   item.Repeatable,
					HoldDuration: time.Duration(holdMs) * time.Millisecond,
					Params:       item.Params,
				}
			}
		}
	}
	return newMap, nil
}
