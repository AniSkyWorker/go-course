package processor

import (
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type VideoProcessor interface {
	GetVideoDuration(videoPath string) (float64, error)
	CreateVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error
}

type FfmpegVideoProcessor struct {
}

func (pr *FfmpegVideoProcessor) GetVideoDuration(videoPath string) (float64, error) {
	result, err := exec.Command(`ffprobe`, `-v`, `error`, `-show_entries`, `format=duration`, `-of`, `default=noprint_wrappers=1:nokey=1`, videoPath).Output()
	if err != nil {
		return 0.0, err
	}

	return strconv.ParseFloat(strings.Trim(string(result), "\n\r"), 64)
}

func (pr *FfmpegVideoProcessor) CreateVideoThumbnail(videoPath string, thumbnailPath string, thumbnailOffset int64) error {
	return exec.Command(`ffmpeg`, `-i`, videoPath, `-ss`, ffmpegTimeFromSeconds(thumbnailOffset), `-vframes`, `1`, thumbnailPath, `-y`).Run()
}

func ffmpegTimeFromSeconds(seconds int64) string {
	return time.Unix(seconds, 0).UTC().Format(`15:04:05.000000`)
}
