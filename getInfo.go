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

func getImageInfo(infoFile string) ImageInfo {
	if infoFile == "" {
		infoFile = "/usr/share/ublue-os/image-info.json"
	}

	data, err := os.ReadFile(infoFile)
	if err != nil {
		return ImageInfo{"", ""}
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
	cmd_grep := exec.Command("grep", "-q", "status is GREEN", "/etc/motd.d/boot-status")
	err := cmd_grep.Run()
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`status is GREEN`)

	isGreen := re.FindString("status is GREEN")
	if isGreen != "" {
		return "healthy"
	} else {
		cmd := exec.Command("cat", "/etc/motd.d/boot-status")
		output, err := cmd.Output()
		if err != nil {
			return ""
		}
		return "`" + string(output) + "`"
	}
}

func isBootcSystem() bool {
	_, err := os.Stat("/run/ostree-booted")
	return err == nil
}

// func getDesktop() string {
// 	cmd := exec.Command("")
// 	return
// }
