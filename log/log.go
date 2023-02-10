package logs

import "github.com/fatih/color"

var info *color.Color
var error *color.Color
var warn *color.Color
var ai *color.Color

func copyright() {
	Warn("\n   ____                   ___    ____\n  / __ \\____  ___  ____  /   |  /  _/\n / / / / __ \\/ _ \\/ __ \\/ /| |  / /  \n/ /_/ / /_/ /  __/ / / / ___ |_/ /   \n\\____/ .___/\\___/_/ /_/_/  |_/___/   \n    /_/                              \n")
}

func Init() {
	info = color.New(color.FgGreen)
	error = color.New(color.FgRed)
	warn = color.New(color.FgYellow)
	ai = color.New(color.FgBlue)
	copyright()
}

func Info(log string) {
	info.Println(log)
}

func InfoF(log string) {
	info.Printf(log)
}

func Error(log string) {
	error.Println(log)
}

func Warn(log string) {
	warn.Println(log)
}

func AI(log string) {
	ai.Println(log)
}
