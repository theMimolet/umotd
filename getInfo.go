package main

import (
	"encoding/json"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type ImageInfo struct {
	ImageRef string `json:"image-ref"`
	ImageTag string `json:"image-tag"`
}

func getImageInfo(file string) ImageInfo {
	infoFile := file
	if infoFile == "" {
		infoFile = "/usr/share/ublue-os/image-info.json"
	}

	data, err := os.ReadFile(infoFile)
	if err != nil {
		return ImageInfo{}
	}

	var info ImageInfo
	json.Unmarshal(data, &info)

	// strip the ostree prefix, same as the sed in bash
	info.ImageRef = strings.TrimPrefix(info.ImageRef, "ostree-image-signed:docker://")

	return info
}

func getOSName() string {
	// Gets the OS name from /etc/os-release
	data, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`NAME="(.*)"`)
	match := re.FindStringSubmatch(string(data))
	if len(match) > 1 {
		return match[1]
	}
	return "Your System"
}

func getGreenbootInfo() string {
	data, err := os.ReadFile("/run/motd.d/boot-status")
	if err == nil {
		return ""
	}
	re := regexp.MustCompile(string(data))

	isGreen := re.FindString("GREEN")
	isSuccess := re.FindString("SUCCESS")
	if isGreen != "" || isSuccess != "" {
		return "healthy"
	} else {
		cmd := exec.Command("cat", "/run/motd.d/boot-status")
		output, err := cmd.Output()
		if err != nil {
			return ""
		}
		return "`" + string(output) + "`"
	}
}

// func getDesktop() string {
// 	cmd := exec.Command("")
// 	return
// }
