package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	var device string = "No Device"

	cmd1 := exec.Command("pamixer", "--get-default-sink")
	defaultDevice, err := cmd1.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	cmd2 := exec.Command("pactl", "list", "sinks")
	stdout, err := cmd2.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defaultId := strings.Split(strings.Split(string(defaultDevice), "\n")[1], " ")[0]

	output := string(stdout)

	devicesMap := make(map[string]map[string]string)

	for _, strSlice := range strings.Split(output, "Sink #") {
		if len(strSlice) == 0 {
			continue
		}

		deviceInfo := strings.SplitN(strSlice, "\n", 2)
		infoMap := make(map[string]string)
		devicesMap[deviceInfo[0]] = infoMap

		for _, str := range strings.Split(deviceInfo[1], "\n") {
			withoutTabs := strings.ReplaceAll(str, "\t", "")
			split := strings.Split(withoutTabs, getSeparator(str))
			if len(split) == 2 {
				infoMap[split[0]] = split[1]
			}
		}
	}

	device = devicesMap[defaultId]["alsa.card_name"]

	fmt.Println(strings.ReplaceAll(device, `"`, ""))
}

func getSeparator(str string) string {
	for _, char := range str {
		if char == '=' {
			return " = "
		}
		if char == ':' {
			return ": "
		}
	}

	return ""
}
