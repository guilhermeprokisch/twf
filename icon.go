
package main

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func getNerdFontIcon(filename string) (string, error) {
	// Run the `exa` command to get the icon for the file
	cmd := exec.Command("lsd", "--icon", "always",filename)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	// Parse the output of `lsd` to get the icon
	iconLine := strings.Split(out.String(), "\n")[0]
	return iconLine[0:3], nil
}

func main() {
	filename := "screenshot.png"
	icon, err := getNerdFontIcon(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("NerdFont icon for %s: %s\n", filename, icon)
	}
}
