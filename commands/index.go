package commands

import (
	"fmt"
	"os/exec"
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
	println("Updating the system...")

	aptUpdate()
	aptUpgrade()
}
