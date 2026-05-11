package config

import (
	"app/internal/handler"
	"log/slog"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func WatchPresetDir(dirPath string, eh *handler.IrEventHandler, logger *slog.Logger) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		logger.Error("Failed to create file watcher", "error", err)
		return
	}

	err = watcher.Add(dirPath)
	if err != nil {
		logger.Error("Failed to watch preset directory", "error", err)
		return
	}

	logger.Info("Watching preset directory for changes", "directory", dirPath)

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				isJson := strings.HasSuffix(event.Name, ".json")

				isRelevantOp := event.Op.Has(fsnotify.Write) || event.Op.Has(fsnotify.Create) || event.Op.Has(fsnotify.Remove)
				if isJson && isRelevantOp {
					logger.Info("Preset file changed, reloading presets", "file", event.Name)

					updatedMap, err := handler.ParsePresetDir(dirPath)
					if err != nil {
						logger.Error("Failed to parse preset directory", "error", err)
						continue
					}

					eh.ReloadKeyActions(updatedMap)
				}

			case event, ok := <-watcher.Errors:
				if !ok {
					return
				}
				logger.Error("File watcher error", "error", event)

			}

		}
	}()
}
