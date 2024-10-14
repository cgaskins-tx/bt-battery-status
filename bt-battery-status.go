package main

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	// display program name & purpose
	fmt.Println("Bluetooth Device Battery Status")
	fmt.Println()

	// Create a new command
	cmd := exec.Command("upower", "-e")

	// Execute the upower -e command command and get the output
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// parse the multiline output output using a scanner
	devicelist := bufio.NewScanner(strings.NewReader(string(out)))

	// iterate over each line
	for devicelist.Scan() {
		line := devicelist.Text() // get current line

		// now execute the upower info --show-info command with the appropriate parm from the first command
		cmd2 := exec.Command("upower", "--show-info", line)
		out2, err2 := cmd2.Output()
		if err2 != nil {
			fmt.Println("Error:", err2)
			return
		}

		// parse the output with a scanner
		devicestatus := bufio.NewScanner(strings.NewReader(string(out2)))

		// iterate through the status data
		native_device := false
		word := ""
		for devicestatus.Scan() {
			statusline := devicestatus.Text() // get current line
			word = "native-path"              // determine if it is a native device
			if strings.Contains(statusline, word) {
				// fmt.Println("Status: ", statusline)
				native_device = true
			}
			word = "model" // look for the model name
			if strings.Contains(statusline, word) && native_device == true {
				device_name := strings.Replace(statusline, "model:", "", -1)
				fmt.Println(strings.TrimLeft(device_name, " "))
			}
			word = "percentage" // look for the battery percentage
			if strings.Contains(statusline, word) && native_device == true {
				battery_percent := strings.Replace(statusline, "percentage:", "", -1)
				fmt.Println("Battery Remaining: ", strings.TrimLeft(battery_percent, " "))
			}
		}
		if native_device == true {
			fmt.Println(" ")
		}

	}

	// check for any errors during scanning
	if err := devicelist.Err(); err != nil {
		fmt.Println("Error while reading lines: ", err)
	}
	fmt.Println("Done.")
}
