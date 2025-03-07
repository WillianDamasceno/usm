package color

func LinesAround(format string) string {
	return "\n" + format + "\n"
}

func PrintLinesAround(format string) {
	print(LinesAround(format))
}

// ANSI escape codes functions
func Reset() string {
	return "\033[0m" // Reset to default
}

func Bold(format string) string {
	return "\033[1m" + format + Reset()
}

func Dim(format string) string {
	return "\033[2m" + format + Reset()
}

func Italic(format string) string {
	return "\033[3m" + format + Reset()
}

func Underline(format string) string {
	return "\033[4m" + format + Reset()
}

func Blink(format string) string {
	return "\033[5m" + format + Reset()
}

func Reverse(format string) string {
	return "\033[7m" + format + Reset()
}

func Hidden(format string) string {
	return "\033[8m" + format + Reset()
}

// Foreground color functions
func Black(format string) string {
	return "\033[30m" + format + Reset()
}

func Red(format string) string {
	return "\033[31m" + format + Reset()
}

func Green(format string) string {
	return "\033[32m" + format + Reset()
}

func Yellow(format string) string {
	return "\033[33m" + format + Reset()
}

func Blue(format string) string {
	return "\033[34m" + format + Reset()
}

func Magenta(format string) string {
	return "\033[35m" + format + Reset()
}

func Cyan(format string) string {
	return "\033[36m" + format + Reset()
}

func White(format string) string {
	return "\033[37m" + format + Reset()
}

// Bright foreground color functions
func BrightBlack(format string) string {
	return "\033[90m" + format + Reset()
}

func BrightRed(format string) string {
	return "\033[91m" + format + Reset()
}

func BrightGreen(format string) string {
	return "\033[92m" + format + Reset()
}

func BrightYellow(format string) string {
	return "\033[93m" + format + Reset()
}

func BrightBlue(format string) string {
	return "\033[94m" + format + Reset()
}

func BrightMagenta(format string) string {
	return "\033[95m" + format + Reset()
}

func BrightCyan(format string) string {
	return "\033[96m" + format + Reset()
}

func BrightWhite(format string) string {
	return "\033[97m" + format + Reset()
}

// Background color functions
func BackgroundBlack(format string) string {
	return "\033[40m" + format + Reset()
}

func BackgroundRed(format string) string {
	return "\033[41m" + format + Reset()
}

func BackgroundGreen(format string) string {
	return "\033[42m" + format + Reset()
}

func BackgroundYellow(format string) string {
	return "\033[43m" + format + Reset()
}

func BackgroundBlue(format string) string {
	return "\033[44m" + format + Reset()
}

func BackgroundMagenta(format string) string {
	return "\033[45m" + format + Reset()
}

func BackgroundCyan(format string) string {
	return "\033[46m" + format + Reset()
}

func BackgroundWhite(format string) string {
	return "\033[47m" + format + Reset()
}

// Bright background color functions
func BrightBackgroundBlack(format string) string {
	return "\033[100m" + format + Reset()
}

func BrightBackgroundRed(format string) string {
	return "\033[101m" + format + Reset()
}

func BrightBackgroundGreen(format string) string {
	return "\033[102m" + format + Reset()
}

func BrightBackgroundYellow(format string) string {
	return "\033[103m" + format + Reset()
}

func BrightBackgroundBlue(format string) string {
	return "\033[104m" + format + Reset()
}

func BrightBackgroundMagenta(format string) string {
	return "\033[105m" + format + Reset()
}

func BrightBackgroundCyan(format string) string {
	return "\033[106m" + format + Reset()
}

func BrightBackgroundWhite(format string) string {
	return "\033[107m" + format + Reset()
}
