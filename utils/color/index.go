package color

import "fmt"

// ANSI escape codes functions
func Reset() string {
	return "\033[0m" // Reset to default
}

func Bold(format string, args ...interface{}) {
	fmt.Printf("\033[1m"+format+Reset(), args...)
}

func Dim(format string, args ...interface{}) {
	fmt.Printf("\033[2m"+format+Reset(), args...)
}

func Italic(format string, args ...interface{}) {
	fmt.Printf("\033[3m"+format+Reset(), args...)
}

func Underline(format string, args ...interface{}) {
	fmt.Printf("\033[4m"+format+Reset(), args...)
}

func Blink(format string, args ...interface{}) {
	fmt.Printf("\033[5m"+format+Reset(), args...)
}

func Reverse(format string, args ...interface{}) {
	fmt.Printf("\033[7m"+format+Reset(), args...)
}

func Hidden(format string, args ...interface{}) {
	fmt.Printf("\033[8m"+format+Reset(), args...)
}

// Foreground color functions
func Black(format string, args ...interface{}) {
	fmt.Printf("\033[30m"+format+Reset(), args...)
}

func Red(format string, args ...interface{}) {
	fmt.Printf("\033[31m"+format+Reset(), args...)
}

func Green(format string, args ...interface{}) {
	fmt.Printf("\033[32m"+format+Reset(), args...)
}

func Yellow(format string, args ...interface{}) {
	fmt.Printf("\033[33m"+format+Reset(), args...)
}

func Blue(format string, args ...interface{}) {
	fmt.Printf("\033[34m"+format+Reset(), args...)
}

func Magenta(format string, args ...interface{}) {
	fmt.Printf("\033[35m"+format+Reset(), args...)
}

func Cyan(format string, args ...interface{}) {
	fmt.Printf("\033[36m"+format+Reset(), args...)
}

func White(format string, args ...interface{}) {
	fmt.Printf("\033[37m"+format+Reset(), args...)
}

// Bright foreground color functions
func BrightBlack(format string, args ...interface{}) {
	fmt.Printf("\033[90m"+format+Reset(), args...)
}

func BrightRed(format string, args ...interface{}) {
	fmt.Printf("\033[91m"+format+Reset(), args...)
}

func BrightGreen(format string, args ...interface{}) {
	fmt.Printf("\033[92m"+format+Reset(), args...)
}

func BrightYellow(format string, args ...interface{}) {
	fmt.Printf("\033[93m"+format+Reset(), args...)
}

func BrightBlue(format string, args ...interface{}) {
	fmt.Printf("\033[94m"+format+Reset(), args...)
}

func BrightMagenta(format string, args ...interface{}) {
	fmt.Printf("\033[95m"+format+Reset(), args...)
}

func BrightCyan(format string, args ...interface{}) {
	fmt.Printf("\033[96m"+format+Reset(), args...)
}

func BrightWhite(format string, args ...interface{}) {
	fmt.Printf("\033[97m"+format+Reset(), args...)
}

// Background color functions
func BackgroundBlack(format string, args ...interface{}) {
	fmt.Printf("\033[40m"+format+Reset(), args...)
}

func BackgroundRed(format string, args ...interface{}) {
	fmt.Printf("\033[41m"+format+Reset(), args...)
}

func BackgroundGreen(format string, args ...interface{}) {
	fmt.Printf("\033[42m"+format+Reset(), args...)
}

func BackgroundYellow(format string, args ...interface{}) {
	fmt.Printf("\033[43m"+format+Reset(), args...)
}

func BackgroundBlue(format string, args ...interface{}) {
	fmt.Printf("\033[44m"+format+Reset(), args...)
}

func BackgroundMagenta(format string, args ...interface{}) {
	fmt.Printf("\033[45m"+format+Reset(), args...)
}

func BackgroundCyan(format string, args ...interface{}) {
	fmt.Printf("\033[46m"+format+Reset(), args...)
}

func BackgroundWhite(format string, args ...interface{}) {
	fmt.Printf("\033[47m"+format+Reset(), args...)
}

// Bright background color functions
func BrightBackgroundBlack(format string, args ...interface{}) {
	fmt.Printf("\033[100m"+format+Reset(), args...)
}

func BrightBackgroundRed(format string, args ...interface{}) {
	fmt.Printf("\033[101m"+format+Reset(), args...)
}

func BrightBackgroundGreen(format string, args ...interface{}) {
	fmt.Printf("\033[102m"+format+Reset(), args...)
}

func BrightBackgroundYellow(format string, args ...interface{}) {
	fmt.Printf("\033[103m"+format+Reset(), args...)
}

func BrightBackgroundBlue(format string, args ...interface{}) {
	fmt.Printf("\033[104m"+format+Reset(), args...)
}

func BrightBackgroundMagenta(format string, args ...interface{}) {
	fmt.Printf("\033[105m"+format+Reset(), args...)
}

func BrightBackgroundCyan(format string, args ...interface{}) {
	fmt.Printf("\033[106m"+format+Reset(), args...)
}

func BrightBackgroundWhite(format string, args ...interface{}) {
	fmt.Printf("\033[107m"+format+Reset(), args...)
}
