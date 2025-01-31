package utils

import "os"

func CheckSudo() {
	// Checks if the program is running as root
	if os.Geteuid() != 0 {
		println("This operation requires elevated permissions. Please run the program with `sudo`.")
		os.Exit(1)
	}
}
