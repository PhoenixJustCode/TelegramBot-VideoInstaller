package services

import (
	"fmt"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

func ConvertMp4ToMp3WithID(inputPath, outputDir, fileID string) (string, error) {
	outputPath := filepath.Join(outputDir, fileID + ".mp3")

	cmd := exec.Command("ffmpeg", "-i", inputPath, "-vn", "-acodec", "libmp3lame", "-ab", "192k", outputPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("ошибка конвертации: %v (%s)", err, string(output))
	}

	log.Infof("✅ MP3 сохранено: %s", outputPath)
	return outputPath, nil
}
