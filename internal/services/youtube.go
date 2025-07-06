package services

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"path/filepath"
	"telegramBotInstaller/internal/utils"
)

func DownloadYouTubeVideo(youtubeURL, downloadDir string) (string, error) {
	outputTemplate := filepath.Join(downloadDir, "%(id)s.%(ext)s")
	cmd := exec.Command("yt-dlp", "-f", "mp4", "-o", outputTemplate, youtubeURL)

	log.Infof("📥 Команда скачивания: yt-dlp -f mp4 -o \"%s\" %s", outputTemplate, youtubeURL)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("❌ yt-dlp error: %s", string(output))
		return "", fmt.Errorf("failed to download YouTube video: %w", err)
	}

	videoID, err := utils.ExtractVideoID(youtubeURL)
	if err != nil {
		return "", err
	}

	videoPath := filepath.Join(downloadDir, videoID+".mp4")
	return videoPath, nil
}
