package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
)

func downloadHLS(inputFile string, outputFile string, segmentDuration int) error {
	ffmpegCmd := exec.Command(
		"ffmpeg",
		"-i", inputFile,
		"-loglevel", "debug",
		"-start_number", "0",
		"-hls_time", strconv.Itoa(segmentDuration),
		"-hls_list_size", "0",
		"-f", "hls",
		"-c", "copy",
		outputFile,
	)

	output, err := ffmpegCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to download HLS: %v\nOutput: %s", err, string(output))
	}

	return nil
}

func main() {
	inputFile := "https://wework-5-us.stream.iot-11.com:8080/cloudrecord/eb675542a3f236ea384fak/csd9belqmqsujpkm66r0tp0rhpNkj9Cv.m3u8"
	outputFile := "videos.mp4"
	segmentDuration := 10

	if err := downloadHLS(inputFile, outputFile, segmentDuration); err != nil {
		log.Fatalf("Error download HLS: %v", err)
	}

	log.Println("HLS downloaded successfully")
}
