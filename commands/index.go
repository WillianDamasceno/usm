package commands

import (
	"fmt"
	"os/exec"
	"usm/utils/color"
)

func aptUpdate() {
	cmd := exec.Command("sudo", "apt", "update")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Output:\n%s\n", output)
}

func aptUpgrade() {
	cmd := exec.Command("sudo", "apt", "upgrade", "-y")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error:%s\n", err)
		return
	}

	fmt.Printf("Output:\n%s\n", output)
}

func UpdateSystem() {
	color.PrintLinesAround(color.BackgroundGreen("Updating the system..."))

	color.PrintLinesAround(color.Green("[1/2] Executing sudo apt update...\n"))
	aptUpdate()

	color.PrintLinesAround(color.Green("[2/2] Executing sudo apt upgrade...\n"))
	aptUpgrade()
}
