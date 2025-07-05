package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DownloadVideo(bot *tgbotapi.BotAPI, fileID, downloadDir string) (string, error) {
	file, err := bot.GetFile(tgbotapi.FileConfig{FileID: fileID})
	if err != nil {
		return "", fmt.Errorf("failed to get file: %v", err)
	}

	fileURL := fmt.Sprintf("https://api.telegram.org/file/bot%s/%s", bot.Token, file.FilePath)

	resp, err := http.Get(fileURL)
	if err != nil {
		return "", fmt.Errorf("download error: %v", err)
	}
	defer resp.Body.Close()

	err = os.MkdirAll(downloadDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create folder: %v", err)
	}

	videoPath := filepath.Join(downloadDir, fileID+".mp4")

	out, err := os.Create(videoPath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %v", err)
	}

	return videoPath, nil
}
